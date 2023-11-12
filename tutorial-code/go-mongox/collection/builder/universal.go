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

package main

import (
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/types"
)

func main() {
	bsonx.M("姓名", "陈明勇") // bson.M{"姓名": "陈明勇"}

	bsonx.Id("陈明勇") // bson.M{"_id": "陈明勇"}

	bsonx.D(types.KV("姓名", "陈明勇"), types.KV("手机号", "1888***1234")) // bson.D{{Key:"姓名", Value:"陈明勇"}, {Key:"手机号", Value:"1888***1234"}}

	bsonx.E("姓名", "陈明勇") // bson.E{Key:"姓名", Value:"陈明勇"}

	bsonx.A("陈明勇", "1888***1234") // bson.A{"陈明勇", "1888***1234"}
}
