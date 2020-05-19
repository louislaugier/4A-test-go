package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/louislaugier/4A-test-go/content"
)

type platformFeed struct {
	Articles       []*content.Article
	Videos         []*content.Video
	Moocs          []*content.Mooc
	Certifications []*content.Certification
}

func importFromCSV(filePath string) ([]*content.Item, content.Err) {
	f, err := os.Open("trajets.csv")
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, content.Err{
			StatusCode: 500,
			Message:    err,
		}
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Println("Error reading file:", err)
		return nil, content.Err{
			StatusCode: 500,
			Message:    err,
		}
	}
	var items []*content.Item
	for indexRecords, record := range records {
		log.Println(record)
		if indexRecords > 0 {
			items = append(items, &content.Item{
				// A continuer ici
				// ID: uuid.New()
			})
		}
	}
	return items, content.Err{
		StatusCode: 200,
		Message:    nil,
	}
}

func (p *platformFeed) itemsGET(items []*content.Item, limit int) {
	for _, item := range items {
		if item.Type == "Article" {
			p.Articles = append(p.Articles, item.Article)
		} else if item.Type == "Video" {
			p.Videos = append(p.Videos, item.Video)
		} else if item.Type == "Mooc" {
			p.Moocs = append(p.Moocs, item.Mooc)
		} else if item.Type == "Certification" {
			p.Certifications = append(p.Certifications, item.Certification)
		}
	}
}

func main() {
	var myFeed platformFeed
	items, err := importFromCSV("content.csv")
	if err.StatusCode != 200 {
		log.Println(err)
	}
	myFeed.itemsGET(items, 7)
	// wg := sync.WaitGroup{}
	// for _, b := range myFeed.Articles {
	// 	wg.Add(1)
	//
	// 	go func(wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 	}(&wg)
	// }
	// wg.Wait()
	// log.Println("Done!")
}
