package fs

import (
	"estore/internal/models/filemap"
	"estore/protoc/fspb"
)

func T1FileToFilestore(a *fspb.File) filemap.FileStore {
	return filemap.FileStore{
		ID:       a.Id,
		FileName: a.Name,
	}
}
