package infer

import "github.com/athomecomar/semantic/semprov"

var TomeAndFolioByFullnameByCategory = map[semprov.Category]tomeAndFolioByFullnameByCategory{
	semprov.Lawyer: tomeAndFolioByFullnameByCategoryLawyer,
}
