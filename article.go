package guardian

import "time"

// Article provides an interface that exposes methods common to articles from different sources
// that may have different implementations, such as the Guardian or the Telegraph.
type Article interface {
	ArticleId() string
	ArticleDate() time.Time
	Title() string
	Body() string
	Json() string
}

type TestArticle struct {
	articleId string
	articleDate time.Time
	title string
	body string
	json string
}

func (ta *TestArticle) ArticleId() string { return ta.articleId }
func (ta *TestArticle) ArticleDate() time.Time { return ta.articleDate }
func (ta *TestArticle) Title() string { return ta.title }
func (ta *TestArticle) Body() string { return ta.body }
func (ta *TestArticle) Json() string { return ta.json }
