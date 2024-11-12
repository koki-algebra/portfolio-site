package callermodel

import (
	"github.com/google/uuid"
)

type Caller struct {
	userID uuid.UUID
	authID string
	email  string
}

func (c *Caller) GetUserID() uuid.UUID {
	return c.userID
}

func (c *Caller) GetAuthID() string {
	return c.authID
}

func (c *Caller) GetEmail() string {
	return c.email
}
