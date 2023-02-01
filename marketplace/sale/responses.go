package sale

import (
	"pawdot.app/models"
)

type CreateBidSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"bid placed successfully"`
	Data    models.Sale
} //	@Name	CreateBidSuccessfulResponse

type CreateSaleSuccessfulResponse struct {
	Success bool        `example:"true"`
	Mesage  string      `example:"sale created successfully"`
	Data    models.Sale `example:"id:clcv8zrk90000vboy7d4tb5ms,createdAt:2023-01-14T02:06:31+01:00,updatedAt:2023-01-14T02:06:31+01:00,type:NEW,category:CAT,priority:1,breed:Cat Dog,title:Bull Dog For Sale,description:BDG for sale,startingBig:20000,traderId:clcnz7h460000s1mo30jmkewi,status:PENDING,sold:false,bidCount:0" swaggertype:"object,string"`
} //	@Name	CreateSaleSuccessfulResponse

type GetAllSalesSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched all current sales"`
	Data    GetAllSales
} // @Name GetAllSalesSuccessfulResponse

type GetSaleSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched all current sales"`
	Data    models.Sale
} // @Name GetSaleSuccessfulResponse

type GetSaleBidsSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched all sale bids"`
	Data    []models.Bid
} // @Name GetSaleBidsSuccessfulResponse

type GetPersonalBidsSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched all sale bids"`
	Data    []models.Bid
} // @Name GetPersonalBidsSuccessfulResponse

type GetPersonalSalesSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched all sale bids"`
	Data    GetAllSales
} // @Name GetPersonalSalesSuccessfulResponse

type CancelSaleSuccessfulResposnse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"sale cancelled successfully"`
} // @Name CancelSaleSuccessfulResposnse

type RepublishSaleSuccessfulResponse struct {
	Success bool        `example:"true"`
	Mesage  string      `example:"sale republished successfully"`
	Data    models.Sale `example:"id:clcv8zrk90000vboy7d4tb5ms,createdAt:2023-01-14T02:06:31+01:00,updatedAt:2023-01-14T02:06:31+01:00,type:NEW,category:CAT,priority:1,breed:Cat Dog,title:Bull Dog For Sale,description:BDG for sale,startingBig:20000,traderId:clcnz7h460000s1mo30jmkewi,status:PENDING,sold:false,bidCount:0" swaggertype:"object,string"`
} //	@Name	RepublishSaleSuccessfulResponse
