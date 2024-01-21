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

func F0CheckIdIsValid(ctx rex.Context, a *clubpb.Club) error {
	if a.Id == "" {
		return errors.New("club uuid is empty")
	}

	return nil
}
