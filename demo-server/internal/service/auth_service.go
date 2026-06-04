package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/houmingqi/lims-demo-server/internal/dto"
	"github.com/houmingqi/lims-demo-server/internal/pkg/apperror"
	"github.com/houmingqi/lims-demo-server/internal/pkg/auth"
	"github.com/houmingqi/lims-demo-server/internal/repository"
)

type AuthService struct {
	users     *repository.UserRepository
	jwtSecret string
}

func NewAuthService(users *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{users: users, jwtSecret: jwtSecret}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginReq) (*dto.LoginResp, error) {
	user, err := s.users.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(40100, 401, "用户名或密码错误", "")
		}
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, apperror.New(40100, 401, "用户名或密码错误", "")
	}
	token, err := auth.IssueAccessToken(s.jwtSecret, user.ID, user.Username, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	name := user.DisplayName
	if name == "" {
		name = user.Username
	}
	return &dto.LoginResp{
		AccessToken: token,
		UserID:      user.ID,
		UserName:    name,
	}, nil
}
