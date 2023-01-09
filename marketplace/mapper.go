package marketplace

import (
	"pawdot.app/models"
	"pawdot.app/utils"
)

type GetAllSales struct {
	Sales []models.Sale    `json:"sales"`
	Meta  utils.Pagination `json:"meta"`
}
