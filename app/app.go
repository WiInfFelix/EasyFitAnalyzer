package app

import (
	"easyfitanalysis/analyser"
	"easyfitanalysis/logger"
	"easyfitanalysis/scraper"
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router          *mux.Router
	currentAverages map[string]map[int]float32
}

func (a *App) RefreshCurrentAverages(analysis map[string]map[int]float32) {
	log.Println("Refreshing in-memory values...")

	a.currentAverages = analysis
	return
}

func (a *App) analysis(writer http.ResponseWriter, request *http.Request) {

	jsonString, err := json.MarshalIndent(a.currentAverages, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(writer, string(jsonString))
}

func (a *App) Init() {

	a.Router = mux.NewRouter()
	a.InitRoutes()
	a.startSchedules()
	a.RefreshCurrentAverages(analyser.ReturnAnalysis())
}

func (a *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitRoutes() {

	fs := http.FileServer(http.Dir("./static/"))

	a.Router.HandleFunc("/analysisResults", a.analysis).Methods("GET")
	a.Router.PathPrefix("/").Handler(http.StripPrefix("", fs)).Methods("GET")
	a.Router.Use(logger.LogRequestHandler)

}

func (a *App) startSchedules() {
	log.Println("Starting scraping agent")

	s := gocron.NewScheduler(time.UTC)

	// build a channel to get return values and asign them to the cached map
	// https://stackoverflow.com/questions/65432808/return-output-data-from-gocron-task
	dataChan := make(chan interface{})

	s.Cron("1,15,30,45 8-23 * * 1-5").Do(scraper.ScrapePage, dataChan)
	s.Cron("1,15,30,45 10-20 * * 0,6").Do(scraper.ScrapePage, dataChan)

	go func() {
		for data := range dataChan {
			a.RefreshCurrentAverages(data.(map[string]map[int]float32))
		}
	}()

	s.StartAsync()
	log.Println("Applied schedules")
}
