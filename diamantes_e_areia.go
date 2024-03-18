package main // Contém a função main para executar o programa

import ( // Pacotes necessários para o programa
	"fmt"
)

// Função para extrair diamantes de uma string
func extrairDiamantes(entrada string) int { // Entrada: string contendo possíveis diamantes
	var pilha []rune // Pilha para ajudar na identificação de diamantes
	diamantes := 0   // Contador de diamantes encontrados

	// Loop através de cada caractere da string
	for _, caractere := range entrada {
		if caractere == '<' { // Se encontrar '<', empilha-o
			pilha = append(pilha, caractere)
		} else if caractere == '>' && len(pilha) > 0 { // Se encontrar '>' e a pilha não estiver vazia
			pilha = pilha[:len(pilha)-1] // Desempilha o '<' correspondente
			diamantes++                  // Incrementa o contador de diamantes
		}
	}

	return diamantes // Retorna o número de diamantes encontrados
}

func main() {
	var numCasos int // Variável para armazenar o número de casos de teste

	fmt.Scanln(&numCasos) // Lê o número de casos de teste

	for i := 0; i < numCasos; i++ { // Loop através de cada caso de teste
		var entrada string // Variável para armazenar a entrada do usuário

		fmt.Scanln(&entrada)                   // Lê a entrada do usuário
		fmt.Println(extrairDiamantes(entrada)) // Imprime o número de diamantes encontrados
	}
}
