package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)

const apiKey = "43c9eda0-c176-4a3e-9034-3cc68ccaf407"

type (
	SearchResponseWrapper struct {
		Response SearchResponse `json:"response"`
	}

	SearchResponse struct {
		Status      string    `json:"status"`
		UserTier    string    `json:"userTier"`
		Total       int       `json:"total"`
		StartIndex  int       `json:"startIndex"`
		PageSize    int       `json:"pageSize"`
		CurrentPage int       `json:"currentPage"`
		Pages       int       `json:"pages"`
		OrderBy     string    `json:"orderBy"`
		Results     []*Article `json:"results"`
	}
	Article struct {
		Id                 string        `json:"id"`
		Type               string        `json:"type"`
		SectionId          string        `json:"sectionId"`
		SectionName        string        `json:"sectionName"`
		WebPublicationDate string        `json:"webPublicationDate"`
		WebTitle           string        `json:"webTitle"`
		WebUrl             string        `json:"webUrl"`
		ApiUrl             string        `json:"apiUrl"`
		Blocks             ArticleBlocks `json:"blocks"`
		IsHosted           bool          `json:"isHosted"`
		PillarId           string        `json:"pillarId"`
		PillarName         string        `json:"pillarName"`
	}

	ArticleBlocks struct {
		Main            *BlockType   `json:"main"`
		Body            []*BlockType `json:"body"`
		TotalBodyBlocks int          `json:"totalBodyBlocks"`
	}

	BlockType struct {
		Id                 string            `json:"id"`
		BodyHtml           string            `json:"bodyHtml"`
		BodyTextSummary    string            `json:"bodyTextSummary"`
		Attributes         interface{}       `json:"attributes"`
		Published          bool              `json:"published"`
		CreatedDate        string            `json:"createdDate"`
		FirstPublishedDate string            `json:"firstPublishedDate"`
		PublishedDate      string            `json:"publishedDate"`
		LastModifedDate    string            `json:"lastModifiedDate"`
		Contributors       interface{}       `json:"contributors"`
		Elements           []ElementsElement `json:"elements"`
	}

	ElementsElement struct {
		Type         string      `json:"type"`
		Assets       interface{} `json:"assets"`
		TextTypeData interface{} `json:"textTypeData"`
	}

	ArticleResponseWrapper struct {
		Response ArticleResponse `json:"response"`
	}

	ArticleResponse struct {
		Status   string  `json:"status"`
		UserTier string  `json:"userTier"`
		Total    int     `json:"total"`
		Content  Article `json:"content"`
	}
)

func searchDefault() {
	response, err := http.Get("https://content.guardianapis.com/search?api-key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var responseWrapper SearchResponseWrapper
	if err := json.Unmarshal(data, &responseWrapper); err != nil {
		panic(err)
	}

	prettyJson, err := json.MarshalIndent(responseWrapper, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(prettyJson))
}

func getArticlesFromDate(startTime time.Time, pageSize int) (articles []Article, err error) {
	response, err := http.Get("https://content.guardianapis.com/search" +
		"?api-key=" + apiKey +
		"&order-by=oldest" +
		"&from-date=" + startTime.Format("2006-01-02T15:04:05Z") +
		"&page-size=" + strconv.Itoa(pageSize))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response %s from search: %v", response.Status, err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var responseWrapper SearchResponseWrapper
	if err := json.Unmarshal(data, &responseWrapper); err != nil {
		return nil, err
	}

	//prettyJson, err := json.MarshalIndent(responseWrapper, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(string(prettyJson))

	for _, v := range responseWrapper.Response.Results {
		//fmt.Printf("i=%d, id=%s\n", i, v.Id)
		articles = append(articles, *v)
	}

	return articles, nil
}

func getArticlesByDatePaginated(pageIndex, pageSize int) (articles []Article, err error) {
	url := "https://content.guardianapis.com/search?order-by=oldest&show-blocks=all" +
		"&page=" + strconv.Itoa(pageIndex) +
		"&page-size=" + strconv.Itoa(pageSize)
	fmt.Printf("%s\n", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("api-key", apiKey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response %s from GET: %v", response.Status, err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var responseWrapper SearchResponseWrapper
	if err := json.Unmarshal(data, &responseWrapper); err != nil {
		return nil, err
	}

	prettyJson, err := json.MarshalIndent(responseWrapper, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyJson))

	for _, v := range responseWrapper.Response.Results {
		//fmt.Printf("i=%d, id=%s\n", i, v.Id)
		articles = append(articles, *v)
	}

	return articles, nil
}

func getArticleById(id string) (*Article, error) {

	url := "https://content.guardianapis.com/" + id + "?show-blocks=all"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("api-key", apiKey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response %s from GET: %v", response.Status, err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var responseWrapper ArticleResponseWrapper
	if err := json.Unmarshal(data, &responseWrapper); err != nil {
		return nil, err
	}

	return &responseWrapper.Response.Content, nil
}

func main() {
	fmt.Println("Starting...")
	startTime := time.Now()

	// Test searching with default parameters
	//searchDefault()

	// Search for articles from a given date-time
	//articles, err := getArticlesFromDate(time.Date(1700, 12, 19, 11, 27, 14, 0, time.UTC), 3)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("main received %d articles\n", len(articles))

	// Retrieve a single article
	//article, err := getArticleById("news/1822/may/07/leadersandreply.mainsection")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Article summary: %s", article.Blocks.Body[0].BodyTextSummary)

	// Get several articles
	const pageSize int = 200
	const numPages int = 3
	//pageStartDate := time.Date(2017, 12, 19, 11, 27, 14, 0, time.UTC)
	for i := 1; i <= numPages; i++ {
		articles, err := getArticlesByDatePaginated(i, pageSize)
		if err != nil {
			panic(err)
		}
		fmt.Printf("main received %d articles\n", len(articles))
		for _, v := range articles {
			fmt.Printf("%s\n", v.Id)
		}
	}

	fmt.Printf("\nFinished in %dms\n", time.Now().Sub(startTime)/1000000)
}
