/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pingTest/pkg"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteServer,err:=cmd.Flags().GetString("server")
		if err!=nil{
			fmt.Printf("flags get server  address error:%v \n",err)
			return
		}
		c,err:=cmd.Flags().GetInt("c")
		if err!=nil{
			fmt.Printf("get packet number  error:%v \n",err)
			return
		}
		w,err:=cmd.Flags().GetInt("w")
		if err!=nil{
			fmt.Printf("get timeout error:%v \n",err)
			return
		}
		pkg.GoPing(remoteServer,c,w)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().String("server","127.0.0.1:8090","remote server ip and port")
	clientCmd.Flags().Int("c",4,"send packet number")
	clientCmd.Flags().Int("w",1,"timeout (s)")
}
