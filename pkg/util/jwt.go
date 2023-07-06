package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-web-template/global"
	"go-web-template/internal/model"
	"go-web-template/pkg/logger"

	"time"
)

func CreateToken(u model.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = u.ID
	claims["userName"] = u.Username
	claims["exp"] = time.Now().Add(time.Second * time.Duration(global.Cfg.JWT.Expire)).Unix()
	tokenString, err := token.SignedString([]byte(global.Cfg.JWT.Secret))
	if err != nil {
		logger.Error(err)
	}
	return tokenString
}

func ParseToken(token string) (*model.User, error) {
	u := model.User{}
	if token == "" {
		return nil, errors.New("token为空")
	}
	// 解析JWT令牌
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.Cfg.JWT.Secret), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userId := uint(claims["userId"].(float64))
		userName := claims["userName"].(string)
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		logger.Debug(fmt.Sprintf("UserID: %d UserName: %s Expiration Time: %s", userId, userName, expirationTime.String()))
		if expirationTime.Before(time.Now()) {
			logger.Error(fmt.Sprintf("token expirationTime:%s", expirationTime.String()))
		} else {
			u.ID = userId
			u.Username = userName
		}
	} else {
		logger.Error(err)
		return nil, err
	}
	return &u, nil
}
