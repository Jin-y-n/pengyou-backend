package common

import (
	"context"
	"go.uber.org/zap"
	string2 "pengyou/utils/check/string"
	"pengyou/utils/log"
	"strconv"
	"sync"
)

type tokenKey struct{}

var tokenCtxKey = tokenKey{}

func SetTokenInContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtxKey, token)
}

func GetTokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenCtxKey).(string)
	return token, ok
}

func SetTokenInContextDefault(token string) context.Context {
	return SetTokenInContext(context.Background(), token)
}

func GetTokenFromContextDefault() (string, bool) {
	return GetTokenFromContext(context.Background())
}

func processToken(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	token, ok := GetTokenFromContext(ctx)
	if !ok {
		log.Logger.Fatal("Token not found in context")
	}
	log.Logger.Info("Processing ", zap.String("token:", token))
}

func CheckUserIdDefault(userId uint) bool {
	tokenFromContext, suc := GetTokenFromContext(context.Background())
	if !suc {
		log.Logger.Fatal("Token not found in context")
		return false
	}

	if !string2.IsNumberString(&tokenFromContext) {
		return false
	}
	id, _ := strconv.Atoi(tokenFromContext)
	return userId == uint(id)
}
