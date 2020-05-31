package validate

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"

	"github.com/athomecomar/athome/backend/identifier/scrap"
	"github.com/athomecomar/semantic/semprov"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

var ByCategory = map[semprov.Category]licenseValidator{
	semprov.Psychologist: licensePsychologist,
}

type licenseValidator func(uint64, uint64) (bool, error)

func licensePsychologist(dni uint64, license uint64) (valid bool, err error) {
	const uri = "http://www.colpsiba.org.ar/autogestion/autogestion/"
	const formName = "F1"
	const licenseFormName = "idmatriculado"
	const dniFormName = "documentonro"

	formCollector := colly.NewCollector()
	formValues := make(url.Values)
	formCollector.OnHTML("form", scrap.GetFormValuesAsUrlValues(formName, formValues))
	err = formCollector.Visit(uri)
	if err != nil {
		return false, errors.Wrap(err, "Visit")
	}
	formValues.Set(dniFormName, strconv.Itoa(int(dni)))
	formValues.Set(licenseFormName, strconv.Itoa(int(license)))

	jar, err := cookiejar.New(nil)
	if err != nil {
		return false, errors.Wrap(err, "cookiejar.New")
	}
	URL, err := url.Parse(uri)
	if err != nil {
		return false, errors.Wrap(err, "url.Parse")
	}
	jar.SetCookies(URL, formCollector.Cookies(uri))
	cli := &http.Client{Jar: jar}
	resp, err := cli.PostForm(uri, formValues)
	if err != nil {
		return false, errors.Wrap(err, "PostForm")
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("PostForm returned %v code", resp.StatusCode)
		return
	}
	return resp.ContentLength < 50000, nil
}

//  func getFormValues(formName string, formValues map[string]string) colly.HTMLCallback {
// 	return func(e *colly.HTMLElement) {
// 		if e.Attr("name") != formName {
// 			return
// 		}
// 		names := e.ChildAttrs("input", "name")
// 		vals := e.ChildAttrs("input", "value")
// 		for i, n := range names {
// 			formValues[n] = vals[i]
// 		}
// }
// }
