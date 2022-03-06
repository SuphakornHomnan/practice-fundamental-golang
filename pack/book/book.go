package book

type book struct {
	name string
}

func New() book {
	return book{name: "Suphakorn"}
}
