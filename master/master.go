package master

import (
	"context"
	"fmt"
	"strings"
)

type Master struct {
	workerClients     []WorkerClient
	idleWorkerClients chan WorkerClient
	linksCh           chan []string
}

func (m *Master) ExploreWebsite(ctx context.Context, siteURL string) {
	done := ctx.Done()

	// Prevent channel from blocking on the with the initial site URL
	m.linksCh = make(chan []string, 1)
	linkCh := make(chan string)
	visited := make(map[string]bool)

	m.linksCh <- []string{siteURL}
	count := 1

	for {
		select {
		case links := <-m.linksCh:
			if len(links) > 0 {
				fmt.Printf("%s\n", strings.Join(links, "\n"))
			}
			count--
			for _, link := range links {
				if visited[link] {
					continue
				}
				count++
				visited[link] = true
				go func(link string) {
					linkCh <- link
				}(link)
			}
			if count == 0 {
				fmt.Printf("Finish exploring %s\n", siteURL)
				return
			}
		case link := <-linkCh:
			go func(link string) {
				// process the current site there is idle workerClients
				worker := <-m.idleWorkerClients
				err := worker.FetchLinks(ctx, link)
				if err != nil {
					m.idleWorkerClients <- worker
				}
			}(link)
		case <-done:
			return
		}
	}
}

func (m *Master) FinishExtractingLinks(workerID int, links []string) {
	m.linksCh <- links
	go func() {
		m.idleWorkerClients <- m.workerClients[workerID]
	}()
}

func (m *Master) RegisterWorker(workerClient WorkerClient) int {
	workID := len(m.workerClients)
	m.workerClients = append(m.workerClients, workerClient)
	go func() {
		m.idleWorkerClients <- workerClient
	}()
	return workID
}

func newMaster() *Master {
	return &Master{
		idleWorkerClients: make(chan WorkerClient),
	}
}
