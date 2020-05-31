package infer

import (
	"encoding/json"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/identifierconf"
	"github.com/athomecomar/athome/backend/identifier/normalise"
	"github.com/athomecomar/semantic/semprov"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type medicInferencesByFullname map[string]uint64

func licensebyFullnameByCategoryMedic(fs afero.Fs, name, surname string) (uint64, error) {
	f, err := fs.Open(identifierconf.GetDATA_DIR() + "/" + ByFullnameFilenames[semprov.Medic])
	if err != nil {
		return 0, errors.Wrap(err, "fs.Open")
	}
	licenseByName := make(medicInferencesByFullname)
	err = json.NewDecoder(f).Decode(&licenseByName)
	if err != nil {
		return 0, errors.Wrap(err, "json.Decode")
	}

	givenSurnameWords, givenNameWords := strings.Split(surname, " "), strings.Split(name, " ")
	var match uint64
	for fullname, license := range licenseByName {
		words := strings.Split(fullname, " ")
		surnameWords := words[0:len(givenSurnameWords)]
		nameWords := words[len(givenSurnameWords):]

		eq, err := normalise.CompareSlice(surnameWords, givenSurnameWords)
		if err != nil {
			return 0, errors.Wrap(err, "compareSlice on surnameWords")
		}
		if !eq {
			continue
		}
		eq, err = normalise.CompareSliceSoft(nameWords, givenNameWords)
		if err != nil {
			return 0, errors.Wrap(err, "compareSlice on nameWords")
		}
		if eq {
			match = license
			break
		}
	}
	return match, nil
}
