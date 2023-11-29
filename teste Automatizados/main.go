package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ceps := []string{"20040010", "20040010", "70002-900", "04101-300"}

	canal := make(chan string, len(ceps))

	for _, cep := range ceps {
		wg.Add(1)
		go getAddress(cep, canal, &wg)
	}

	go func() {
		wg.Wait()
		close(canal)
	}()

	for message := range canal {
		fmt.Println(message)
	}
}

func getAddress(cep string, canal chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	address, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Println(err)
		canal <- fmt.Sprintf("Error fetching data for CEP %s", cep)
		return
	}
	defer address.Body.Close()

	body, err := ioutil.ReadAll(address.Body)
	if err != nil {
		fmt.Println(err)
		canal <- fmt.Sprintf("Error reading data for CEP %s", cep)
		return
	}

	canal <- string(body)
}
