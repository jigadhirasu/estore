package fs

import (
	"context"
	"estore/protoc/fspb"
	"slices"
	"testing"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"github.com/stretchr/testify/assert"
)

func TestExpandDeleteDir1(t *testing.T) {

	rootDir := &fspb.Dir{
		Name: "/PV/src/file",
		Dirs: []*fspb.Dir{
			{
				Name: "A",
				Files: []*fspb.File{
					{
						Name:    "file1",
						Dataurl: "dataurl1",
					},
				},
			},
			{
				Name: "B",
			},
		},
	}

	ctx := rex.NewContext(context.TODO())

	result, err := H1ExpandDeleteDir(ctx, rootDir).ToSlice()

	assert.NoError(t, err)

	dirnames := array.Map[*fspb.Dir, string](func(a *fspb.Dir) string {
		return a.Name
	})(result)
	slices.Sort(dirnames)

	assert.Equal(t, []string{"/PV/src/file/A", "/PV/src/file/B"}, dirnames)

}

func TestExpandDeleteDir2(t *testing.T) {

	rootDir := &fspb.Dir{
		Name: "/PV/src/file",
		Dirs: []*fspb.Dir{
			{
				Name: "A",
				Files: []*fspb.File{
					{
						Name:    "file1",
						Dataurl: "dataurl1",
					},
				},
			},
			{
				Name: "B",
			},
		},
		Files: []*fspb.File{
			{Name: "file1.txt"},
		},
	}

	ctx := rex.NewContext(context.TODO())

	result, err := H1ExpandDeleteDir(ctx, rootDir).ToSlice()

	assert.NoError(t, err)

	dirnames := array.Map[*fspb.Dir, string](func(a *fspb.Dir) string {
		return a.Name
	})(result)
	slices.Sort(dirnames)

	assert.Equal(t, []string{"/PV/src/file", "/PV/src/file/A", "/PV/src/file/B"}, dirnames)

}
