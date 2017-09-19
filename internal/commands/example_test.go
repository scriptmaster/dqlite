// Copyright 2017 Canonical Ltd.
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
package commands_test

import (
	"fmt"
	"log"

	"github.com/CanonicalLtd/dqlite/internal/commands"
)

// Create a new dqlite Command with a few parameters, serialize it, and the
// finally deserialize it and read back its parameters.
func Example() {
	data, err := commands.Marshal(commands.NewOpen("test.db"))
	if err != nil {
		log.Fatalf("failed to marshal open command: %v", err)
	}
	cmd, err := commands.Unmarshal(data)
	if err != nil {
		log.Fatalf("failed to unmarshal open command: %v", err)
	}
	params := cmd.Params.(*commands.Command_Open)

	// Output:
	// 11
	// test.db
	fmt.Println(len(data))
	fmt.Println(params.Open.Name)
}