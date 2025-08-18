package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// getJsonContent makes a GET request to a URL for JSON
func getJsonContent(urlToRequest string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlToRequest, nil)

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

	fmt.Print(string(bodyBytes))
}
