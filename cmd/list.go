package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List the files in the specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		// Verificar se o diretório foi passado como argumento
		
		var dir string
		if len(args) > 0 {
			dir = args[0]
		} else {
			dir, _ = os.Getwd() // Pega o diretório atual se nenhum for especificado
		}

	
		// Listar os arquivos do diretório
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error to list the files:", err)
			return
		}

		fmt.Println(color.CyanString("\n📂 Directory: %s", dir))
		fmt.Println(color.YellowString("──────────────────────────────────"))

		// Imprimir os arquivos encontrados
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(color.BlueString("📁 %s", file.Name()))
			} else {
				// Ira me retornar oa extensao do arquivo(e.g .go,.exe)
				ext := filepath.Ext(file.Name())
				icon := getFileIcon(ext)
				fmt.Printf("%s %s\n", icon, file.Name())
			}
		}

		fmt.Println(color.YellowString("──────────────────────────────────\n"))
	},
}

func getFileIcon(ext string) string {
	// Ícones para diferentes tipos de arquivo
	switch ext {
	case ".txt":
		return "📄"
	case ".go":
		return "🐹"
	case ".jpg", ".png", ".gif":
		return "🖼️"
	case ".zip", ".tar", ".gz":
		return "📦"
	case ".exe":
		return "⚙️"
	default:
		return "📜"
	}
}

