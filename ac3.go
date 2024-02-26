package main

import (
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(contato Contato, listaContatos []Contato) ([]Contato, error) {
	for i := range listaContatos {
		if listaContatos[i].Nome == "" || listaContatos[i].Email == "" {
			listaContatos[i] = contato
			return listaContatos, nil
		}
	}
	return listaContatos, fmt.Errorf("A lista de contatos está cheia. Não foi possível adicionar o contato.")
}

func excluirContato(listaContatos []Contato) ([]Contato, error) {
	contatosVazios := 0
	for i := len(listaContatos) - 1; i >= 0; i-- {
		if listaContatos[i].Nome == "" && listaContatos[i].Email == "" {
			contatosVazios++
		}
		if listaContatos[i].Nome != "" || listaContatos[i].Email != "" {
			listaContatos[i] = Contato{}
			return listaContatos, nil
		}
	}
	if contatosVazios == len(listaContatos) {
		return listaContatos, fmt.Errorf("A lista de contatos está vazia. Não há contatos para excluir.")
	}
	return listaContatos, nil
}

func main() {
	listaContatos := make([]Contato, 5)
	var err error

	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Adicionar Contato")
		fmt.Println("2. Excluir Contato")
		fmt.Println("3. Sair")

		var opcao int
		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Nome do Contato: ")
			var nome string
			fmt.Scanln(&nome)

			fmt.Print("E-mail do Contato: ")
			var email string
			fmt.Scanln(&email)

			novoContato := Contato{
				Nome:  nome,
				Email: email,
			}

			listaContatos, err = adicionarContato(novoContato, listaContatos)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Contato adicionado com sucesso.")
			}

		case 2:
			listaContatos, err = excluirContato(listaContatos)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Último contato não vazio excluído com sucesso.")
			}

		case 3:
			fmt.Println("Saindo...")

			// Exibindo toda a lista de contatos ao sair
			fmt.Println("\nLista de Contatos:")
			for _, contato := range listaContatos {
				if contato.Nome != "" || contato.Email != "" {
					fmt.Printf("Nome: %s, E-mail: %s\n", contato.Nome, contato.Email)
				}
			}

			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
