package v1

import (
	"context"
	"estore/assets"
	"estore/configs"
	"estore/internal/fs"
	"estore/internal/mariadb"
	"estore/protoc/fspb"
	"fmt"
	"testing"
	"time"

	_ "embed"

	"github.com/jigadhirasu/rex"
	"github.com/stretchr/testify/assert"
)

func init() {
	configs.Develop()
}

var timestamp = time.Now().Unix()

// var timestamp = 123

func testCreateFileData(t *testing.T) []*fspb.Dir {
	return []*fspb.Dir{
		{
			Name: "create_A",
			Files: []*fspb.File{
				{
					Dataurl: string(assets.HtmlBase64),
				},
			},
		},
		{
			Name: "create_B",
			Files: []*fspb.File{
				{
					Name:    fmt.Sprintf("b-%d", timestamp),
					Dataurl: string(assets.JpegBase64),
				},
			},
			Dirs: []*fspb.Dir{
				{
					Name: "create_C",
					Files: []*fspb.File{
						{
							Name:    fmt.Sprintf("c-%d", timestamp),
							Dataurl: string(assets.HtmlBase64),
						},
					},
				},
			},
		},
	}
}

func TestCreateFile(t *testing.T) {

	dbname := "test"

	ctx := rex.NewContext(context.TODO())

	dirs := testCreateFileData(t)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*fspb.Dir](dbname)),
		rex.Once(mariadb.F0Begin[*fspb.Dir](dbname)),
		fs.PipeCreateFile(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*fspb.CreateOrUpdateDirFileResponse](dbname)),
	)(rex.From(dirs...))(ctx)

	defer mariadb.Rollback(ctx, dbname)
	item := <-pipe()
	resp1, err := item()

	assert.NoError(t, err)

	if err != nil {
		return
	}

	assert.Equal(t, int32(3), resp1.Response.AffectedRows)
}

func TestFindFile(t *testing.T) {

	ctx := rex.NewContext(context.TODO())

	pipe := fs.PipeFindFile()(rex.From(&fspb.Dir{
		Name: "create_A",
	}))(ctx)

	item := <-pipe()

	resp, err := item()

	assert.NoError(t, err)

	assert.LessOrEqual(t, 1, len(resp.RootDir.Files))
}

func TestDeleteFile(t *testing.T) {

	dbname := "test"

	ctx := rex.NewContext(context.TODO())

	dirs := testCreateFileData(t)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*fspb.Dir](dbname)),
		rex.Once(mariadb.F0Begin[*fspb.Dir](dbname)),
		fs.PipeDeleteFile(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*fspb.DeleteDirFileResponse](dbname)),
	)(rex.From(dirs...))(ctx)

	defer mariadb.Rollback(ctx, dbname)
	item := <-pipe()
	resp1, err := item()

	assert.NoError(t, err)

	if err != nil {
		return
	}

	assert.Equal(t, int32(3), resp1.Response.AffectedRows)
}
