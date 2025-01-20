package utils

import (
	"context"
	"errors"

	"github.com/twinbeard/goLearning/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	SUUID, ok := ctx.Value("subjectUUID").(string)

	if !ok {
		return "", errors.New("failed to get subject UUID")
	}

	return SUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	SUUID, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}

	// get infoUser Redis from uuid
	var inforUser InfoUserUUID
	if err := cache.GetCache(ctx, SUUID, &inforUser); err != nil {
		return 0, err
	}

	return inforUser.UserId, nil
}
