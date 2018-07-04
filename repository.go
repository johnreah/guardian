package guardian

type Repository interface {
	Get(id string) (*Article, error)
	Put(article *Article) error
	Count() int
	GetMostRecent() (*Article, error)
}
