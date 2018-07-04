package guardian

import "time"

// Article provides an interface that exposes methods common to articles from different sources
// that may have different implementations, such as the Guardian or the Telegraph.
type Article interface {
	IdString() string
	Title() string
	ArticleDate() time.Time
	Body() string
	Json() string
}
