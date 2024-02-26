package main

import (
  "fmt"
  "math"
  "geometria"
)

type Ponto struct {
  X float64
  Y float64
}

func main() {
  vetor := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  fmt.Println("Elementos do vetor:")
  for i := 0; i < len(vetor); i++ {
    fmt.Println(vetor[i])
  }

  fmt.Println("Digite uma string:")
  var input string
  fmt.Scanln(&input)
  resultadoInversao := inverteString(input)
  fmt.Println("String invertida:", resultadoInversao)

  ponto := Ponto{X: 3, Y: 4}
  distancia := ponto.DistanciaOrigem()
  fmt.Println("Distância até a origem:", distancia)

  fmt.Println("Digite o comprimento do retângulo:")
  var comprimento float64
  fmt.Scanln(&comprimento)

  fmt.Println("Digite a largura do retângulo:")
  var largura float64
  fmt.Scanln(&largura)

  area := geometria.CalculaAreaRetangulo(comprimento, largura)
  fmt.Println("Área do retângulo:", area)

  perimetro := geometria.CalculaPerimetroRetangulo(comprimento, largura)
  fmt.Println("Perímetro do retângulo:", perimetro)
}

func inverteString(str string) string {
  runes := []rune(str)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func (p Ponto) DistanciaOrigem() float64 {
  return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
