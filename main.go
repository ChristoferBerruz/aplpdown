/*
This is a program that crawls a given URL and downloads all the files and folders found on that page.
It basically keeps the same structure as the original website.

This was created based on my need to download a lot of files from the CS252 course website.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func isFileLink(link string) bool {
	base := path.Base(link)
	return strings.Contains(base, ".")
}

func CrawlAndDownload(url, destination string, dryRun bool, visited map[string]bool) {
	if visited[url] {
		return
	}
	visited[url] = true

	fmt.Printf("Crawling: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("An error occurred while crawling %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch %s: %s\n", url, resp.Status)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Failed to parse HTML at %s: %v\n", url, err)
		return
	}

	var fileLinks, folderLinks []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val
					// Handle relative URLs
					switch {
					case strings.HasPrefix(href, "./"):
						href = href[2:]
					case strings.HasPrefix(href, "../"):
						href = href[3:]
					case strings.HasPrefix(href, "/"):
						href = href[1:]
					}
					fullURL := url
					if strings.HasSuffix(url, "/") {
						fullURL += href
					} else {
						fullURL += "/" + href
					}
					if fullURL == url || visited[fullURL] {
						continue
					}
					if isFileLink(fullURL) {
						fileLinks = append(fileLinks, fullURL)
					} else {
						folderLinks = append(folderLinks, fullURL)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for _, fileLink := range fileLinks {
		fileName := path.Base(fileLink)
		filePath := filepath.Join(destination, fileName)
		if dryRun {
			fmt.Printf("[Dry-run] Would download: %s to %s\n", fileLink, filePath)
		} else {
			fmt.Printf("Downloading: %s to %s\n", fileLink, filePath)
			fileResp, err := http.Get(fileLink)
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", fileLink, err)
				continue
			}
			defer fileResp.Body.Close()
			out, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Failed to create file %s: %v\n", filePath, err)
				continue
			}
			_, err = io.Copy(out, fileResp.Body)
			out.Close()
			if err != nil {
				fmt.Printf("Failed to save file %s: %v\n", filePath, err)
			} else {
				fmt.Printf("Saved %s to %s\n", fileLink, filePath)
			}
		}
	}

	for _, folderLink := range folderLinks {
		folderName := path.Base(folderLink)
		folderPath := filepath.Join(destination, folderName)
		if dryRun {
			fmt.Printf("[Dry-run] Would create: %s\n", folderPath)
		} else {
			os.MkdirAll(folderPath, os.ModePerm)
		}
		CrawlAndDownload(folderLink, folderPath, dryRun, visited)
	}
}
func main() {
	dryRun := flag.Bool("dry-run", false, "If set, only prints the files to be downloaded without downloading them.")
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("Usage: cli <url> <destination> [-dry-run]")
		os.Exit(1)
	}
	fmt.Println("Dry run mode:", *dryRun)
	url := flag.Arg(0)
	destination := flag.Arg(1)

	os.MkdirAll(destination, os.ModePerm)
	visited := make(map[string]bool)
	CrawlAndDownload(url, destination, *dryRun, visited)
}
