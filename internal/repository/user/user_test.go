package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	domain "art-sso/internal/domain/user"
	repository "art-sso/internal/repository/user"
)

func TestMySQLUserRepository(t *testing.T) {
	// MySQL Repository 테스트이지만, 일단은 SQLite를 사용한다. 현재 메소드는 모두 서로 호환되는 상황.
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		t.Fatal(err)
	}

	userRepo := repository.NewMySQLUserRepository(db)

	t.Run("Create user", func(t *testing.T) {
		user := domain.User{
			Email:         "user@example.com",
			Salt:          "randomsalt",
			EmailVerified: false,
		}

		err := userRepo.CreateUser(&user)

		assert.NoError(t, err)
	})

	t.Run("Get user by ID", func(t *testing.T) {
		user, err := userRepo.GetUserByID(1)
		t.Log(user)

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Get user by email", func(t *testing.T) {
		user, err := userRepo.GetUserByEmail("user@example.com")

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Update user", func(t *testing.T) {
		user, err := userRepo.GetUserByID(1)

		assert.NoError(t, err)

		user.EmailVerified = true
		err = userRepo.UpdateUser(user)

		assert.NoError(t, err)
	})

	t.Run("Delete user", func(t *testing.T) {
		user, err := userRepo.GetUserByID(1)

		assert.NoError(t, err)

		err = userRepo.DeleteUser(user)

		assert.NoError(t, err)
	})
}