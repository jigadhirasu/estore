package fs

import (
	"context"
	"estore/assets"
	"estore/protoc/fspb"
	"slices"
	"testing"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"github.com/stretchr/testify/assert"
)

func TestMkdirAndSaveFile(t *testing.T) {

	files := []*fspb.File{
		{
			Name:    "/PV/src/B/index",
			Dataurl: string(assets.HtmlBase64),
		},
		{
			Name:    "/PV/src/A",
			Dataurl: string(assets.JpegBase64),
		},
	}

	pipe := rex.Pipe1(
		rex.Map(F1MkdirAndSaveFile),
	)(rex.From(files...))

	ctx := rex.NewContext(context.Background())
	result, err := pipe(ctx).ToSlice()

	assert.NoError(t, err)

	names := array.Map[*fspb.File, string](func(a *fspb.File) string {
		return a.Name
	})(result)
	slices.Sort(names)

	assert.Equal(t, []string{"/PV/src/A.jpeg", "/PV/src/B/index.html"}, names)

}
