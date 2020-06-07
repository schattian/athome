package imageconf

import (
	"github.com/athomecomar/athome/backend/images/store"
	"github.com/athomecomar/envconf"
	"github.com/spf13/afero"
)

func GetSTORE() (s store.Store) {
	switch envconf.GetENV() {
	case envconf.Development:
		s = store.NewDiskImageStore(GetFILESYSTEM(), "tmp")
	case envconf.Staging, envconf.Production:
		// port = ":9903"
	}
	return
}

func GetFILESYSTEM() (fs afero.Fs) {
	switch envconf.GetENV() {
	case envconf.Development:
		fs = afero.NewOsFs()
	case envconf.Staging, envconf.Production:
		// port = ":9903"
	}
	return
}

func GetMAX_IMAGE_SIZE() (sz int64) {
	switch envconf.GetENV() {
	case envconf.Development:
		sz = 5 * mb
	case envconf.Staging, envconf.Production:
		sz = 10 * mb
	}
	return
}

const (
	kb = 1e3
	mb = 1e6
)
