package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "inventorybook-svc",
	Short: "Inventory Book Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use serve to start a server")
		fmt.Println("Use -h to see the list of command")
	},
}

func Run() {
	rootCommand.AddCommand(restCommand)

	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
