// +build integration

package guardian

import (
	"fmt"
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func testSearchDefault(t *testing.T) {
	// Test searching with default parameters
	searchDefault()
}

func testGetArticlesFromDate(t *testing.T) {
	const pageSize = 3
	// Search for articles from a given date-time
	articles, err := getArticlesFromDate(time.Date(1700, 12, 19, 11, 27, 14, 0, time.UTC), pageSize)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, pageSize, len(articles), "Unexpected number of articles in page")
}

// Get a single article
func testGetSingleArticle(t *testing.T) {
	article, err := getArticleById("news/1822/may/07/leadersandreply.mainsection")
	if err != nil {
		t.Errorf("getArticleById error: %v", err)
	}
	assert.Contains(t, article.Blocks.Body[0].BodyTextSummary, "Thackeray")

	//fmt.Printf("GuardianArticle summary: %s", article.Blocks.Body[0].BodyTextSummary)
	//prettyJson, err := json.MarshalIndent(article, "", "  ")
	//if err != nil {
	//	t.Errorf("TestGetSingleArticle marshalling error: %v", err)
	//}
	//fmt.Println(string(prettyJson))
}

// Get pages of articles ordered by date
func testGetArticlesByDatePaginated(t *testing.T) {
	const pageSize int = 5
	const numPages int = 3
	pagesRetrieved := 0
	fromTime := time.Date(2000, 12, 19, 0, 0, 0, 0, time.UTC)
	for i := 1; i <= numPages; i++ {
		articles, err := GetArticlesByDatePaginated(i, pageSize, fromTime)
		if err != nil {
			panic(err)
		}
		log.Printf("main received %d articles\n", len(articles))
		for _, v := range articles {
			log.Printf("%s\n", v.Id)
			pagesRetrieved++
		}
	}
	assert.Equal(t, pageSize * numPages, pagesRetrieved)
}

func TestScraping(t *testing.T) {
	fmt.Println("Starting...")
	startTime := time.Now()

	const pageSize int = 200
	const numPages int = 2
	fromTime := time.Date(1700, 1, 1, 0, 0, 0, 0, time.UTC)
	pagesRetrieved := 0
	for i := 1; i <= numPages; i++ {
		articles, err := GetArticlesByDatePaginated(i, pageSize, fromTime)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got %d articles\n", len(articles))
		for _, v := range articles {
			fmt.Printf("%s %s\n", v.WebPublicationDate, v.Id)
			pagesRetrieved++
		}
	}
	fmt.Printf("\nFinished in %dms\n", time.Now().Sub(startTime)/1000000)
}
