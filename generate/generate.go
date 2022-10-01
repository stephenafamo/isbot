//go:build ignore

// This program downloads the crawler-user-agents.json file https://github.com/monperrus/crawler-user-agents
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	downloadList(
		ctx, "crawler-user-agents.json", "GET",
		"https://raw.githubusercontent.com/monperrus/crawler-user-agents/master/crawler-user-agents.json", nil,
	)

	downloadList(
		ctx, "user-agents-bots.txt", "POST",
		"https://user-agents.net/download",
		strings.NewReader("crawler=true&download=txt"),
	)
}

func downloadList(ctx context.Context, name, method, uri string, body io.Reader) {
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if body != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	die(err)

	rsp, err := http.DefaultClient.Do(req)
	die(err)
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(rsp.Body)
		die(err)

		fmt.Println(string(body))
		return
	}

	f, err := os.Create(name)
	die(err)
	defer f.Close()

	_, err = io.Copy(f, rsp.Body)
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
