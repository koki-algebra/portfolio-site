package converter

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	usermodel "backend/internal/domain/model/user"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func TestConvertUser(t *testing.T) {
	assert.Nil(t, ConvertUser(nil))

	userID := uuid.New()
	authID := uuid.NewString()
	email := "example@example.com"

	user := &usermodel.User{
		UserID: userID,
		AuthID: authID,
		Email:  email,
	}

	got := ConvertUser(user)

	want := &userv1.User{
		UserId: userID.String(),
		AuthId: authID,
		Email:  email,
	}

	assert.Equal(t, want, got)
}
