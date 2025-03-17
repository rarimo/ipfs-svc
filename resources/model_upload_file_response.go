/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UploadFileResponse struct {
	Key
	Attributes UploadFileResponseAttributes `json:"attributes"`
}
type UploadFileResponseResponse struct {
	Data     UploadFileResponse `json:"data"`
	Included Included           `json:"included"`
}

type UploadFileResponseListResponse struct {
	Data     []UploadFileResponse `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *UploadFileResponseListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UploadFileResponseListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUploadFileResponse - returns UploadFileResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUploadFileResponse(key Key) *UploadFileResponse {
	var uploadFileResponse UploadFileResponse
	if c.tryFindEntry(key, &uploadFileResponse) {
		return &uploadFileResponse
	}
	return nil
}
