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
	"fmt"

	"github.com/chenmingyong0423/go-mongox/bsonx"
)

func main() {
	// bson.M{"姓名": "陈明勇"}
	m := bsonx.M("姓名", "陈明勇")
	fmt.Printf("%#v\n\n", m)

	// bson.M{"_id": "陈明勇"}
	id := bsonx.Id("陈明勇")
	fmt.Printf("%#v\n\n", id)

	// bson.D{bson.E{Key:"姓名", Value:"陈明勇"}, bson.E{Key:"手机号", Value:"1888***1234"}}
	d := bsonx.D(bsonx.KV("姓名", "陈明勇"), bsonx.KV("手机号", "1888***1234"))
	fmt.Printf("%#v\n\n", d)

	// bson.E{Key:"姓名", Value:"陈明勇"}
	e := bsonx.E("姓名", "陈明勇")
	fmt.Printf("%#v\n\n", e)

	// bson.A{"陈明勇", "1888***1234"}
	a := bsonx.A("陈明勇", "1888***1234")
	fmt.Printf("%#v", a)
}
