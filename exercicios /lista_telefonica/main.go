package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um mapa para armazenar a lista telefônica
	listaTelefonica := map[string]string{
		"João":      "123-456-789",
		"Maria":     "987-654-321",
		"Carlos":    "555-123-456",
		"Ana":       "111-222-333",
		"Paulo":     "444-555-666",
		"Juliana":   "777-888-999",
		"Fernando":  "666-555-444",
		"Roberta":   "321-654-987",
		"Gustavo":   "999-888-777",
		"Amanda":    "456-789-012",
		"Lucas":     "111-333-555",
		"Renata":    "444-666-888",
		"Diego":     "999-555-111",
		"Sandra":    "123-789-456",
		"Henrique":  "987-654-321",
		"Tatiane":   "777-333-888",
		"Ricardo":   "555-111-777",
		"Camila":    "321-789-456",
		"Alexandre": "888-444-999",
	}

	// Iniciando a medição do tempo
	startTime := time.Now()

	// Acessando número de telefone específico
	// nomeBusca := "Gustavo"
	// if numero, existe := listaTelefonica[nomeBusca]; existe {
	// 	fmt.Printf("O número de telefone de %s é %s.\n", nomeBusca, numero)
	// } else {
	// 	fmt.Printf("%s não está na lista telefônica.\n", nomeBusca)
	// }

	numeroDaBusca := "Gustavo"
	if numero, existe := listaTelefonica[numeroDaBusca]; existe {
		fmt.Printf("O número de telefone de %s é %s.\n", numeroDaBusca, numero)
	} else {
		fmt.Printf("%s não está na lista telefônica.\n", numeroDaBusca)
	}
	// Finalizando a medição do tempo
	endTime := time.Now()

	// Calculando e imprimindo o tempo de execução
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Tempo de execução: %s\n", elapsedTime)
}




accessKey=AKIAYAE7CLM6GFCPZ4ER
accessKey=AKIAYAE7CLM6H54CKEJ3
accessKey=AKIAYAE7CLM6KF7RXNWV
accessKey=AKIAYAE7CLM6BPQAEIFK
accessKey=AKIAYAE7CLM6DC2SA7XB
accessKey=AKIAYAE7CLM6EN7T5C6J