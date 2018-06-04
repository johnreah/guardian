package guardian

import (
	"fmt"
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

func testSearchDefault(t *testing.T) {
	// Test searching with default parameters
	searchDefault()
}

func testGetArticlesFromDate(t *testing.T) {
	// Search for articles from a given date-time
	articles, err := getArticlesFromDate(time.Date(1700, 12, 19, 11, 27, 14, 0, time.UTC), 3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("main received %d articles\n", len(articles))
}

func testGetSingleArticle(t *testing.T) {
	// Retrieve a single article
	article, err := getArticleById("news/1822/may/07/leadersandreply.mainsection")
	if err != nil {
		t.Errorf("getArticleById error: %v", err)
	}
	assert.Contains(t, article.Blocks.Body[0].BodyTextSummary, "Thackeray")
	//fmt.Printf("Article summary: %s", article.Blocks.Body[0].BodyTextSummary)
	//prettyJson, err := json.MarshalIndent(article, "", "  ")
	//if err != nil {
	//	t.Errorf("TestGetSingleArticle marshalling error: %v", err)
	//}
	//fmt.Println(string(prettyJson))

}

func TestGuardian(t *testing.T) {
	fmt.Println("Starting...")
	startTime := time.Now()

	// Get several articles
	const pageSize int = 200
	const numPages int = 3
	fromTime := time.Date(2017, 12, 19, 11, 27, 14, 0, time.UTC)
	for i := 1; i <= numPages; i++ {
		articles, err := getArticlesByDatePaginated(i, pageSize, fromTime)
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
