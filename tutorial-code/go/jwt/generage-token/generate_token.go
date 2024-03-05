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
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(key any, method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(key)
}

func main() {
	jwtKey := make([]byte, 32) // 生成32字节（256位）的密钥
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}
	jwtStr, err := GenerateJwt(jwtKey, jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "程序员陈明勇",
		"sub": "chenmingyong.cn",
		"aud": "Programmer",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(jwtStr)
}
