package v1

import (
	"context"
	"estore/internal/fs"
	"estore/internal/mariadb"
	"estore/protoc/fspb"

	"github.com/jigadhirasu/rex"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm/clause"
)

type FSService struct {
	fspb.UnimplementedFileSystemServer
}

func (FSService) Ping(cin context.Context, req *fspb.PingPong) (*fspb.PingPong, error) {
	req.Message = "PONG"
	return req, nil
}

func (FSService) CreateDirFile(cin context.Context, req *fspb.CreateOrUpdateDirFileRequest) (*fspb.CreateOrUpdateDirFileResponse, error) {

	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*fspb.Dir](dbname)),
		rex.Once(mariadb.F0Begin[*fspb.Dir](dbname)),
		fs.PipeCreateFile(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*fspb.CreateOrUpdateDirFileResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()
}

func (FSService) UpdateDirFile(cin context.Context, req *fspb.CreateOrUpdateDirFileRequest) (*fspb.CreateOrUpdateDirFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDirFile not implemented")
}
func (FSService) CreateOrUpdateDirFile(cin context.Context, req *fspb.CreateOrUpdateDirFileRequest) (*fspb.CreateOrUpdateDirFileResponse, error) {

	dbname := "test"

	ctx := rex.NewContext(cin)

	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "file_name"}},
		DoNothing: true,
	}

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*fspb.Dir](dbname)),
		rex.Once(mariadb.F0Begin[*fspb.Dir](dbname)),
		fs.PipeCreateFile(mariadb.TxKey(dbname), onConflict),
		rex.Tap(mariadb.F0Commit[*fspb.CreateOrUpdateDirFileResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)
	item := <-pipe()

	return item()
}

func (FSService) DeleteDirFile(cin context.Context, req *fspb.DeleteDirFileRequest) (*fspb.DeleteDirFileResponse, error) {

	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*fspb.Dir](dbname)),
		rex.Once(mariadb.F0Begin[*fspb.Dir](dbname)),
		fs.PipeDeleteFile(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*fspb.DeleteDirFileResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)
	item := <-pipe()

	return item()
}

func (FSService) FindDirFile(cin context.Context, req *fspb.FindDirFileRequest) (*fspb.FindDirFileResponse, error) {

	ctx := rex.NewContext(cin)

	pipe := fs.PipeFindFile()(rex.From(req.RootDir))(ctx)

	item := <-pipe()

	return item()
}
