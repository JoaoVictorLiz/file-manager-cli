package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use: "rename [old-name] [new-name]",
	Short: "Rename a file or directory",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileOldName := args[0]
		fileNewName := args[1]

		if _, err := os.Stat(fileOldName); os.IsNotExist(err) {
			fmt.Println("❌ Error: The file does not exist:", fileOldName)
			return
		}


		err := os.Rename(fileOldName, fileNewName)
		if err != nil {
			fmt.Println("❌ Failed to rename the file:", err)
			return
		}

		fmt.Printf("✅ File '%s' renamed to '%s' successfully!\n", fileOldName, fileNewName)
	},
}

