package fs

import (
	"estore/protoc/fspb"

	"github.com/jigadhirasu/rex"
)

// 檔案資料夾階層結構攤平
func H1ExpandCreateFile(ctx rex.Context, dir *fspb.Dir) rex.Iterable[*fspb.File] {
	files := []rex.Iterable[*fspb.File]{}

	for _, file := range dir.Files {
		file.Name = dir.Name + "/" + file.Name
	}
	files = append(files, rex.From(dir.Files...))

	if len(dir.Dirs) > 0 {
		for _, subdir := range dir.Dirs {
			subdir.Name = dir.Name + "/" + subdir.Name

			files = append(files, H1ExpandCreateFile(ctx, subdir))
		}
	}

	return rex.Merge(files...)(rex.NewContext(ctx))
}
