package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Xploitable/todo-app"
	"github.com/Xploitable/todo-app/package/repository"
)

const SALT = "00sdfk12ppxcmlDSF=qd4"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	_hash := sha1.New()
	_hash.Write([]byte(password))
	return fmt.Sprintf("%x", _hash.Sum([]byte(SALT)))
}
