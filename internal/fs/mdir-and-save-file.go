package fs

import (
	"estore/pkg/module/logger"
	"estore/protoc/fspb"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/jigadhirasu/rex"
	"github.com/vincent-petithory/dataurl"
	"go.uber.org/zap"
)

// 創建資料夾
// 產生uuid
// 產生完整檔案名稱
// 寫入檔案
func F1MkdirAndSaveFile(ctx rex.Context, file *fspb.File) (*fspb.File, error) {

	du, err := dataurl.DecodeString(file.Dataurl)
	if err != nil {
		logger.Out(ctx).Error("MkdirAndSaveFile", zap.Error(err))
		return nil, err
	}
	file.Dataurl = ""

	idx := strings.LastIndex(file.Name, "/")
	if err := os.MkdirAll(file.Name[:idx], 0755); err != nil {
		logger.Out(ctx).Error("MkdirAndSaveFile", zap.Error(err))
		return nil, err
	}

	if file.Id == "" {
		file.Id = uuid.NewString()
	}

	if len(file.Name) == idx+1 {
		file.Name = file.Name + file.Id
	}

	file.Name = file.Name + "." + du.MediaType.Subtype

	if err := os.WriteFile(file.Name, du.Data, 0755); err != nil {
		logger.Out(ctx).Error("MkdirAndSaveFile", zap.Error(err))
		return nil, err
	}

	return file, nil

}
