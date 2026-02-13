package user

import (
	"fmt"

	"github.com/orgmange/order-service/internal/util"
)

type Model struct {
	ID    uint
	Name  string
	Email string
}

func (u *Model) CheckValid() error {
	if !util.NAME_REGEX.MatchString(u.Name) {
		return fmt.Errorf("not valid name")
	}

	if !util.EMAIL_REGEX.MatchString(u.Email) {
		return fmt.Errorf("not valid email")
	}

	return nil
}
