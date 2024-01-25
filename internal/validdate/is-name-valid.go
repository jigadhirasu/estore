package validate

import (
	"errors"

	"github.com/jigadhirasu/rex"
)

type IName interface {
	GetName() string
}

func IsNameValid[A IName](ctx rex.Context, a A) error {
	if a.GetName() == "" {
		return errors.New("name is empty")
	}

	return nil
}
