# Relatorio APA   
   
   
# Candies   
   
## qual o problema? 
   
**Instruções**:   
Alice é uma professora de jardim de infância. Ela quer dar alguns doces para as crianças de sua classe. Todas as crianças se sentam em uma fila, e cada uma delas tem uma pontuação de classificação de acordo com seu desempenho na classe. Alice quer dar pelo menos 1 doce para cada criança. Se duas crianças se sentam uma ao lado da outra, então aquela com a pontuação mais alta deve receber mais doces. Alice quer minimizar o número total de doces que ela precisa comprar.   
   
### Exemplo

Ela dá doces aos alunos nas seguintes quantidades mínimas: . Ela precisa comprar no mínimo 10 doces.   
   
### Descrição da Função

Complete a função candies no editor abaixo.

candies tem os seguintes parâmetros:

int n: o número de crianças na classe
int arr[n]: as classificações de cada aluno
Retorna

int: o número mínimo de doces que Alice deve comprar
Formato de Entrada

A primeira linha contém um número inteiro, , o tamanho de .
Cada uma das próximas linhas contém um número inteiro  indicando a classificação do aluno na posição .

Restrições   
### Entrada de Amostra 0
   
```

3
1
2
2

```
**Saída de Amostra 0**:   
```

4

```
   
### Entrada de Amostra 1:

   
```
10
2
4
2
6
1
7
8
9
2
1
```
**Saída de Amostra 1**:   
```

19







```
### Entrada de Amostra 2:   
   
```
8
2
4
3
5
2
6
4
5
```
**Saída de Amostra 2**:   
```

12

```
   
   
## Linguagem para resolver o problemas:

   
Para resolver o problema utilzamos 3 linguagens diferentes:
   
1. Java   
2. Golang   
3. Rust   
   
   
Java e golang foram utilizadas pela familiaridade dos integrantes do grupo. A linguagem rust foi utilizada por demonstrar desempenho de execução e uso de memória

   
# Estratégia para Solução
   
   
   
1. Percorre vetor da esquerda para direita   
    - Vetor é percorrido começando do item 0 (esquerda) usando um loop `for`, o codigo deve verificar se a atual classificação lida é maior do que a anterior   
    - Quando a classificação atual é maior, a criança recebe um doce a mais que a criança anterior, quando não maior é dado apenas 1 doce   
    - A lógica acontece no primeiro loop do codigo    
   
       
   
2.  Percorre o vetor da direita pra esquerda   
  - Vetor é percorrido do último item para a esquerda usando loop `for`, dessa forma é garantido que a criança tenha tantos doces quanto a criança a direita com uma classficação maior   
  - é ajustado a quantidade de doces da criança atual para que seja o máximo entre o número atual de doces e o numéro de doces da criança a direita   
  - A lógica acontece no segundo loop do codigo   
   
   
- Antes de processar os loops, é verificado se o vetor tem tamanho menor ou igual a 1. Quando é verdadeiro, o próprio tamanho é retornado, já que não será preciso atribuir doces   
   
   
   
**Inicialização e Contagem Total**:   
- O algortimo inicaliza o vetor candies que armazena a quantidade de doces que é dada para cada criança   
- Usando a variável `total` fazemos a contagem de todos os doces que foram atribuidos que no fim é retornada como resultado final.    
   
**Utilização de Função Auxiliar** max   
- A função max faz o trabalho de calcular o máximo entre 2 numeros. Portanto, usamos para garantir que a quantidade destribuida seja feita corretamente durante passagem do vetor. Ela foi separada para manter o código mais legível.   
- No java é substituido um import equivalante: `Math.max`   
   
   
#    
   
```


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


```
## Imports   
   
```

import (
	"fmt"
	"log"
)

```
![image.png](files/image.png)    
```
public class Candy {
    public int candy(int[] ratings) {
        // Obtém o tamanho do array
        int n = ratings.length;

        // Se o tamanho do array for menor ou igual a 1, retorna o próprio tamanho
        if (n <= 1) {
            return n;
        }

        // Cria um array para armazenar a quantidade de doces atribuídos a cada criança
        int[] candies = new int[n];
        
        // Atribui 1 doce à primeira criança
        candies[0] = 1;

        // Loop para percorrer as classificações da esquerda para a direita
        for (int i = 1; i < n; i++) {
            // Se a classificação atual for maior que a anterior, atribui um doce a mais
            if (ratings[i] > ratings[i - 1]) {
                candies[i] = candies[i - 1] + 1;
            } else {
                // Se não, atribui apenas 1 doce
                candies[i] = 1;
            }
        }

        // Inicializa a variável total com a quantidade de doces da última criança
        int total = candies[n - 1];

        // Loop para percorrer as classificações da direita para a esquerda
        for (int i = n - 2; i >= 0; i--) {
            // Se o rating atual é maior que o próximo, ajusta para que tenha 1 a mais que o próximo
            if (ratings[i] > ratings[i + 1]) {
                // Se a próxima criança tem mais doces, nada muda
                candies[i] = Math.max(candies[i], candies[i + 1] + 1);
            }
            // Adiciona a quantidade de doces da criança atual ao total
            total += candies[i];
        }

        // Retorna a quantidade total de doces atribuídos
        return total;
    }
}

```
![image.png](files/image_v.png)    
```
// Importa a biblioteca de logging do Rust
use log::{info, debug};

impl Solution {
    pub fn candy(ratings: Vec<i32>) -> i32 {
        // Inicializa o sistema de logging
        env_logger::init();

        // Obtém o tamanho do vetor
        let n = ratings.len();

        // Log da entrada
        debug!("Classificações recebidas: {:?}", ratings);

        // Se o tamanho do vetor for menor ou igual a 1, retorna o próprio tamanho
        if n <= 1 {
            info!("Apenas uma criança ou nenhum. Retornando {} como a quantidade total de doces.", n);
            return n as i32;
        }

        // Cria um vetor para armazenar a quantidade de doces atribuídos a cada criança
        let mut candies = vec![0; n];

        // Atribui 1 doce à primeira criança
        candies[0] = 1;

        // Loop para percorrer as classificações da esquerda para a direita
        for i in 1..n {
            // Se a classificação atual for maior que a anterior, atribui um doce a mais
            if ratings[i] > ratings[i - 1] {
                candies[i] = candies[i - 1] + 1;
            } else {
                // Se não, atribui apenas 1 doce
                candies[i] = 1;
            }
        }

        // Inicializa a variável total com a quantidade de doces da última criança
        let mut total = candies[n - 1];

        // Loop para percorrer as classificações da direita para a esquerda
        for i in (0..n - 1).rev() {
            // Se o rating atual é maior que o próximo, ajusta para que tenha 1 a mais que o próximo
            if ratings[i] > ratings[i + 1] {
                // Se a próxima criança tem mais doces, nada muda
                candies[i] = std::cmp::max(candies[i], candies[i + 1] + 1);
            }
            // Adiciona a quantidade de doces da criança atual ao total
            total += candies[i];
        }

        // Log do resultado
        info!("Quantidade total de doces atribuídos: {}", total);

        // Retorna a quantidade total de doces atribuídos
        total
    }
}

```
## Imports   
```
[dependencies]
env_logger = "0.12"
log = "0.4"

```
![image.png](files/image_f.png)    
   
Os argumentos para resolver são um pouco diferentes do hackeRank, no hacker rank os argumento são um int n e um vetor, enquanto no leetcode existe apenas um argumento de entrada.    
   
algoritmo deve ser testado em máquinas diferentes (pelo menos duas), e o
tempo de execução deve ser analisado.   
