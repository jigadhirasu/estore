package club

import (
	"errors"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
)

func F0CheckNameIsValid(ctx rex.Context, a *clubpb.Club) error {
	if a.Name == "" {
		return errors.New("club name is empty")
	}

	return nil
}
