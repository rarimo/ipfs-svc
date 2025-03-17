package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/ipfs-svc/resources"
)

func NewUploadJSON(r *http.Request) (req resources.UploadJsonRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	reqData := req.Data.Attributes
	err = validation.Errors{
		"data/type":            validation.Validate(req.Data.Type, validation.Required, validation.In(resources.JSON)),
		"data/attributes/salt": validation.Validate(reqData.Metadata, validation.Required),
	}.Filter()

	return req, err
}

func newDecodeError(what string, err error) error {
	return validation.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}
