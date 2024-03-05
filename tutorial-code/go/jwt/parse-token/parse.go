// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func ParseJwt(key any, jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}

func main() {
	jwtKey := make([]byte, 32) // 生成32字节（256位）的密钥
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "程序员陈明勇",
		"sub": "chenmingyong.cn",
		"aud": "Programmer",
		"exp": time.Now().Add(time.Second * 10).UnixMilli(),
	})
	jwtStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	// 解析 jwt
	claims, err := ParseJwt(jwtKey, jwtStr, jwt.WithExpirationRequired())
	if err != nil {
		panic(err)
	}
	fmt.Println(claims)
}
