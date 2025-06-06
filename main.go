package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Station struct {
	count     int
	totalTemp float64
	minTemp   float64
	maxTemp   float64
}

func main() {
	// Captura o estado inicial da memória
	var m runtime.MemStats
	start := time.Now()
	runtime.ReadMemStats(&m)

	// Abre o arquivo
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Cria um novo scanner
	scanner := bufio.NewScanner(file)

	stations := map[string]Station{}
	linesRead := 0

	// Lê o arquivo linha por linha
	for scanner.Scan() {
		linesRead++
		line := scanner.Text()

		parts := strings.Split(line, ";")
		station := parts[0]
		temp, err := strconv.ParseFloat(parts[1], 64)

		if err != nil {
			continue
		}

		if _, ok := stations[station]; ok {
			payload := stations[station]
			payload.count++
			payload.totalTemp += temp
			payload.minTemp = math.Min(payload.minTemp, temp)
			payload.maxTemp = math.Max(payload.maxTemp, temp)
			stations[station] = payload
		} else {
			stations[station] = Station{
				count:     1,
				totalTemp: temp,
				minTemp:   temp,
				maxTemp:   temp,
			}
		}

	}

	// Verifica se houve erro durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro durante a leitura do arquivo: %v\n", err)
	}

	// Ordena as estações por ordem alfabética
	stationNames := make([]string, 0, len(stations))
	for name := range stations {
		stationNames = append(stationNames, name)
	}
	sort.Strings(stationNames)

	// Imprime os resultados em ordem alfabética
	fmt.Println()
	for _, name := range stationNames {
		station := stations[name]

		avg := station.totalTemp / float64(station.count)

		fmt.Printf("%s=%.1f/%.1f/%.1f\n",
			name, station.minTemp, avg, station.maxTemp)
	}

	// Captura o estado final da memória
	runtime.ReadMemStats(&m)

	fmt.Println()
	fmt.Println("Tempo de execução: ", time.Since(start))
	fmt.Printf("Número de linhas lidas: %d\n", linesRead)
	fmt.Printf("Memória final: %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("Total de alocações: %v\n", m.TotalAlloc)
	fmt.Printf("Número de coletas de lixo: %v\n", m.NumGC)
}
