package main

import (
	"fmt"
	"log"

	"github.com/b-open/jobbuzz/internal/config"
	"github.com/b-open/jobbuzz/pkg/scraper"
	"github.com/b-open/jobbuzz/pkg/service"
)

func main() {

	configuration, err := config.LoadConfig("../../")

	if err != nil {
		log.Fatal("Fail to load db config", err)
	}

	db, err := configuration.GetDb()

	if err != nil {
		log.Fatal("Fail to get db connection", err)
	}

	service := service.Service{DB: db}

	// Scrape JobCenter

	fmt.Println("Fetching jobs from JobCenter")
	jobs := scraper.ScrapeJobcenter()

	for _, job := range jobs {
		_, err = service.CreateJob(&job)

		if err != nil {
			fmt.Println("Fail to create job")
		}
	}

	// Scrape Bruneida

	fmt.Println("Fetching jobs from Bruneida")
	jobs = scraper.ScrapeBruneida()

	for _, job := range jobs {
		_, err = service.CreateJob(&job)

		if err != nil {
			fmt.Println("Fail to create job")
		}
	}
}
