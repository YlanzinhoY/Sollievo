package processfile

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateFiles(backEndPath, filename, path string) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return err
	}
	fmt.Println("Diretório de trabalho atual:", wd)

	// Atualize o caminho do arquivo de origem conforme necessário
	srcFile := filepath.Join(wd, path)
	dstFile := filepath.Join(backEndPath, filename)

	// Adicionando depuração
	absSrcFile, err := filepath.Abs(srcFile)
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto do arquivo de origem:", err)
		return err
	}
	fmt.Println("Caminho absoluto do arquivo de origem:", absSrcFile)

	absDstFile, err := filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto do arquivo de destino:", err)
		return err
	}
	fmt.Println("Caminho absoluto do arquivo de destino:", absDstFile)

	// Verifique se o arquivo de origem existe
	if _, err := os.Stat(absSrcFile); os.IsNotExist(err) {
		fmt.Println("Arquivo de origem não existe:", absSrcFile)
		return err
	}

	err = copyFile(absSrcFile, absDstFile)
	if err != nil {
		fmt.Println("Erro ao copiar o arquivo:", err)
		return err
	}

	fmt.Printf("%s copiado com sucesso\n", filename)

	return nil
}
