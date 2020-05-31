package infer

import (
	"encoding/json"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/identifierconf"
	"github.com/athomecomar/athome/backend/identifier/normalize"
	"github.com/athomecomar/semantic/semprov"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type Lawyer struct {
	Tome    uint64 `json:"tome,omitempty"`
	Folio   uint64 `json:"folio,omitempty"`
	Surname string `json:"surname,omitempty"`
	Name    string `json:"name,omitempty"`
}

func tomeAndFolioByFullnameAttorneyAndLawyers(c semprov.Category) tomeAndFolioByFullnameByCategory {
	return func(fs afero.Fs, name, surname string) (tome uint64, folio uint64, err error) {
		f, err := fs.Open(identifierconf.GetDATA_DIR() + "/" + ByFullnameFilenames[semprov.Lawyer])
		if err != nil {
			err = errors.Wrap(err, "fs.Open")
			return
		}

		var lawyers []*Lawyer
		err = json.NewDecoder(f).Decode(&lawyers)
		if err != nil {
			err = errors.Wrap(err, "json.Decode")
			return
		}
		givenSurnameWords, givenNameWords := strings.Split(surname, " "), strings.Split(name, " ")

		var eq bool
		for _, lawyer := range lawyers {
			surnameWords, nameWords := strings.Split(lawyer.Surname, " "), strings.Split(lawyer.Name, " ")

			eq, err = normalize.CompareSlice(surnameWords, givenSurnameWords)
			if err != nil {
				err = errors.Wrap(err, "compareSlice on surnameWords")
				return
			}
			if !eq {
				continue
			}
			eq, err = normalize.CompareSliceSoft(nameWords, givenNameWords)
			if err != nil {
				err = errors.Wrap(err, "compareSlice on nameWords")
				return
			}
			if eq {
				tome, folio = lawyer.Tome, lawyer.Folio
				break
			}
		}
		return
	}
}
