use log::{info, debug};
use sys_info;

struct Solution;

impl Solution {
    pub fn main() {
        // memória no início do programa
        let mem_info = sys_info::mem_info().unwrap();

        // Configuração do sistema de logging
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

        // Cria um vetor de classificações
        let ratings = vec![1, 2,3,2,1,4];

        // Chama a função candy
        let total = Solution::candy(ratings);

        // Imprime a quantidade total de doces
        println!("Quantidade total de doces atribuídos: {}", total);

        // Imprime a memória total (em KB)
        println!("Memória total: {} KB", mem_info.total);

        // Imprime a memória livre (em KB)
        println!("Memória livre: {} KB", mem_info.free);

        // Imprime a memória usada (em MB)
        println!("Memória usada: {:.2} MB", mem_info.total as f64 / 1024.0 / 1024.0);
        println!("Memória Livre  {:.2} MB", mem_info.free as f64 / 1024.0 / 1024.0 )

    }

    pub fn candy(ratings: Vec<i32>) -> i32 {
        let start_time = std::time::Instant::now();

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
                info!("Ajustando a quantidade de doces da criança {} para {}", i, candies[i]);
            }

            // Adiciona a quantidade de doces da criança atual ao total
            total += candies[i];
        }

        let elapsed_time = start_time.elapsed();
        info!("Tempo gasto na função candy: {:?}", elapsed_time);
        total
    }
}

fn main() {
    Solution::main();
}