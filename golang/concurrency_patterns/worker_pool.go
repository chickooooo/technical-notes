/*
This Go program implements a simple Worker Pool to concurrently fetch the titles of multiple web pages.
It defines a Job struct representing a web page URL to process, and a JobResponse struct to hold the
result of fetching the page title.

The Worker function continuously processes jobs from a channel, fetching the HTML title for each URL,
and sends the results to a result channel. The WorkerPool function initializes a list of URLs, spawns
a fixed number of worker goroutines, sends all jobs to the workers, and collects and prints the results.

This example demonstrates concurrency in Go using goroutines and channels, as well as basic HTML parsing
using the "golang.org/x/net/html" package.
*/

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type Job struct {
	ID  int
	URL string
}

type JobResponse struct {
	Job   *Job
	Title string
	Err   error
}

func (j *Job) Title() (string, error) {
	// Make GET request
	resp, err := http.Get(j.URL)
	if err != nil {
		return "", err
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Parse HTML body
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	// Extract title tag
	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if title == "" {
		return "", fmt.Errorf("no title element found")
	}
	return title, nil
}

func Worker(id int, jobCh <-chan *Job, resultCh chan<- *JobResponse) {
	fmt.Printf("Worker %d ready\n", id)

	// Keep receiving jobs while jobs channel is open
	for j := range jobCh {
		fmt.Printf("Worker %d picked up job %d\n", id, j.ID)
		t, err := j.Title()
		fmt.Printf("Worker %d finished job %d\n", id, j.ID)

		// Send response to result channel
		resultCh <- &JobResponse{
			Job:   j,
			Title: t,
			Err:   err,
		}
	}
}

func WorkerPool() {
	urls := []string{
		"https://chatgpt.com/",
		"https://blinkit.com/",
		"https://gobyexample.com/",
		"https://aliexpress.com/",
		"https://apple.com/",
		"https://lg.com/",
		"https://google.com/",
		"https://alibaba.com/",
		"https://xbox.com/",
		"https://youtube.com/",
		"https://samsung.com/",
		"https://hotmail.com/",
		"https://zomato.com/",
		"https://amazon.com/",
		"https://microsoft.com/",
		"https://myntra.com/",
		"https://goibibo.com/",
		"https://zepto.com/",
		"https://lenskart.com/",
	}
	workers := 3

	// Job channel will hold all the jobs
	jobCh := make(chan *Job, len(urls))
	// Result channel will hold the results
	resultCh := make(chan *JobResponse, len(urls))

	// Spawn workers
	for i := range workers {
		go Worker(i+1, jobCh, resultCh)
	}

	// Send jobs
	for i, url := range urls {
		jobCh <- &Job{
			ID:  i + 1,
			URL: url,
		}
	}
	close(jobCh)

	// Print results
	for i := 0; i < len(urls); i++ {
		res := <-resultCh
		fmt.Println("\n-------------------------------------")
		fmt.Printf("Received result for job: %d, %s\n", res.Job.ID, res.Job.URL)
		fmt.Printf("Title: %s, Error: %v\n", res.Title, res.Err)
		fmt.Println("-------------------------------------\n")
	}
}

func main() {
	// Start worker pool
	WorkerPool()
}
