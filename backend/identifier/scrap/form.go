package scrap

import (
	"net/url"

	"github.com/gocolly/colly"
)

func GetFormValuesAsUrlValues(formName string, formValues url.Values) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		if e.Attr("name") != formName {
			return
		}
		names := e.ChildAttrs("input", "name")
		vals := e.ChildAttrs("input", "value")
		for i, n := range names {
			formValues.Set(n, vals[i])
		}
	}
}
