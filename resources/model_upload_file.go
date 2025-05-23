/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UploadFile struct {
	Key
	Attributes UploadFileAttributes `json:"attributes"`
}
type UploadFileRequest struct {
	Data     UploadFile `json:"data"`
	Included Included   `json:"included"`
}

type UploadFileListRequest struct {
	Data     []UploadFile    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UploadFileListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UploadFileListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUploadFile - returns UploadFile from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUploadFile(key Key) *UploadFile {
	var uploadFile UploadFile
	if c.tryFindEntry(key, &uploadFile) {
		return &uploadFile
	}
	return nil
}
