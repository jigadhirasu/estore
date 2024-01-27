package validate

import (
	"errors"

	"github.com/jigadhirasu/rex"
)

type IClubId interface {
	GetClubId() string
}

func IsClubIdValid[A IClubId](ctx rex.Context, a A) error {
	if a.GetClubId() == "" {
		return errors.New("club_id is empty")
	}

	return nil
}
