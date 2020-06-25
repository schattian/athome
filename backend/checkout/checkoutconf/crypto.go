package checkoutconf

import "github.com/athomecomar/envconf"

func GetNUMBER_SECRET() (s []byte) {
	s = []byte(envconf.Get("NUMBER_SECRET", "passphrasewhichneedstobe32bytes!"))
	if isSilly(s) && envconf.NotInDevelopment() {
		panic("silly number_secret secret given")
	}
	return
}

func isSilly(x []byte) bool {
	return len(x) != 32
}
