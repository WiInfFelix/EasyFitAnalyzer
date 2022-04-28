# EasyFitAnalser

## What is it?

This web app shows an aggregated usage of the EasyFitnessStudio Lübeck by scraping the usage from their website every
fifteen minutes.

The data is then aggregated by weekday and hour and displayed in some bar charts.

The app uses Go as the backend with the following external dependencies:

* Colly: For scraping the data
* Gorilla/Mux: For routing
* Gocron: For scheduling the scraping jobs


The frontend is served by the Go server and uses Bootstrap (for making it look pretty) and ChartJS to plot the bar charts