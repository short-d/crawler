package master

import (
	"context"
	"fmt"
	"strings"
)

type Master struct {
	workers     []WorkerClient
	idleWorkers chan WorkerClient
	linksCh     chan []string
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
				// process the current site there is idle workers
				worker := <-m.idleWorkers
				err := worker.FetchLinks(ctx, link)
				if err != nil {
					m.idleWorkers <- worker
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
		m.idleWorkers <- m.workers[workerID]
	}()
}

func (m *Master) RegisterWorker(worker WorkerClient) int {
	workID := len(m.workers)
	m.workers = append(m.workers, worker)
	go func() {
		m.idleWorkers <- worker
	}()
	return workID
}

func newMaster() *Master {
	return &Master{
		idleWorkers: make(chan WorkerClient),
	}
}
