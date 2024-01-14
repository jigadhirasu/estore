package fs

import (
	"estore/protoc/fspb"
	"os"
	"path/filepath"
	"strings"

	"github.com/jigadhirasu/rex"
)

func F1ResourceRoot(ctx rex.Context, dir *fspb.Dir) (*fspb.Dir, error) {
	root := os.Getenv("RESOURCE_ROOT")
	if root == "" {
		return dir, nil
	}

	dir.Name = filepath.Clean(root + "/" + dir.Name)
	dir.Name = strings.ReplaceAll(dir.Name, "\\", "/")
	return dir, nil
}
