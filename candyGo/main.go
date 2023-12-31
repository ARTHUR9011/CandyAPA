package main

import (
	"log"
	"runtime"
	"time"
)



func main() {
	// Classificações de exemplo
	ratings := []int{1, 2, 3,2, 1, 4}

	// Inicia a contagem do tempo
	startTime := time.Now()

	// Chama a função candy com as classificações de exemplo
	result := candy(ratings)

	// Finaliza a contagem do tempo
	endTime := time.Now()

	// Imprime o resultado
	log.Printf("Resultado: %d", result)

	// Calcula e imprime o tempo de execução
	elapsedTime := endTime.Sub(startTime)
	log.Printf("Tempo de execução: %s", elapsedTime)

	// Obtém e imprime o uso de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Convertendo para megabytes (MB)
	memoryUsageMB := float64(m.Alloc) / 1024 / 1024
	log.Print("Uso de memória: ", m.Alloc, " bytes")
	log.Printf("Uso de memória: ~%.2f MB", memoryUsageMB)
}

func candy(ratings []int) int {
	// Define n como o tamanho do vetor
	n := len(ratings)

	// Log da entrada
	log.Printf("Classificações recebidas: %v", ratings)

	// Se o tamanho do vetor for menor ou igual a 1, retorna o próprio tamanho
	if n <= 1 {
		log.Printf("Apenas uma criança ou nenhum. Retornando %d como a quantidade total de doces.", n)
		return n
	}

	// Cria um vetor para armazenar a quantidade de doces atribuídos a cada criança
	candies := make([]int, n)

	// Atribui 1 doce à primeira criança
	candies[0] = 1

	// Loop para percorrer as classificações da esquerda para a direita
	for i := 1; i < n; i++ {
		// Se a classificação atual for maior que a anterior, atribui um doce a mais
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			// Se não, atribui apenas 1 doce
			candies[i] = 1
		}
		// Log do número de doces atribuídos a cada criança durante o loop
		log.Printf("Criança %d recebeu %d doces", i, candies[i])
	}

	// Inicializa a variável total com a quantidade de doces da última criança
	total := candies[n-1]

	// Loop para percorrer as classificações da direita para a esquerda
	for i := n - 2; i >= 0; i-- {
		// Se o rating atual é maior que o próximo, ajusta para que tenha 1 a mais que o próximo
		if ratings[i] > ratings[i+1] {
			// Se a próxima criança tem mais doces, nada muda
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		// Adiciona a quantidade de doces da criança atual ao total
		total += candies[i]
		// Log do número de doces atribuídos a cada criança durante o loop reverso
		log.Printf("Criança %d recebeu %d doces (total acumulado: %d)", i, candies[i], total)
	}

	// Log do resultado
	log.Printf("Quantidade total de doces atribuídos: %d", total)

	// Retorna a quantidade total de doces atribuídos
	return total
}


// Calcula o máximo entre dois inteiros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}