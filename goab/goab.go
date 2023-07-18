package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	//args
	var concurrency = flag.Int("c", 1, "Concurrency")
	var totalRequests = flag.Int("n", 1, "Number of Requests")
	var keepAlive = flag.Bool("k", false, "Keep-Alive")

	flag.Parse()

	url := flag.Arg(0)

	//error
	if *concurrency <= 0 || *totalRequests <= 0 || url == "" || *totalRequests < *concurrency {
		log.Fatal("Bad argument")
	}

	var wg sync.WaitGroup
	wg.Add(*concurrency)

	responseTimes := make(chan time.Duration, *totalRequests)
	tErrorCount := 0
	startTime := time.Now()

	for i := 0; i < *concurrency; i++ {

		go func() { //anonymous func because of the arguments
			defer wg.Done()

			client := &http.Client{}
			client.Transport = &http.Transport{

				DisableKeepAlives: !*keepAlive,
			}
			/*requ, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println(err)
			}*/
			//requ.Proto = "HTTP/1.0"
			//requ.ProtoMajor = 1
			//requ.ProtoMinor = 0

			for j := 0; j < *totalRequests / *concurrency; j++ {
				timeE := time.Now()

				resp, err := client.Get(url)
				if err != nil {
					log.Println("Error: ", err)
					tErrorCount++
					continue
				}
				resp.Body.Close()

				responseTimes <- time.Since(timeE)

				if resp.StatusCode != 200 {
					tErrorCount++
				}
			}

		}()

	}

	wg.Wait()
	close(responseTimes)

	elapsed := time.Since(startTime)
	tps := float64(*totalRequests) / elapsed.Seconds()
	tLatency := time.Duration(0)
	errorPercentage := (float64(tErrorCount) / float64(*totalRequests)) * 100
	successCount := 0

	for rt := range responseTimes {
		tLatency += rt
		successCount++
	}
	if successCount == 0 {
		successCount = 1
	}
	avgLatency := tLatency / time.Duration(successCount)

	fmt.Println("")
	fmt.Printf("Transactions Per Second (TPS): %.2f\n", tps)
	fmt.Printf("Average Latency: %v\n", avgLatency)
	fmt.Printf("Errored Responses: %d (%.2f%%)\n", tErrorCount, errorPercentage)
}

