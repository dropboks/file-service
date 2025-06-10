package service

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/dropboks/file-service/internal/domain/repository"
	"github.com/dropboks/file-service/pkg/constant"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type (
	UserService interface {
		SaveProfileImage(context context.Context, imageByte []byte, imageExt string) (string, error)
		RemoveProfileImage(context context.Context, imageName string) error
	}
	userService struct {
		userRepository repository.UserRepository
		logger         zerolog.Logger
	}
)

func NewUserService(userRepository repository.UserRepository, logger zerolog.Logger) UserService {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (u *userService) SaveProfileImage(context context.Context, imageBytes []byte, imageExt string) (string, error) {
	imageName := fmt.Sprintf("%s.%s", uuid.New().String(), imageExt)
	imagePath := fmt.Sprintf("%s/%s", constant.PROFILE_IMAGE_FOLDER, imageName)
	start := time.Now()
	err := u.userRepository.SaveProfileImage(context, constant.APP_BUCKET, imagePath, bytes.NewReader(imageBytes), int64(len(imageBytes)))
	duration := time.Since(start)
	u.logger.Info().Dur("duration", duration).Msg("profile image saved")
	if err != nil {
		return "", err
	}
	return imagePath, nil
}

func (u *userService) RemoveProfileImage(context context.Context, imageName string) error {
	return nil
	// panic("unimplemented")
}
