package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flags
var createFileFlag string
var createDirFlag string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if createFileFlag != "" {
			createFile(createFileFlag)
		} else if createDirFlag != "" {
			createDir(createDirFlag)
		} else {
			fmt.Println("❌ Use -f to create a file or -d to create a directory.")
		}
	},
}

// Criar Arquivo
func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("❌ Failed to create a file:", err)
		return
	}
	file.Close()
	fmt.Println("✅ File create successfully:", name)
}

// Criar Diretório
func createDir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Println("❌ Failed to create a directory:", err)
		return
	}
	fmt.Println("✅ Directory created successfully:", name)
}

// Inicializar Flags
func init() {
	createCmd.Flags().StringVarP(&createFileFlag, "file", "f", "", "Create a new file")
	createCmd.Flags().StringVarP(&createDirFlag, "directory", "d", "", "Create a new directory")
}
