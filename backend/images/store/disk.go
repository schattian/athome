package store

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/athomecomar/athome/backend/images/img"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type DiskStore struct {
	mutex  *sync.RWMutex
	dir    string
	images map[string]*DiskData
	fs     afero.Fs
}

type DiskData struct {
	Filename string
}

func entityFromDiskFilename(f string) (ent img.Entity, err error) {
	split := strings.Split(f, "_")
	withExt := split[len(split)]
	strEntity := strings.TrimSuffix(withExt, filepath.Ext(withExt))
	entity := strings.Split(strEntity, "-")
	if len(entity) != 2 {
		err = errors.New("entity len is invalid")
		return
	}
	var strEntityId string
	strEntityId, ent.Table = entity[0], entity[1]
	entId, err := strconv.Atoi(strEntityId)
	if err != nil {
		err = errors.Wrap(err, "strconv.Atoi")
		return
	}
	ent.Id = uint64(entId)
	return
}

func (d *DiskData) Metadata() (*img.Metadata, error) {
	meta := &img.Metadata{
		Ext: img.Ext(filepath.Ext(d.Filename)),
	}
	var err error
	meta.Entity, err = entityFromDiskFilename(d.Filename)
	if err != nil {
		return nil, errors.Wrap(err, "userIdFromDiskFilename")
	}
	return meta, nil
}
func (d *DiskData) URI() string {
	return d.Filename
}

func (d *DiskData) Id() string {
	return strings.Split(d.Filename, ".")[0]
}

func NewDiskImageStore(fs afero.Fs, dir string) *DiskStore {
	return &DiskStore{
		mutex:  &sync.RWMutex{},
		dir:    dir,
		fs:     fs,
		images: make(map[string]*DiskData),
	}
}

func (store *DiskStore) Read(dd Data) (io.Reader, error) {
	file, err := store.fs.Open(store.dir + "/" + dd.URI())
	if err != nil {
		return nil, errors.Wrap(err, "fs.Open")
	}
	return file, nil
}

func (store *DiskStore) Create(ctx context.Context, meta *img.Metadata, data *bytes.Buffer) (Data, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot generate image id: %w", err)
	}
	id := uid.String()
	imagePath := fmt.Sprintf("%s/%s_%d-%s.%s", store.dir, id, meta.Entity.Id, meta.Entity.Table, meta.Ext)
	file, err := store.fs.Create(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = data.WriteTo(file)
	if err != nil {
		return nil, fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	dd := &DiskData{
		Filename: filepath.Base(imagePath),
	}
	store.images[id] = dd

	return dd, nil
}

func (store *DiskStore) Retrieve(ctx context.Context, id string) (Data, error) {
	dd, ok := store.images[id]
	if ok {
		return dd, nil
	}
	files, err := afero.ReadDir(store.fs, store.dir)
	if err != nil {
		return nil, errors.Wrap(err, "ReadDir")
	}
	var fi os.FileInfo
	for _, f := range files {
		fid := strings.Split(f.Name(), "_")[0]
		if fid == id {
			fi = f
		}
	}
	if fi == nil {
		return nil, errors.New("couldnt find file with id: " + id)
	}
	dd = &DiskData{Filename: fi.Name()}
	store.images[fi.Name()] = dd
	return dd, nil
}

func (store *DiskStore) RetrieveMany(ctx context.Context, entityId uint64, entityTable string) ([]Data, error) {
	files, err := afero.ReadDir(store.fs, store.dir)
	if err != nil {
		return nil, errors.Wrap(err, "ReadDir")
	}
	var dds []Data
	var fi os.FileInfo
	for _, f := range files {
		ent, err := entityFromDiskFilename(f.Name())
		if err != nil {
			return nil, errors.Wrap(err, "entityFromDiskFilename")
		}
		if ent.Table != entityTable || ent.Id != entityId {
			continue
		}
		dds = append(dds, &DiskData{Filename: fi.Name()})
	}
	return dds, nil
}

func (store *DiskStore) Delete(ctx context.Context, id string) error {
	dd, ok := store.images[id]
	if ok {
		return fmt.Errorf("image with id %s wasnt found", id)
	}
	err := store.fs.Remove(dd.Filename)
	if err != nil {
		return errors.Wrap(err, "fs.Remove")
	}
	return nil
}
