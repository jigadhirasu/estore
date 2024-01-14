package fs

import (
	"estore/protoc/fspb"

	"github.com/jigadhirasu/rex"
)

// 把刪除需要處理的dir抓出來
func H1ExpandDeleteDir(ctx rex.Context, upDir *fspb.Dir) rex.Iterable[*fspb.Dir] {
	iterables := []rex.Iterable[*fspb.Dir]{}

	// 每個下層的目錄獨立送出
	for _, dir := range upDir.Dirs {
		dir.Name = upDir.Name + "/" + dir.Name
		iterables = append(iterables, H1ExpandDeleteDir(ctx, dir))
	}

	// 如果有檔案就是要刪除這個目錄下的檔案
	if len(upDir.Files) > 0 {
		iterables = append(iterables, rex.From(upDir))
	}

	// 下面都沒東西就是要刪除這個目錄
	if len(upDir.Dirs)+len(upDir.Files) == 0 {
		iterables = append(iterables, rex.From(upDir))
	}

	return rex.Merge(iterables...)(ctx)
}
