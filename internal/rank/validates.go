package rank

import (
	"errors"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
)

func F0CheckNameIsValid(ctx rex.Context, a *clubpb.Rank) error {
	if a.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

func F0CheckIdIsValid(ctx rex.Context, a *clubpb.Rank) error {
	if a.Id == "" {
		return errors.New("uuid is empty")
	}

	return nil
}

func F0CheckConditionIsValid(ctx rex.Context, a *clubpb.Rank) error {
	if len(a.Condition) < 1 {
		return errors.New("condition less than 1")
	}

	return nil
}
