package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// getJsonContent makes a GET request to a URL for JSON
func getJsonContent(urlToRequest string, prettyPrint bool) string {
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

	if prettyPrint {
		var prettyJson bytes.Buffer

		json.Indent(&prettyJson, bodyBytes, "", "  ")

		return prettyJson.String()
	} else {
		return string(bodyBytes)
	}
}
