// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vorbis_test

import (
	"os"
	"testing"

	"github.com/hajimehoshi/go-vorbis"
)

func TestDecode(t *testing.T) {
	file, err := os.Open("testdata/punch.ogg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out, err := vorbis.Decode(file)
	if err != nil {
		panic(err)
	}
	_ = out
}
