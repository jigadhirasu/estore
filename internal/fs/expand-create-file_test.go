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

func TestExpandCreateFile(t *testing.T) {

	rootDir := &fspb.Dir{
		Name: "/PV/src",
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
		},
		Files: []*fspb.File{
			{
				Dataurl: "dataurl2",
			},
			{
				Name:    "file3",
				Dataurl: "dataurl3",
			},
		},
	}

	ctx := rex.NewContext(context.TODO())

	result, err := H1ExpandCreateFile(ctx, rootDir).ToSlice()

	assert.NoError(t, err)

	dataurls := array.Map[*fspb.File, string](func(a *fspb.File) string {
		return a.Dataurl
	})(result)
	slices.Sort(dataurls)
	assert.Equal(t, []string{"dataurl1", "dataurl2", "dataurl3"}, dataurls)

	names := array.Map[*fspb.File, string](func(a *fspb.File) string {
		return a.Name
	})(result)
	slices.Sort(names)
	assert.Equal(t, []string{"/PV/src/", "/PV/src/A/file1", "/PV/src/file3"}, names)

}
