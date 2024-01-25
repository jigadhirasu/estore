package rank

import (
	"errors"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
)

func IsConditionValid(ctx rex.Context, a *clubpb.Rank) error {
	if len(a.Condition) < 1 {
		return errors.New("condition less than 1")
	}

	return nil
}
