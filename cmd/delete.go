package cmd


import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use: "remove [file-name]",
	Short: "Delete a file from the system",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		if _, err := os.Stat(fileName); 	
		os.IsNotExist(err) {
			fmt.Println("❌ Error: The file does not exist:", fileName)
			return
		}

		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("❌ Failed to delete the file:", err)
			return
		}
		
		fmt.Printf("✅ File '%s' successfuly deleted", fileName)
	},
}

