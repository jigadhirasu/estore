package fs

import (
	"estore/pkg/module/logger"
	"estore/protoc/fspb"
	"os"

	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FileInfo struct {
	Name           string
	Size           int32
	CreationTime   int64
	LastAccessTime int64
	LastWriteTime  int64
}

func (info FileInfo) Attach(fileOrDir any) {
	f := func(a int64) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{Seconds: a / 1e9, Nanos: int32(a % 1e9)}
	}

	switch v := fileOrDir.(type) {
	case *fspb.Dir:
		v.CreationTime = f(info.CreationTime)
		v.LastAccessTime = f(info.LastAccessTime)
		v.LastWriteTime = f(info.LastWriteTime)
	case *fspb.File:
		v.Size = info.Size
		v.CreationTime = f(info.CreationTime)
		v.LastAccessTime = f(info.LastAccessTime)
		v.LastWriteTime = f(info.LastWriteTime)
	}

}

// 讀取當前目錄資訊
func F1ReadRootDirInfo(ctx rex.Context, dir *fspb.Dir) (*fspb.Dir, error) {
	info, err := F1ReadFileInfo(ctx, dir.Name)
	if err != nil {
		logger.Out(ctx).Error("ReadRootDirInfo", zap.Error(err))
		return nil, err
	}
	info.Attach(dir)

	return dir, nil
}

// 讀取當前目錄下的資訊
func F1ReadNextDirInfo(ctx rex.Context, upDir *fspb.Dir) (*fspb.Dir, error) {

	files, err := os.ReadDir(upDir.Name)
	if err != nil {
		logger.Out(ctx).Error("ReadNextDirInfo", zap.Error(err))
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {

			dir := &fspb.Dir{Name: upDir.Name + "/" + file.Name()}

			info, err := F1ReadFileInfo(ctx, dir.Name)
			if err != nil {
				logger.Out(ctx).Error("ReadNextDirInfo", zap.Error(err))
				return nil, err
			}
			info.Attach(dir)

			upDir.Dirs = append(upDir.Dirs, dir)
			continue
		}

		f := &fspb.File{Name: upDir.Name + "/" + file.Name()}
		info, err := F1ReadFileInfo(ctx, f.Name)
		if err != nil {
			logger.Out(ctx).Error("ReadNextDirInfo", zap.Error(err))
			return nil, err
		}
		info.Attach(f)
		upDir.Files = append(upDir.Files, f)
	}

	return upDir, nil
}
