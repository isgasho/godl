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
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)
}

// listCmd represents the list command
var list = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the downloaded versions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := getDownloadDir()
		if err != nil {
			return err
		}

		return listDownloadedBinaryArchives(d)
	},
}

func listDownloadedBinaryArchives(downloadDir string) error {
	const (
		archiveSuffix = ".darwin-amd64.tar.gz"
		archivePrefix = "go"
	)

	// Create download directory and its parent
	must(os.MkdirAll(downloadDir, os.ModePerm))

	files, err := ioutil.ReadDir(downloadDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, archiveSuffix) {
			archiveVersion := strings.TrimSuffix(
				strings.TrimPrefix(name, archivePrefix),
				archiveSuffix,
			)
			fmt.Println(archiveVersion)
		}
	}

	return nil
}
