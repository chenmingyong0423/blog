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
	"context"
	"log/slog"
	"os"
)

func main() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).WithGroup("information")
	jsonLogger.InfoContext(context.Background(), "json-log", slog.String("name", "chenmingyong"), slog.Int("phone", 1234567890))

	textLogger := slog.New(slog.NewTextHandler(os.Stdout, nil)).WithGroup("information")
	textLogger.InfoContext(context.Background(), "json-log", slog.String("name", "chenmingyong"), slog.Int("phone", 1234567890))
}
