package worker

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Worker struct to encapsulate all the worker variables and methods
type Worker struct {
	id           int
	secret       string
	masterClient MasterClient
}

// ExtractLinks calls the worker method extractLinks and after informs master FinishExtractingLinks
func (w Worker) ExtractLinks(secret string, sourceURL string) error {
	if secret != w.secret {
		return errors.New("incorrect credential")
	}
	links, err := extractLinks(sourceURL)
	if len(links) > 0 {
		fmt.Printf("%s\n", strings.Join(links, "\n"))
	}

	if err != nil {
		return w.masterClient.FinishExtractingLinks(w.id, []string{})
	}
	return w.masterClient.FinishExtractingLinks(w.id, links)
}

// Init connects to master and register itself to the master as Idle worker.
func (w *Worker) Init(ip string, port int) error {
	err := w.masterClient.Connect()
	if err != nil {
		return err
	}

	id, err := w.masterClient.RegisterWorker(ip, port, w.secret)
	if err != nil {
		return err
	}
	fmt.Printf("Registered with master: id(%d)\n", id)
	w.id = id
	return nil
}

func extractLinks(sourceURL string) ([]string, error) {
	var links []string
	res, err := http.Get(sourceURL)
	if err != nil {
		return nil, err
	}
	tokenizer := html.NewTokenizer(res.Body)
	defer res.Body.Close()
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken:
			tk := tokenizer.Token()
			if tk.Data != "a" {
				continue
			}
			val, err := findAttr(tk.Attr, "href")
			if err != nil {
				continue
			}
			links = append(links, val)
		}
	}
}

func findAttr(attrs []html.Attribute, key string) (string, error) {
	for _, attr := range attrs {
		if attr.Key != key {
			continue
		}
		return attr.Val, nil
	}
	return "", errors.New("attribute not found")
}

func newWorker(masterClient MasterClient, secret string) Worker {
	return Worker{masterClient: masterClient, secret: secret}
}
