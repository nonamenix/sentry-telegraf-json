package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Project struct {
	Slug  string      `json:"slug"`
	Stats [][]float64 `json:"stats"`
	Team  struct {
		ID   string `json:"id"`
		Slug string `json:"slug"`
	} `json:"team"`
}

type Config struct {
	sentryURL          string
	organization       string
	statsPeriod        string
	query              string
	authorizationToken string
}

type Metric struct {
	Project string `json:"project"`
	State   int    `json:"state"`
}

func fetchErrorsFromSentry(config Config) ([]Project, error) {
	requestURL := config.sentryURL + "/api/0/organizations/" + config.organization + "/projects/?statsPeriod=24h&query=" + config.query

	client := &http.Client{}
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+config.authorizationToken)

	resp, err := client.Do(request)
	if err != nil && resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects = []Project{}

	decodeErr := json.NewDecoder(resp.Body).Decode(&projects)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return projects, nil
}

func sentry2telegraf(config Config) {
	projects, _ := fetchErrorsFromSentry(config)
	var metrics []Metric
	for _, project := range projects {
		stat := project.Stats[len(project.Stats)-1]
		errorsCount := int(stat[1])
		metrics = append(metrics, Metric{Project: project.Slug, State: errorsCount})
	}
	json_metrics, _ := json.Marshal(metrics)
	fmt.Println(string(json_metrics))
}

func main() {
	var (
		sentryURL          = flag.String("sentry-url", "https://sentry.io", "The sentry url")
		organization       = flag.String("organization", "XXX", "Organization name in sentry")
		statsPeriod        = flag.String("stats-period", "24h", "Sentry stats period")
		query              = flag.String("query", "", "Sentry query for projects filtering")
		authorizationToken = flag.String("token", "", "Sentry API authorization token")
	)
	flag.Parse()

	var config = Config{
		sentryURL:          *sentryURL,
		organization:       *organization,
		query:              *query,
		statsPeriod:        *statsPeriod,
		authorizationToken: *authorizationToken,
	}
	sentry2telegraf(config)
}
