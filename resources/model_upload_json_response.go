/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UploadJsonResponse struct {
	Key
	Attributes UploadJsonResponseAttributes `json:"attributes"`
}
type UploadJsonResponseResponse struct {
	Data     UploadJsonResponse `json:"data"`
	Included Included           `json:"included"`
}

type UploadJsonResponseListResponse struct {
	Data     []UploadJsonResponse `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *UploadJsonResponseListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UploadJsonResponseListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUploadJsonResponse - returns UploadJsonResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUploadJsonResponse(key Key) *UploadJsonResponse {
	var uploadJSONResponse UploadJsonResponse
	if c.tryFindEntry(key, &uploadJSONResponse) {
		return &uploadJSONResponse
	}
	return nil
}
