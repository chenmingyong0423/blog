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
	"github.com/chenmingyong0423/go-mongox/builder/update"
)

func main() {
	// bson.D{bson.E{Key:"$set", Value:bson.M{"name":"陈明勇"}}}
	u := update.BsonBuilder().Set(bsonx.M("name", "陈明勇")).Build()
	fmt.Printf("%#v\n\n", u)

	// bson.D{bson.E{Key:"$inc", Value:bson.D{bson.E{Key:"orders", Value:1}, bson.E{Key:"ratings", Value:-1}}}}
	u = update.BsonBuilder().Inc(bsonx.D(bsonx.KV("orders", 1), bsonx.KV("ratings", -1))).Build()
	fmt.Printf("%#v\n\n", u)

	// bson.D{bson.E{Key:"$push", Value:bson.M{"scores":95}}}
	u = update.BsonBuilder().Push(bsonx.M("scores", 95)).Build()
	fmt.Printf("%#v\n\n", u)

	// bson.D{bson.E{Key:"$unset", Value:bson.D{primitive.E{Key:"quantity", Value:""}, bson.E{Key:"instock", Value:""}}}}
	u = update.BsonBuilder().Unset("quantity", "instock").Build()
	fmt.Printf("%#v", u)
}
