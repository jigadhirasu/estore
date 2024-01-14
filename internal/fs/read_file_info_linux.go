//go:build linux

package fs

import (
	"errors"
	"estore/pkg/module/logger"
	"os"
	"syscall"

	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
)

func F1ReadFileInfo(ctx rex.Context, name string) (FileInfo, error) {

	var fileinfo FileInfo

	info, err := os.Stat(name)
	if err != nil {
		logger.Out(ctx).Error("readFileInfo", zap.Error(err))
		return fileinfo, err
	}

	fileinfo.Name = name
	fileinfo.Size = int32(info.Size())

	sysfile, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		logger.Out(ctx).Error("readFileInfo", zap.Error(errors.New("not a Stat_t")))
		return fileinfo, nil
	}

	fileinfo.CreationTime = sysfile.Ctim.Sec*1e9 + sysfile.Ctim.Nsec
	fileinfo.LastAccessTime = sysfile.Atim.Sec*1e9 + sysfile.Atim.Nsec
	fileinfo.LastWriteTime = sysfile.Mtim.Sec*1e9 + sysfile.Mtim.Nsec

	return fileinfo, nil
}
