// Copyright © 2019 Onyedikachi Solomon Okwa <solozyokwa@gmail.com>
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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(version)
}

var (
	godlVersion = "unknown version"
	gitHash     = "unknown commit"
	goVersion   = "unknow go version"
	buildDate   = "unknown build date"
)

// versionCmd represents the version command
var version = &cobra.Command{
	Use:   "version",
	Short: "Show the godl version information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nGo version: %s\nGit hash: %s\nBuilt: %s\n",
			godlVersion, goVersion, gitHash, buildDate)
	},
}
