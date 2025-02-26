package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move [file-origin] [directory-destiny]",
	Short: "Move a file to a specific directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[0]
		destDir := args[1]

		// Verifica se o arquivo existe
		if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
			fmt.Println("❌ Error: The file does not exist:", sourceFile)
			return
		}

		// Verifica se o diretório de destino existe, senão cria
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			fmt.Println("⚠️ Warning: Destination directory does not exist. Creating...")
			err = os.MkdirAll(destDir, os.ModePerm)
			if err != nil {
				fmt.Println("❌ Error: Failed to create directory:", err)
				return
			}
		}

		// Constrói corretamente o caminho do novo arquivo
		newPath := filepath.Join(destDir, filepath.Base(sourceFile))

		// Move o arquivo
		err := os.Rename(sourceFile, newPath)
		if err != nil {
			fmt.Println("❌ Failed to move the file:", err)
			return
		}

		fmt.Printf("✅ File '%s' moved to '%s' successfully!\n", sourceFile, newPath)
	},
}
