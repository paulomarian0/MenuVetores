package main

import (
	"fmt"
	"sort"
)

/*
Fac¸a um programa que leia um vetor de 5 posic¸oes para n ˜ umeros reais e, depois, um ´
codigo inteiro. Se o c ´ odigo for zero, finalize o programa; se for 1, mostre o vetor na ordem ´
direta; se for 2, mostre o vetor na ordem inversa. Caso, o codigo for diferente de 1 e 2 ´
escreva uma mensagem informando que o codigo ´ e inv ´ alido.
*/

func main() {
	const tamanho = 5

	var vetor [tamanho]int

	slice := []int{} //slice pra poder armazenar todos valores do vetor e poder ordenar depois

	// for pra entrar todos os dados
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite o valor [%d]:", i)
		fmt.Scan(&vetor[i])
		slice = append(slice, vetor[i])

	}
	for {

		fmt.Println("Digite '0' para finalizar o programa")
		fmt.Println("Digite '1' para exibir o vetor na ordem crescente")
		fmt.Println("Digite '2' para exibir o vetor na ordem decrescente")
		fmt.Println("Digite '3' para adicionar um valor")

		var entrada int // entrada pra switch case
		fmt.Scan(&entrada)

		if entrada == 0 {
			fmt.Println("Finalizando o programa")
			break
		} else {

			//SWITCH CASE
			switch entrada {

			case 1:
				sort.Ints(slice) // Ordenar Crescente

				fmt.Println("Vetor exibido em ordem crescente")
				fmt.Println(slice)
				break

			case 2:

				sort.Sort(sort.Reverse(sort.IntSlice(slice)))
				fmt.Println("Vetor exibido na ordem decrescente")
				fmt.Println(slice)
				break

			case 3:
				fmt.Println("Digite o valor a ser adicionado")
				var valorAdicionado int
				fmt.Scan(&valorAdicionado)

				slice = append(slice, valorAdicionado)

				// Sem break para que os outros casos leiam o valor adicionado

			default:
				fmt.Println("Entrada inválida")
				break

			}

		}

	}
}
