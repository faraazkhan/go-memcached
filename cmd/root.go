// Copyright © 2016 Faraaz Khan <faraaz@rationalizeit.us>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var connectionString, host string
var port int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-memcached",
	Short: "Simple CLI to interact with memcached",
	Long: `
:'######::::'#######:::::::::::'##::::'##:'########:'##::::'##::'######:::::'###:::::'######::'##::::'##:'########:'########::
'##... ##::'##.... ##:::::::::: ###::'###: ##.....:: ###::'###:'##... ##:::'## ##:::'##... ##: ##:::: ##: ##.....:: ##.... ##:
 ##:::..::: ##:::: ##:::::::::: ####'####: ##::::::: ####'####: ##:::..:::'##:. ##:: ##:::..:: ##:::: ##: ##::::::: ##:::: ##:
 ##::'####: ##:::: ##:'#######: ## ### ##: ######::: ## ### ##: ##:::::::'##:::. ##: ##::::::: #########: ######::: ##:::: ##:
 ##::: ##:: ##:::: ##:........: ##. #: ##: ##...:::: ##. #: ##: ##::::::: #########: ##::::::: ##.... ##: ##...:::: ##:::: ##:
 ##::: ##:: ##:::: ##:::::::::: ##:.:: ##: ##::::::: ##:.:: ##: ##::: ##: ##.... ##: ##::: ##: ##:::: ##: ##::::::: ##:::: ##:
. ######:::. #######::::::::::: ##:::: ##: ########: ##:::: ##:. ######:: ##:::: ##:. ######:: ##:::: ##: ########: ########::
:......:::::.......::::::::::::..:::::..::........::..:::::..:::......:::..:::::..:::......:::..:::::..::........::........:::
Go-Memcached is a CLI library for memcached written in GoLang.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().IntVar(&port, "port", 11211, "Port to run Application server on")
	viper.BindPFlag("port", RootCmd.Flags().Lookup("port"))
	RootCmd.PersistentFlags().StringVar(&host, "host", "127.0.0.1", "host memcached is running on.")
	viper.BindPFlag("host", RootCmd.Flags().Lookup("host"))
	connectionString = fmt.Sprintf("%s:%v", host, port)
}
