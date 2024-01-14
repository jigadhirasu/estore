//go:build windows

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

	sysfile, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		logger.Out(ctx).Error("readFileInfo", zap.Error(errors.New("not a Win32FileAttributeData")))
		return fileinfo, nil
	}

	fileinfo.CreationTime = sysfile.CreationTime.Nanoseconds()
	fileinfo.LastAccessTime = sysfile.LastAccessTime.Nanoseconds()
	fileinfo.LastWriteTime = sysfile.LastWriteTime.Nanoseconds()

	return fileinfo, nil
}
