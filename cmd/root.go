package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var listFlag bool
var createFlag string
var deleteFlag string


// Comando raiz do CLI
var rootCmd = &cobra.Command{
	Use:   "file-manager-cli",
	Short: "File Manager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		
		switch {
			case listFlag:
				listCmd.Run(cmd, args)
				return
			case createFlag != "":
				createCmd.Run(cmd, []string{createFlag})
				return
			case deleteFlag != "":
				removeCmd.Run(cmd,  []string{deleteFlag})
				return
			default:
				fmt.Println(color.YellowString("────────────────────────────────────────────────────────────────────"))
				fmt.Println("❗ Use 'file-manager-cli --help' to see the available commands❗")
				fmt.Println(color.YellowString("────────────────────────────────────────────────────────────────────"))
				return
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "List the files in the specified directory")

	rootCmd.AddCommand(createCmd)

	rootCmd.AddCommand(removeCmd)
	rootCmd.Flags().StringVarP(&deleteFlag, "remove", "r", "", "Delete a file from the system")

	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(renameCmd)

}

// Função para executar os comandos
func Execute() error {
	return rootCmd.Execute()
}

