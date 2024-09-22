package caller

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"backend/internal/domain/model"
)

type Caller interface {
	GetUserID() uuid.UUID
	GetEmail() string
}

type callerImpl struct {
	userID uuid.UUID
	email  string
}

func NewCaller(ctx context.Context) (Caller, error) {
	c := new(callerImpl)

	user, ok := model.UserFromContext(ctx)
	if !ok {
		return nil, errors.New("empty user")
	}

	c.userID = user.UserID
	c.email = user.Email

	return c, nil
}

func (c *callerImpl) GetUserID() uuid.UUID {
	return c.userID
}

func (c *callerImpl) GetEmail() string {
	return c.email
}
