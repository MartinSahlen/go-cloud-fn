package shimV2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ObjectOwner struct {
	Entity   string `json:"entity,omitempty"`
	EntityId string `json:"entityId,omitempty"`
}

type Object struct {
	Acl                []*ObjectAccessControl `json:"acl,omitempty"`
	Bucket             string                 `json:"bucket,omitempty"`
	CacheControl       string                 `json:"cacheControl,omitempty"`
	ComponentCount     int64                  `json:"componentCount,omitempty"`
	ContentDisposition string                 `json:"contentDisposition,omitempty"`
	ContentEncoding    string                 `json:"contentEncoding,omitempty"`
	ContentLanguage    string                 `json:"contentLanguage,omitempty"`
	ContentType        string                 `json:"contentType,omitempty"`
	Crc32c             string                 `json:"crc32c,omitempty"`
	Etag               string                 `json:"etag,omitempty"`
	Generation         int64                  `json:"generation,omitempty,string"`
	Id                 string                 `json:"id,omitempty"`
	Kind               string                 `json:"kind,omitempty"`
	Md5Hash            string                 `json:"md5Hash,omitempty"`
	MediaLink          string                 `json:"mediaLink,omitempty"`
	Metadata           map[string]string      `json:"metadata,omitempty"`
	Metageneration     int64                  `json:"metageneration,omitempty,string"`
	Name               string                 `json:"name,omitempty"`
	Owner              *ObjectOwner           `json:"owner,omitempty"`
	SelfLink           string                 `json:"selfLink,omitempty"`
	Size               uint64                 `json:"size,omitempty,string"`
	StorageClass       string                 `json:"storageClass,omitempty"`
	TimeDeleted        string                 `json:"timeDeleted,omitempty"`
	Updated            string                 `json:"updated,omitempty"`
}

type ObjectAccessControl struct {
	Bucket     string `json:"bucket,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Email      string `json:"email,omitempty"`
	Entity     string `json:"entity,omitempty"`
	EntityId   string `json:"entityId,omitempty"`
	Etag       string `json:"etag,omitempty"`
	Generation int64  `json:"generation,omitempty,string"`
	Id         string `json:"id,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Object     string `json:"object,omitempty"`
	Role       string `json:"role,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
}

type BucketHandlerFunc func(object Object)

func HandleBucketEvent(h BucketHandlerFunc) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var object Object
	err = json.Unmarshal(stdin, &object)
	if err != nil {
		log.Fatal(err)
	}
	h(object)
}
