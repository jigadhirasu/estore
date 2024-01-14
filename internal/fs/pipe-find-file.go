package fs

import (
	"estore/protoc/fspb"

	"github.com/jigadhirasu/rex"
)

func PipeFindFile() rex.PipeLine[*fspb.Dir, *fspb.FindDirFileResponse] {
	return rex.Pipe4(
		rex.Map(F1ResourceRoot),
		rex.Map(F1ReadRootDirInfo),
		rex.Map(F1ReadNextDirInfo),
		rex.Map[*fspb.Dir, *fspb.FindDirFileResponse](func(ctx rex.Context, a *fspb.Dir) (*fspb.FindDirFileResponse, error) {
			return &fspb.FindDirFileResponse{
				RootDir: a,
			}, nil
		}),
	)
}
