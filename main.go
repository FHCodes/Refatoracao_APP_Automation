package main

import (
	"Refatoracao_App_Automation/utils"
	"fmt"
)

func main() {
	// Teste criação de pasta
	pathName, err := utils.CreateFolder("C:/Codigo_pessoal/Refatoracao_APP_Automation/cara_de_toba.xlsx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pasta criada em:", pathName)
	}

	// Teste deleção de pasta
	erro := utils.DeleteFolder(pathName)
	fmt.Print(erro)
}
