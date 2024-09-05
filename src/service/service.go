package service

import (
	"fmt"
	"sync"
)

type Location struct {
	Title string `json:"title"`
	WoeID int    `json:"woeid"`
}

type City struct {
	Name string
	Rep  string
}

func Somar(result chan float64, primeiro float64, segundo float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- primeiro + segundo
}

func Dividir(result chan float64, primeiro float64, segundo float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- primeiro / segundo
}

func Multiplicar(result chan float64, primeiro float64, segundo float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- primeiro * segundo
}

func Execute() {
	// channel to comunicate in goroutines
	resultChannel := make(chan float64)
	resultChannel2 := make(chan float64)

	// Usando WaitGroup para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Adiciona o número de cidades ao WaitGroup
	wg.Add(3)

	// Lança uma goroutine para cada cidade
	go Somar(resultChannel, 10, 2.32, &wg)
	go Dividir(resultChannel, 10, 2.22333, &wg)
	go Multiplicar(resultChannel2, 1, 2.01, &wg)

	// Goroutine para fechar o canal após todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(resultChannel)
		close(resultChannel2)
	}()

	wgReceiver := sync.WaitGroup{}
	wgReceiver.Add(2)

	// Goroutine para receber dados de resultChannel
	go func() {
		defer wgReceiver.Done()
		for result := range resultChannel {
			fmt.Println("Resultado do resultChannel:", result)
		}
	}()

	go func() {
		defer wgReceiver.Done()
		for result := range resultChannel2 {
			fmt.Println("Resultado do resultChannel2:", result)
		}
	}()

	wgReceiver.Wait()
}
