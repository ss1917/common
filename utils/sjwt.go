package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	//"io"
	"time"
)

var (
	key []byte = []byte("mQKmWgojkuiUv73ZmazmWK60562AF4")
)

type MyCustomClaims struct {
	Username string `json:"username"`
	UserID   int    `json:"userid"`
	RegionId int    `json:"regionid"`
	jwt.StandardClaims
}

// 产生json web token
func GenToken(name string, id, regionid int) string {
	claims := MyCustomClaims{
		name,
		id,
		regionid,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 100),
			ExpiresAt: int64(time.Now().Unix() + 1000),
			Id:        "15618718060",
			Issuer:    "shenshuo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return ""
	}
	return ss
}

// 校验token是否有效,并返回信息
func CheckToken(ss string) (name string, id, regionid int) {
	token, _ := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		id = claims.UserID
		name = claims.Username
		regionid = claims.RegionId
	}
	return
}
