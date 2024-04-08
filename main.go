package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var ()

func main() {
	reader := bufio.NewReader(os.Stdin)
	config, _ := reader.ReadString('\n')
	config = strings.TrimSuffix(config, "\n")
	println(config)
	input := strings.Split(config, " ")

	//Number of requests
	N, _ := strconv.Atoi(input[0])
	//rate limit
	R, _ := strconv.Atoi(input[1])

	cached := newRateLimitChecker(R)

	for range N {
		request, _ := reader.ReadString('\n')
		request = strings.TrimSuffix(request, "\n")
		requestTime, err := time.Parse(time.RFC3339, request)
		if err != nil {
			log.Println("Request invalid: ", err)
			continue
		}

		//Remove expired request log
		cached.removeExpiredLog(requestTime)

		//check request overhead
		if cached.isRequestOverflow() {
			fmt.Printf("request [%s]: false\n", request)
			continue
		} else {
			fmt.Printf("request [%s]: true\n", request)
		}

		//add request log
		cached.addLog(requestTime)

	}

	log.Println("Program exited")

}
