package saucenao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// A SaucenaoClient is used to make requests to the SauceNao API.
type SaucenaoClient struct {
	APIKey            string
	MinimumSimilarity int
	DatabaseBitmask   int
	AmountOfResults   int
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

	//// To allow for other websites, add their fields here. ////

}

func New(APIKey string) (s *SaucenaoClient) {
	s = &SaucenaoClient{
		APIKey:            APIKey,
		MinimumSimilarity: 80,
		DatabaseBitmask:   999,
		AmountOfResults:   1,
	}

	return
}

// Makes a GET request to the SauceNao API given an url.
// It is the responsibility of the user to make sure that this url leads to an image.
func (s SaucenaoClient) FromURL(imageurl string) (res SaucenaoResult, err error) {
	parsedUrl, _ := url.Parse("http://saucenao.com/search.php")
	queryString := parsedUrl.Query()

	jsonOutput := "2"
	queryString.Set("output_type", jsonOutput)
	queryString.Set("numres", strconv.Itoa(s.AmountOfResults))
	queryString.Set("minsim", strconv.Itoa(s.MinimumSimilarity))
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
