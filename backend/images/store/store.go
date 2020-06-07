package store

import (
	"bytes"
	"context"
	"fmt"
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

type Store interface {
	Create(ctx context.Context, meta *img.Metadata, data *bytes.Buffer) (Data, error)
	Retrieve(ctx context.Context, id string) (Data, error)
	Delete(ctx context.Context, id string) error
}

type Data interface {
	URI() string
	Metadata() (*img.Metadata, error)
	Id() string
}

type DiskStore struct {
	mutex  *sync.RWMutex
	dir    string
	images map[string]*DiskData
	fs     afero.Fs
}

type DiskData struct {
	Filename string
}

func (d *DiskData) UserId() (uint64, error) {
	return userIdFromDiskFilename(d.Filename)
}

func userIdFromDiskFilename(f string) (uint64, error) {
	split := strings.Split(f, "_")
	withExt := split[len(split)]
	strId := strings.TrimSuffix(withExt, filepath.Ext(withExt))
	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (d *DiskData) Metadata() (*img.Metadata, error) {
	meta := &img.Metadata{
		Ext: img.Ext(filepath.Ext(d.Filename)),
	}
	var err error
	meta.UserId, err = userIdFromDiskFilename(d.Filename)
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

func (store *DiskStore) Create(ctx context.Context, meta *img.Metadata, data *bytes.Buffer) (Data, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot generate image id: %w", err)
	}
	id := uid.String()
	imagePath := fmt.Sprintf("%s/%s_%s.%s", store.dir, id, strconv.Itoa(int(meta.UserId)), meta.Ext)
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
