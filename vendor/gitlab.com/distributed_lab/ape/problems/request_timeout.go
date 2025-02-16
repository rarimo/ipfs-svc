package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func RequestTimeout() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusRequestTimeout),
		Status: fmt.Sprintf("%d", http.StatusRequestTimeout),
	}
}
