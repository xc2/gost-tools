package main

import (
	"log"
	"os"

	gost_cmd "github.com/go-gost/x/config/cmd"

	"github.com/spf13/cobra"
)

var (
	VERSION = "0.0.0-PLACEHOLDER"
)

func main() {
	config := &cobra.Command{
		Use:   "config",
		Short: "generate config file content",
		Run: func(cmd *cobra.Command, args []string) {
			services, _ := cmd.Flags().GetStringSlice("services")
			nodes, _ := cmd.Flags().GetStringSlice("nodes")
			json, _ := cmd.Flags().GetBool("json")
			c, err := gost_cmd.BuildConfigFromCmd(services, nodes)
			if err != nil {
				log.Fatalln(err)
				return
			}
			if json {
				c.Write(os.Stdout, "json")
			} else {
				c.Write(os.Stdout, "yaml")
			}
		},
	}

	config.Flags().StringSliceP("services", "L", []string{}, "service list")
	config.Flags().StringSliceP("nodes", "F", []string{}, "chain node list")
	config.Flags().BoolP("json", "j", false, "use json instead of yaml")
	config.MarkFlagRequired("services")

	root := &cobra.Command{Use: os.Args[0], Version: VERSION}
	root.AddCommand(config)

	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
