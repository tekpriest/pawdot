package utils

import (
	"log"
	"os"
)

type Utils interface {
	AppError(err error, msg ...string)
}

type Pagination struct {
	Size         int   `json:"size"`
	TotalItems   int64 `json:"totalItems"`
	NextPage     int   `json:"nextPage"`
	PreviousPage int   `json:"previousPage"`
} //	@Name	Pagination

func AppError(err error, msg ...string) {
	log.Fatalf("%s %s", err, msg)
	os.Exit(1)
}

// func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		switch {
// 		case limit > 100:
// 			limit = 100
// 		case limit <= 0:
// 			limit = 20
// 		}
// 		offset := (page - 1) * limit
// 		return db.Offset(offset).Limit(limit)
// 	}
// }

func Paginate(count int64, totalData, page, size int) Pagination {
	nextPage := page + 1
	pages := float64(count / int64(size))

	if nextPage > int(pages) {
		nextPage = int(pages)
	}

	if page-1 == 0 {
		page = 1
	}

	return Pagination{
		Size:         totalData,
		TotalItems:   count,
		NextPage:     nextPage,
		PreviousPage: page,
	}
}
