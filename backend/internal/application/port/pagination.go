package port

type PaginationOutput interface {
	GetTotalItems(table string) int64
	GetTableName(structTable interface{}) string
}

type PaginationInput interface {
	GetTotalItems(object interface{}) int64
}
