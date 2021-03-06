// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/iost-official/Go-IOS-Protocol/core/tx"
	"github.com/iost-official/Go-IOS-Protocol/vm"
	"github.com/iost-official/Go-IOS-Protocol/vm/lua"
	"github.com/spf13/cobra"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile contract files to smart contract",
	Long:  `Compile contract files to smart contract. `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(`Error: source file not given`)
			return
		}
		path := args[0]
		fd, err := ReadFile(path)
		if err != nil {
			fmt.Println("Read file failed: ", err.Error())
			return
		}
		rawCode := string(fd)

		var contract vm.Contract
		switch Language {
		case "lua":
			parser, _ := lua.NewDocCommentParser(rawCode)
			contract, err = parser.Parse()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		mTx := tx.NewTx(int64(Nonce), contract)

		bytes := mTx.Encode()

		if dest == "default" {
			dest = ChangeSuffix(args[0], ".sc")
		}

		err = SaveTo(dest, bytes)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var Language string
var dest string
var Nonce int

func init() {
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flags().StringVarP(&Language, "language", "l", "lua", "Set language of contract, Support lua")
	compileCmd.Flags().IntVarP(&Nonce, "nonce", "n", 1, "Set Nonce of this Transaction")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
