package marketplace

type ICreateSale struct {
	Category    string  `json:"category"    validate:"required"`
	Breed       string  `json:"breed"       validate:"required"`
	Title       string  `json:"title"       validate:"required"`
	Description string  `json:"description" validate:"required"`
	StartingBid float32 `json:"startingBid" validate:"required"`
}

type IPaginateQuery struct {
	Page  int
	Limit int
}

type IQuerySales struct {
	Breed    string `json:"breed"`
	Category string `json:"category"`
	IPaginateQuery
}

type ICreateBid struct {
	Amount float32 `json:"amount" validate:"required"`
}
