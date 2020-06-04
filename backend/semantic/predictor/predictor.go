package predictor

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Predictor struct {
	Ctx context.Context
	http.Client
}

func NewPredictor(ctx context.Context) *Predictor {
	p := &Predictor{Ctx: ctx}
	p.Client = *http.DefaultClient
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	p.Client.Transport = tr
	return p
}

// RouteTo retrieves a route given a path alias to the desired resource
// Notice: it returns a trailing slash on the return value in case it can contain childs.
// For example, in the case of /items, it'll return /items/ instead (alerting it is a sort of dir of sub-nodes)
// Path can be "auth", "product", "category_predict", "category", "category_attributes"
func (p *Predictor) RouteTo(path string, params url.Values, ids ...interface{}) (string, error) {
	base := "https://api.mercadolibre.com"
	if ids != nil {
		base += fmt.Sprintf(path, ids...)
	} else {
		base += strings.ReplaceAll(path, "%v", "")
	}
	URL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	URL.RawQuery = params.Encode()
	base = URL.String()
	return base, nil
}

func (ml *Predictor) Get(url string) (resp *http.Response, err error) {

	req, err := http.NewRequestWithContext(ml.Ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return ml.Do(req)
}

type Prediction struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (ml *Predictor) Predict(title string) (*Prediction, error) {
	params := url.Values{}
	params.Set("q", title)
	params.Set("limit", "1")
	URL, err := ml.RouteTo("/sites/MLA/domain_discovery/search", params)
	if err != nil {
		return nil, errors.Wrap(err, "RouteTo")
	}
	resp, err := ml.Get(URL)
	if err != nil {
		return nil, errors.Wrap(err, "Get")
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, errors.Wrap(errFromReader(resp.Body), "errFromReader")
	}
	var predictions []*Prediction
	err = json.NewDecoder(resp.Body).Decode(&predictions)
	if err != nil {
		return nil, errors.Wrap(err, "json.Decode")
	}
	if len(predictions) != 1 {
		return nil, errors.Wrap(err, "predictions length != 1")
	}
	return predictions[0], nil
}
