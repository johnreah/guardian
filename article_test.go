package guardian

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestGuardianArticleCast(t *testing.T) {
	ga := &GuardianArticle{
		Id: "testId",
		WebTitle: "testTitle",
		Blocks: ArticleBlocks{
			Body: []*BlockType{
				{
					BodyTextSummary: "testBody",
				},
			},
		},
	}
	a := Article(ga)
	assert.Equal(t, a.ArticleId(), "testId")
	assert.Equal(t, a.Title(), "testTitle")
	assert.Equal(t, a.Body(), "testBody")
}

func TestTestArticleCast(t *testing.T) {
	ta := makeTestArticle("testIdString", time.Now(), "testTitle", "testBody", "testJson")
	a := Article(ta)
	assert.Equal(t, a.ArticleId(), "testIdString")
	assert.Equal(t, a.Title(), "testTitle")
	assert.Equal(t, a.Body(), "testBody")
	assert.Equal(t, a.Json(), "testJson")
}

func makeTestArticle(articleId string, articleDate time.Time, title, body, json string) *TestArticle {
	return &TestArticle{
		articleId: articleId,
		articleDate: articleDate,
		title: title,
		body: body,
		json: json,
	}
}

/*
type TestArticle struct {
	idString string
	title string
	articleDate time.Time
	body string
	json string
}

func (ta *TestArticle) ArticleId() string { return ta.idString }
func (ta *TestArticle) Title() string { return ta.title }
func (ta *TestArticle) ArticleDate() time.Time { return ta.articleDate }
func (ta *TestArticle) Body() string { return ta.body }
func (ta *TestArticle) Json() string { return ta.json }
*/