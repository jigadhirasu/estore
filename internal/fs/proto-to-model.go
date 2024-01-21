package fs

import (
	"estore/internal/models/filemap"
	"estore/protoc/fspb"

	"github.com/jigadhirasu/rex"
)

func T1FileToFilestore(a *fspb.File) filemap.FileStore {
	return filemap.FileStore{
		ID:       a.Id,
		FileName: a.Name,
	}
}
func F1FileToFilestore(ctx rex.Context, a *fspb.File) (filemap.FileStore, error) {
	return T1FileToFilestore(a), nil
}
