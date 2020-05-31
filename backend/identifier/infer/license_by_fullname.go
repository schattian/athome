package infer

import (
	"github.com/athomecomar/semantic/semprov"
	"github.com/spf13/afero"
)

var LicenseByFullnameByCategory = map[semprov.Category]licenseByFullnameByCategory{
	semprov.Medic: licensebyFullnameByCategoryMedic,
}

var ByFullnameFilenames = map[semprov.Category]string{
	semprov.Medic:    "medic_licenses_by_fullname.json",
	semprov.Lawyer:   "lawyers.json",
	semprov.Attorney: "attorneys.json",
}

type licenseByFullnameByCategory func(afero.Fs, string, string) (uint64, error)

type tomeAndFolioByFullnameByCategory func(afero.Fs, string, string) (uint64, uint64, error)
