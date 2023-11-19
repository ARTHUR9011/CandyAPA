public class Main {
    public static void main(String[] args) {
        Candy candyDistributor = new Candy();

        // Classificações de exemplo
        int[] ratings = {3, 1, 5, 2, 4, 6};

        // Chamada do método candy
        int totalCandies = candyDistributor.candy(ratings);

        System.out.println("Total de doces atribuídos: " + totalCandies);
    }
}
