package scraper

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/identifierconf"
	"github.com/athomecomar/semantic/semprov"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

var InferrorByFullnameByCategory = map[semprov.Category]inferrorByFullnameByCategory{
	semprov.Medic: inferrorByFullnameByCategoryMedic,
}

type inferrorByFullnameByCategory func(afero.Fs, string, string) (uint64, error)

const medicLicensesInferencesFilename = "medic_licenses_by_fullname.json"

func inferrorByFullnameByCategoryMedic(fs afero.Fs, name, surname string) (uint64, error) {
	f, err := fs.Open(identifierconf.GetDATA_DIR() + "/" + medicLicensesInferencesFilename)
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
		if len(givenSurnameWords)+len(givenNameWords) > len(words) {
			continue
		}
		surnameWords := words[0:len(givenSurnameWords)]
		nameWords := words[len(givenSurnameWords) : len(givenNameWords)+len(givenSurnameWords)]

		eq, err := compareSlice(surnameWords, givenSurnameWords)
		if err != nil {
			return 0, errors.Wrap(err, "compareSlice on surnameWords")
		}
		if !eq {
			continue
		}
		log.Printf("comparing %v with %v", nameWords, givenNameWords)
		eq, err = compareSlice(nameWords, givenNameWords)
		if err != nil {
			return 0, errors.Wrap(err, "compareSlice on nameWords")
		}
		log.Println(eq)
		if eq {
			match = license
		}
	}
	return match, nil
}

type medicInferencesByFullname map[string]uint64
