package main

import (
	"concurrency/structs"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	data, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal("File can't be readed error :", err)
	}
	stats := make(map[string][]structs.Visit)
	err = json.Unmarshal(data, &stats)
	if err != nil {
		log.Fatalf("An error was returned while parsing the Json file %v", err)
	}
	wait := &sync.WaitGroup{}
	wait.Add(len(stats))
	inChan := make(chan structs.Task, 10)
	outChan := make(chan structs.DailyStat, len(stats))
	for i := 0; i < len(stats); i++ {
		go worker(inChan, outChan, i, wait)
	}
	for date, visits := range stats {
		inChan <- structs.Task{
			Date:   date,
			Visits: visits,
		}
	}
	close(inChan)
	wait.Wait()
	close(outChan)
	done := make([]structs.DailyStat, 0, len(stats))
	for out := range outChan {
		done = append(done, out)
	}
	res, err := json.Marshal(done)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("result.json", res, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DONE")
}

func worker(in chan structs.Task, out chan structs.DailyStat, workerId int, wait *sync.WaitGroup) {
	for received := range in {
		m := make(map[string]int)
		for _, v := range received.Visits {
			m[v.Page]++
		}
		out <- structs.DailyStat{
			Date:   received.Date,
			Bypage: m,
		}
		fmt.Printf("worker %d finished task \n", workerId)
	}
	log.Println("workers done")
	wait.Done()
}
