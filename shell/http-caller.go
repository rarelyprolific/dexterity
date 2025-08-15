package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func makeExampleHttpCall() {
	// Example HTTP call
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.rarelyprolific.co.uk", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	rawHTML := string(bodyBytes)

	fmt.Println("HTML HTTP Response Body:", rawHTML)
}
