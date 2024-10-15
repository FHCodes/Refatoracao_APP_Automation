package utils

import (
	"fmt"
	"os"
)

func DeleteFolder(pathName string) error {
	// Remove o diretório e todo o seu conteúdo
	err := os.RemoveAll(pathName)
	if err != nil {
		return fmt.Errorf("erro ao deletar diretório: %v", err)
	}
	return nil
}
