package validate

import (
	"errors"

	"github.com/jigadhirasu/rex"
)

type IId interface {
	GetId() string
}

func IsIdValid[A IId](ctx rex.Context, a A) error {
	if a.GetId() == "" {
		return errors.New("id is empty")
	}

	return nil
}
