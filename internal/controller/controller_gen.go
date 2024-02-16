package controller

import (
	"log"
	"os"
)

func ControllerGenerator(nomeArquivo, genfile string) {
	caminhoArquivo := nomeArquivo

	err := os.WriteFile(caminhoArquivo, []byte(genfile), 0666)

	if err != nil {
		panic(err)
	}

	log.Printf("Docker-compose gerado com sucesso!")
}
