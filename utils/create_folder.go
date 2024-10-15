package utils

import (
	"fmt"
	"os"
	"strings"
)

func CreateFolder(pathNameXlsm string) (string, error) {
	// 1. Divide o caminho pelo separador "/"
	temp := strings.Split(pathNameXlsm, "/")

	// 2. Pega a última parte (nome do arquivo) e remove a extensão
	nomePack := strings.Split(temp[len(temp)-1], ".")
	// Remove a extensão do nome do arquivo
	nomePack = nomePack[:len(nomePack)-1]

	// 3. Remonta o caminho sem o nome do arquivo
	temp = temp[:len(temp)-1]

	path := strings.Join(temp, "/")

	// 4. Monta o nome do diretório com "MASTERFILES_" e o nome sem extensão
	nomePackStr := strings.Join(nomePack, ".")
	pathName := fmt.Sprintf("%s/MASTERFILES_%s", path, nomePackStr)

	// 5. Verifica se o diretório existe, se não, cria o diretório
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		err = os.MkdirAll(pathName, os.ModePerm) // Cria o diretório com todas as permissões
		if err != nil {
			return "", fmt.Errorf("erro ao criar diretório: %v", err)
		}
	}

	// 6. Retorna o nome do caminho criado
	return pathName, nil
}
