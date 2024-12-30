package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	width, height := 800, 600
	words, err := getWordsFromDB()
	if err != nil {
		panic(fmt.Sprintf("Erro ao consultar palavras no banco de dados: %v", err))
	}

	file, err := os.Create("wordcloud.svg")
	if err != nil {
		fmt.Println("Erro ao criar arquivo SVG:", err)
		return
	}
	defer file.Close()
	generateCloudWords(file, words, width, height)
	fmt.Println("Arquivo 'wordcloud.svg' gerado com sucesso!")
}
