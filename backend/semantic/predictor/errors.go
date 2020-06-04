package predictor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	ErrRemoteInconsistency = errors.New("the SERVER had an inconsistency while performing a request (status code != real behaviour)")
)

type Error struct {
	Message     string      `json:"message,omitempty"`
	ResponseErr string      `json:"error,omitempty"`
	Status      int         `json:"status,omitempty"`
	Cause       []*errCause `json:"cause,omitempty"`
}

type errCause struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (svErr *Error) Error() string {
	if len(svErr.Cause) == 0 {
		return fmt.Sprintf("%s: %s", svErr.ResponseErr, svErr.Message)
	}
	var strErr string
	for _, cause := range svErr.Cause {
		strErr += fmt.Sprintf("%s: %s", cause.Code, cause.Message)
		strErr += "; "
	}
	return strErr
}

func errFromReader(stream io.Reader) error {
	body := &Error{}
	err := json.NewDecoder(stream).Decode(body)
	if err != nil {
		return err
	}
	if body.ResponseErr == "" && body.Message == "" {
		body.ResponseErr = "remote_inconsistency"
		body.Message = ErrRemoteInconsistency.Error()
	}
	return body
}
