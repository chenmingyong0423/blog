// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject

package wire

import (
	"chenmingyong0423/blog/tutorial-code/wire/internal/post/handler"
	"chenmingyong0423/blog/tutorial-code/wire/internal/post/service"
	"chenmingyong0423/blog/tutorial-code/wire/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(
		handler.NewPostHandler,
		service.NewPostService,
		ioc.NewGinEngineAndRegisterRoute,
	)
	return &gin.Engine{}
}

func InitializeAppV2() *gin.Engine {
	wire.Build(
		handler.PostSet,
		ioc.NewGinEngineAndRegisterRoute,
	)
	return &gin.Engine{}
}

func InitializeAppV3() *gin.Engine {
	wire.Build(
		handler.NewPostHandler,
		service.NewPostServiceV2,
		ioc.NewGinEngineAndRegisterRoute,
		wire.Bind(new(service.IPostService), new(*service.PostService)),
	)
	return &gin.Engine{}
}

type Name string

func NewName() Name {
	return "陈明勇"
}

type PublicAccount string

func NewPublicAccount() PublicAccount {
	return "公众号：Go技术干货"
}

type User struct {
	MyName          Name
	MyPublicAccount PublicAccount
}

func InitializeUser() *User {
	wire.Build(
		NewName,
		NewPublicAccount,
		wire.Struct(new(User), "MyName", "MyPublicAccount"),
	)
	return &User{}
}

func InjectUser() User {
	wire.Build(wire.Value(User{MyName: "陈明勇"}))
	return User{}
}

func InjectPostService() service.IPostService {
	wire.Build(wire.InterfaceValue(new(service.IPostService), &service.PostService{}))
	return nil
}

func NewUser() User {
	return User{MyName: Name("陈明勇"), MyPublicAccount: PublicAccount("公众号：Go技术干货")}
}

func GetUserName() Name {
	wire.Build(
		NewUser,
		wire.FieldsOf(new(User), "MyName"),
	)
	return ""
}
