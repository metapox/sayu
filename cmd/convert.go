/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/metapox/sayu/converter"
	"github.com/metapox/sayu/converter/converters"
	"github.com/metapox/sayu/input"
	"github.com/metapox/sayu/output"

	"github.com/spf13/cobra"
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
		inp := input.NewSolrLog("test/test.log")
		outp := output.NewLocalFile("test/new_test.log")
		pipeline := converter.NewPipeline(inp, outp)
		converter := converters.NewLogConverter()
		pipeline.RegistConverter(converter)
		fmt.Println(pipeline.ShowConvertersInfo())
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
