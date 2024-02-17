package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Xploitable/todo-app"
	"github.com/Xploitable/todo-app/package/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const TOKEN_TTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
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
	salt := viper.GetString("app.salt")
	_hash.Write([]byte(password))
	return fmt.Sprintf("%x", _hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	signin_key := viper.GetString("app.signinkey")
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signin_key))
}
