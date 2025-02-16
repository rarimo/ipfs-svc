/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UploadJson struct {
	Key
	Attributes UploadJsonAttributes `json:"attributes"`
}
type UploadJsonRequest struct {
	Data     UploadJson `json:"data"`
	Included Included   `json:"included"`
}

type UploadJsonListRequest struct {
	Data     []UploadJson    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UploadJsonListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UploadJsonListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUploadJson - returns UploadJson from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUploadJson(key Key) *UploadJson {
	var uploadJSON UploadJson
	if c.tryFindEntry(key, &uploadJSON) {
		return &uploadJSON
	}
	return nil
}
