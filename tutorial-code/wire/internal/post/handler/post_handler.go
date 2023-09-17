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

package handler

import (
	"chenmingyong0423/blog/tutorial-code/wire/internal/post/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var PostSet = wire.NewSet(NewPostHandler, service.NewPostService)

func NewPostHandler(serv service.IPostService) *PostHandler {
	return &PostHandler{serv: serv}
}

type PostHandler struct {
	serv service.IPostService
}

func (h *PostHandler) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/post/:id", h.GetPostById)
}

func (h *PostHandler) GetPostById(ctx *gin.Context) {
	content := h.serv.GetPostById(ctx, ctx.Param("id"))
	ctx.String(http.StatusOK, content)
}
