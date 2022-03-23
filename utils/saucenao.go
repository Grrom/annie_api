package saucenao

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SaucenaoClient struct {
	APIKey          string
	DatabaseBitmask int
	AmountOfResults int
	OutputType      int
}

// The result struct, the result of the query is directly parsed into this struct.
type SaucenaoResult struct {
	Header SaucenaoHeader    `json:"header"`
	Data   []SaucenaoResults `json:"results"`
}

type SaucenaoHeader struct {
	UserId            string                         `json:"user_id"`
	AccountType       string                         `json:"account_type"`
	ShortLimit        string                         `json:"short_limit"`
	LongLimit         string                         `json:"long_limit"`
	LongRemaining     int                            `json:"long_remaining"`
	ShortRemaining    int                            `json:"short_remaining"`
	Status            int                            `json:"status"`
	ResultsRequested  int                            `json:"results_requested"`
	Index             map[string]SaucenaoHeaderIndex `json:"index"`
	SearchDepth       string                         `json:"search_depth"`
	MinimumSimilarity float32                        `json:"minimum_similarity"`
	QueryImageDisplay string                         `json:"query_image_display"`
	QueryImage        string                         `json:"query_image"`
	ResultsReturned   int                            `json:"results_returned"`
}

type SaucenaoHeaderIndex struct {
	Status   int `json:"status"`
	ParentId int `json:"parent_id"`
	Id       int `json:"id"`
	Results  int `json:"results"`
}

type SaucenaoResults struct {
	Header SaucenaoResultsHeader `json:"header"`
	Data   SaucenaoResultsData   `json:"data"`
}

type SaucenaoResultsHeader struct {
	Similarity string `json:"similarity"`
	Thumbnail  string `json:"thumbnail"`
	IndexId    int    `json:"index_id"`
	IndexName  string `json:"index_name"`
}

type SaucenaoResultsData struct {
	ExtUrls []string `json:"ext_urls"`
	Title   string   `json:"title"`
	Source  string   `json:"source"`

	// it's just a general struct, it doesn't get all of the data from all other sources
}

type Sauce struct {
	Link      string  `json:"link"`
	Accuracy  float64 `json:"accuracy"`
	Title     string  `json:"title"`
	Thumbnail string  `json:"thumbnail"`
}

func saucenaoToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error: .Err:%s", err)
	}
	return os.Getenv("SAUCENAO_TOKEN")
}

func New() (s *SaucenaoClient) {
	s = &SaucenaoClient{
		APIKey:          saucenaoToken(),
		DatabaseBitmask: 999,
		AmountOfResults: 1,
		OutputType:      2,
	}
	return
}

// Makes a GET request to the SauceNao API given an url.
// It is the responsibility of the user to make sure that this url leads to an image.
func (s SaucenaoClient) FromURL(imageurl string) (res SaucenaoResult, err error) {
	parsedUrl, _ := url.Parse("http://saucenao.com/search.php")
	queryString := parsedUrl.Query()

	queryString.Set("output_type", strconv.Itoa(s.OutputType))
	queryString.Set("numres", strconv.Itoa(s.AmountOfResults))
	queryString.Set("dbmask", strconv.Itoa(s.DatabaseBitmask))
	queryString.Set("api_key", s.APIKey)
	queryString.Set("url", imageurl)

	parsedUrl.RawQuery = queryString.Encode()

	var req *http.Request
	req, err = http.NewRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return
	}

	var resp *http.Response
	client := http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &res)

	return
}
