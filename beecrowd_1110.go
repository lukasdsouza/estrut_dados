package main

import (
	"fmt"
	"strings"
)

// Definição do tipo Discarded, que é um heap que armazena as cartas descartadas.
type Discarded []int

// Método Len retorna o tamanho do heap.
func (d Discarded) Len() int {
	return len(d)
}

// Método Less compara dois elementos do heap e retorna true se o primeiro for menor que o segundo.
func (d Discarded) Less(i, j int) bool {
	return d[i] < d[j]
}

// Método Swap troca os dois elementos do heap.
func (d Discarded) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Método Push adiciona um elemento ao heap.
func (d *Discarded) Push(x interface{}) {
	*d = append(*d, x.(int))
}

// Método Pop remove e retorna o elemento do topo do heap.
func (d *Discarded) Pop() interface{} {
	n := len(*d)
	card := (*d)[n-1]
	*d = (*d)[:n-1]
	return card
}

// Método String converte o heap em uma string formatada.
func (d Discarded) String() string {
	s := make([]string, d.Len())
	for i := 0; i < d.Len(); i++ {
		s[i] = fmt.Sprintf("%d", d[i])
	}
	return strings.Join(s, ", ")
}

// Definição do tipo Queue, que é uma fila que armazena as cartas.
type Queue struct {
	cards []int
}

// Função NewQueue inicializa uma nova fila.
func NewQueue() *Queue {
	return &Queue{cards: []int{}}
}

// Método Len retorna o tamanho da fila.
func (q *Queue) Len() int {
	return len(q.cards)
}

// Método Enqueue adiciona um elemento ao final da fila.
func (q *Queue) Enqueue(card int) {
	q.cards = append(q.cards, card)
}

// Método Dequeue remove e retorna o elemento do início da fila.
func (q *Queue) Dequeue() int {
	front := q.cards[0]
	q.cards = q.cards[1:]
	return front
}

// Método Peek retorna o elemento do início da fila sem removê-lo.
func (q *Queue) Peek() int {
	return q.cards[0]
}

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		queue := NewQueue()
		for i := 1; i <= n; i++ {
			queue.Enqueue(i)
		}

		discarded := &Discarded{}
		for queue.Len() > 1 {
			discarded.Push(queue.Dequeue())
			queue.Enqueue(queue.Dequeue())
		}

		fmt.Println("Cartas descartadas:", discarded)
		fmt.Println("Carta restante:", queue.Peek())
	}
}
