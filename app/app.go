package app

import (
	"easyfitanalysis/analyser"
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

func (a *App) RefreshCurrentAverages() {
	log.Println("Refreshing in-memory values...")
	a.currentAverages = analyser.ReturnAnalysis()
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
	a.RefreshCurrentAverages()
}

func (a *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitRoutes() {

	fs := http.FileServer(http.Dir("./static/"))

	a.Router.HandleFunc("/analysisResults", a.analysis).Methods("GET")
	a.Router.PathPrefix("/").Handler(http.StripPrefix("", fs)).Methods("GET")

}

func (a *App) startSchedules() {
	log.Println("Starting scraping agent")

	s := gocron.NewScheduler(time.UTC)

	s.Cron("1,15,30,45 8-23 * * 1-5").Do(scraper.ScrapePage)
	s.Cron("1,15,30,45 10-20 * * 0,6").Do(scraper.ScrapePage)

	s.StartAsync()
	log.Println("Applied schedules")
}
