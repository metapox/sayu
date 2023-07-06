/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/metapox/sayu/config"
	"github.com/metapox/sayu/converter"
	"github.com/metapox/sayu/input"
	"github.com/metapox/sayu/output"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config{}
		data, err := os.ReadFile("config.yaml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// check execute from input
		fmt.Println("以下の設定で実行します")
		fmt.Println(string(data))
		fmt.Println("実行しますか？[y/N]: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text != "y" {
			fmt.Println("終了します")
			os.Exit(0)
		}
		yaml.Unmarshal(data, &c)

		inp, err := input.CreateInput(c.Input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		outp, err := output.CreateOutput(c.Output)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pipeline := converter.NewPipeline(inp, outp)
		for _, cc := range c.Converters {
			conv, err := converter.CreateConverter(cc)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pipeline.RegistConverter(conv)
		}
		pipeline.Start()
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
