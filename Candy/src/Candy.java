import java.util.Arrays;
import java.util.logging.*;

public class Candy {
    private static final Logger LOGGER = Logger.getLogger(Candy.class.getName());

    static {
        LOGGER.setUseParentHandlers(false);
        Handler[] handlers = LOGGER.getHandlers();
        for (Handler handler : handlers) {
            LOGGER.removeHandler(handler);
        }

        ConsoleHandler handler = new ConsoleHandler();
        handler.setFormatter(new SimpleFormatter() {
            @Override
            public synchronized String format(LogRecord lr) {
                return lr.getMessage() + System.getProperty("line.separator");
            }
        });
        LOGGER.addHandler(handler);
    }

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
                LOGGER.info("Criança na posição " + i + " recebeu " + candies[i] + " doce(s).");
            } else {
                // Se não, atribui apenas 1 doce
                candies[i] = 1;
                LOGGER.info("Criança na posição " + i + " recebeu " + candies[i] + " doce(s).");
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
            LOGGER.info("Criança na posição " + i + " recebeu " + candies[i] + " doce(s).");
        }

        // Log das quantidades de doces atribuídos a cada criança
        LOGGER.info("Quantidade de doces atribuídos a cada criança: " + Arrays.toString(candies));
        LOGGER.info("Total de doces atribuídos: " + total);

        // Retorna a quantidade total de doces atribuídos
        return total;
    }
}