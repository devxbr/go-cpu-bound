package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// StressCPU consome uma porcentagem específica da CPU por um tempo determinado.
// percent: porcentagem de CPU a ser consumida (0-100).
// duration: duração do estresse (em segundos).
func StressCPU(percent int, duration int) {
	if percent <= 0 || percent > 100 {
		panic("O percentual deve estar entre 1 e 100")
	}

	// Calcula o tempo de trabalho e espera com base na porcentagem
	workTime := time.Duration(percent) * time.Millisecond
	sleepTime := time.Duration(100-percent) * time.Millisecond

	// Obtém o número de CPUs disponíveis
	numCPU := runtime.NumCPU()

	fmt.Printf("running: %d%% da CPU por %d segundos usando %d núcleos.\n", percent, duration, numCPU)

	// Canal para sincronizar a duração
	stop := make(chan bool)
	for i := 0; i < numCPU; i++ {
		go func() {
			for {
				select {
				case <-stop:
					return
				default:
					// lock busy loop
					start := time.Now()
					for time.Since(start) < workTime {
					}
					time.Sleep(sleepTime)
				}
			}
		}()
	}

	// Aguarda o tempo de duração do estresse
	time.Sleep(time.Duration(duration) * time.Second)

	// Para todas as goroutines
	close(stop)

	fmt.Println("Teste concluído.")
}

func main() {
	percent := flag.Int("percent", 50, "Porcentagem da CPU a ser consumida (1-100)")
	duration := flag.Int("duration", 10, "Duração do teste em segundos")
	flag.Parse()
	if *percent < 1 || *percent > 100 {
		fmt.Println("o percentual deve estar entre 1 e 100.")
		return
	}
	if *duration < 1 {
		fmt.Println("a duração deve ser maior que 0.")
		return
	}

	StressCPU(*percent, *duration)
}
