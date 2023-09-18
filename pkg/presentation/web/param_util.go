package web

const (
	defaultOffset int = 0
	defaultLimit  int = 15
)

type Pagination struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func NewPagination() Pagination {
	return Pagination{
		Offset: defaultOffset,
		Limit:  defaultLimit,
	}
}
