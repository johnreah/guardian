package guardian

type Article interface {
	Id() string
	Title() string
	Body() string
	Source() string
}
