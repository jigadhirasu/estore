package validate

import (
	"errors"

	"github.com/jigadhirasu/rex"
)

type IRankId interface {
	GetRankId() string
}

func IsRankIdValid[A IRankId](ctx rex.Context, a A) error {
	if a.GetRankId() == "" {
		return errors.New("rank_id is empty")
	}

	return nil
}
