package redis

import (
	"go.uber.org/zap"
	"time"
)

// 存放user_id:token
func UserIdToken(userID string, token string, exp time.Duration) error {
	if err := rdb.Set(userID, token, exp).Err(); err != nil {
		zap.L().Error("设置user_id和token进redis失败")
		return err
	}
	return nil
}

func GetTokenKey(userID string) (string, error) {
	value, err := rdb.Get(userID).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
