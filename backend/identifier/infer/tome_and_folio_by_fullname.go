package infer

import "github.com/athomecomar/semantic/semprov"

var TomeAndFolioByFullnameByCategory = map[*semprov.Category]tomeAndFolioByFullnameByCategory{
	semprov.Lawyer:   tomeAndFolioByFullnameAttorneyAndLawyers(semprov.Lawyer),
	semprov.Attorney: tomeAndFolioByFullnameAttorneyAndLawyers(semprov.Attorney),
}
