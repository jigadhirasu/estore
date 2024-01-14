package transfer1

import (
	"estore/internal/models/filemap"
	"estore/protoc/fspb"
)

func FileToFilestore(a *fspb.File) filemap.FileStore {
	return filemap.FileStore{
		ID:       a.Id,
		FileName: a.Name,
	}
}
