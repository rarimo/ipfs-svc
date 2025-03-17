package requests

import (
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/ipfs-svc/resources"
)

func NewUploadFile(r *http.Request) (req resources.UploadFileRequest, err error) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		return req, fmt.Errorf("content-type must be multipart/form-data")
	}

	req.Data.Type = resources.FILE

	err = validation.Errors{
		"data/type": validation.Validate(req.Data.Type, validation.Required, validation.In(resources.FILE)),
	}.Filter()

	return req, err
}
