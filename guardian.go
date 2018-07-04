package guardian

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
	"os"
	"log"
)

var apiKey string = os.Getenv("GUARDIAN_API_KEY")

type (
	SearchResponseWrapper struct {
		Response SearchResponse `json:"response"`
	}

	SearchResponse struct {
		Status      string             `json:"status"`
		UserTier    string             `json:"userTier"`
		Total       int                `json:"total"`
		StartIndex  int                `json:"startIndex"`
		PageSize    int                `json:"pageSize"`
		CurrentPage int                `json:"currentPage"`
		Pages       int                `json:"pages"`
		OrderBy     string             `json:"orderBy"`
		Results     []*GuardianArticle `json:"results"`
	}

	GuardianArticle struct {
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
		Id                 string         `json:"id"`
		BodyHtml           string         `json:"bodyHtml"`
		BodyTextSummary    string         `json:"bodyTextSummary"`
		Attributes         interface{}    `json:"attributes"`
		Published          bool           `json:"published"`
		CreatedDate        string         `json:"createdDate"`
		FirstPublishedDate string         `json:"firstPublishedDate"`
		PublishedDate      string         `json:"publishedDate"`
		LastModifedDate    string         `json:"lastModifiedDate"`
		Contributors       interface{}    `json:"contributors"`
		Elements           []*ElementType `json:"elements"`
	}

	ElementType struct {
		Type         string      `json:"type"`
		Assets       interface{} `json:"assets"`
		TextTypeData interface{} `json:"textTypeData"`
	}

	ArticleResponseWrapper struct {
		Response ArticleResponse `json:"response"`
	}

	ArticleResponse struct {
		Status   string          `json:"status"`
		UserTier string          `json:"userTier"`
		Total    int             `json:"total"`
		Content  GuardianArticle `json:"content"`
	}
)

func (ga *GuardianArticle) IdString() string {
	return ga.Id
}

func (ga *GuardianArticle) Title() string {
	return ga.WebTitle
}

func (ga *GuardianArticle) ArticleDate() time.Time {
	return time.Now() //TODO
}

func (ga *GuardianArticle) Body() string {
	return ga.Blocks.Body[0].BodyTextSummary
}

func (ga *GuardianArticle) Json() string {
	var ba []byte
	var err error
	if ba, err = json.Marshal(ga); err != nil {
		panic(err)
	}
	return string(ba[:])
}

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

	//prettyJson, err := json.MarshalIndent(responseWrapper, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(string(prettyJson))
}

func getArticlesFromDate(startTime time.Time, pageSize int) (articles []GuardianArticle, err error) {
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
		//fmt.Printf("i=%d, id=%s\n", i, v.Id_)
		articles = append(articles, *v)
	}

	return articles, nil
}

// GetArticlesByDatePaginated retrieves a page of articles. Repeated calls with
// increasing values for pageIndex can retrieve larger sets of articles.
// PageIndex starts at 1. Maximum page size appears to be 200.
func GetArticlesByDatePaginated(pageIndex, pageSize int, startTime time.Time) (articles []GuardianArticle, err error) {
	url := "https://content.guardianapis.com/search?" +
		"order-by=oldest" +
		"&show-blocks=all" +
		"&page=" + strconv.Itoa(pageIndex) +
		"&page-size=" + strconv.Itoa(pageSize) +
		"&from-date=" + startTime.Format("2006-01-02T15:04:05Z")
	log.Printf("%s\n", url)
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

	//prettyJson, err := json.MarshalIndent(responseWrapper, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(prettyJson))

	for _, v := range responseWrapper.Response.Results {
		//fmt.Printf("i=%d, id=%s\n", i, v.Id_)
		articles = append(articles, *v)
	}

	return articles, nil
}

func getArticleById(id string) (*GuardianArticle, error) {

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
