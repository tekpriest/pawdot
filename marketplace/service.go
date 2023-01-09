package marketplace

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pawdot.app/models"
	"pawdot.app/utils"
)

type Service interface {
	CreateSale(userID string, data ICreateSale) (*models.Sale, error)
	PlaceBid(saleID, userId string, amount float32) (*models.Sale, error)
	FetchBid(userID, saleID string) (*models.Bid, error)
	FetchAllSales(query IQuerySales) (GetAllSales, error)
	FetchSale(saleID string) (*models.Sale, error)
	FetchPersonalSales(userID string, query IQuerySales) (GetAllSales, error)
	SetWinner(saleID, bidderID string) (*models.Sale, error)
	FetchSaleBids(saleID string) ([]models.Bid, error)
	FetchPersonalBids(userID string) ([]models.Bid, error)
}

type service struct {
	db *gorm.DB
}

func InitService(db utils.DatabaseConnection) Service {
	return &service{
		db: db.GetDB(),
	}
}

func queryByBreed(breed ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(breed) > 0 && breed[0] != "" {
			return db.Where("breed = ?", breed[0])
		}
		return db.Where("breed IS NOT NULL")
	}
}

func queryByCategory(category ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(category) > 0 && category[0] != "" {
			return db.Where("category = ?", category[0])
		}
		return db.Where("category IS NOT NULL")
	}
}

// CreateSale implements Service
func (s *service) CreateSale(userID string, data ICreateSale) (*models.Sale, error) {
	newSale := &models.Sale{
		Category:    data.Category,
		Breed:       data.Breed,
		Title:       data.Title,
		Description: data.Description,
		StartingBid: data.StartingBid,
		TraderID:    userID,
	}
	if err := s.db.Create(newSale).Error; err != nil {
		return nil, err
	}

	return s.FetchSale(newSale.ID)
}

// FetchAllSales implements Service
func (s *service) FetchAllSales(query IQuerySales) (GetAllSales, error) {
	var sales []models.Sale
	var count int64

	if err := s.db.
		Table("sales").
		Preload(clause.Associations).
		Scopes(
			queryByBreed(query.Breed),
			queryByCategory(query.Category),
		).
		Order("created_at DESC").
		Limit(query.Limit).
		Offset((query.Page - 1) * query.Limit).
		Find(&sales).Error; err != nil {
		return GetAllSales{}, err
	}

	return GetAllSales{
		Sales: sales,
		Meta:  utils.Paginate(count, len(sales), query.Page, query.Limit),
	}, nil
}

// FetchPersonalSales implements Service
func (s *service) FetchPersonalSales(
	userID string, query IQuerySales,
) (GetAllSales, error) {
	var sales []models.Sale
	var count int64

	if err := s.db.
		Table("sales").
		Preload(clause.Associations).
		Scopes(
			queryByBreed(query.Breed),
			queryByCategory(query.Category),
		).
		Order("created_at DESC").
		Limit(query.Limit).
		Offset((query.Page-1)*query.Limit).
		Where("trader_id = ?", userID).
		Find(&sales).Error; err != nil {
		return GetAllSales{}, err
	}
	return GetAllSales{
		Sales: sales,
		Meta: utils.
			Paginate(count, len(sales), query.Page, query.Limit),
	}, nil
}

// PlaceBid implements Service
func (s *service) PlaceBid(saleID, userID string, amount float32) (*models.Sale, error) {
	sale, err := s.FetchSale(saleID)
	if err != nil {
		return nil, err
	}

	bid, _ := s.FetchBid(userID, sale.ID)

	if bid == nil {
		bid = &models.Bid{
			UserID: userID,
			Amount: sale.StartingBid + amount,
			SaleID: sale.ID,
		}
		if err := s.db.Create(bid).Error; err != nil {
			return nil, err
		}
	} else {
		if err := s.db.Table("bids").
			Where("id = ?", bid.ID).
			Update("amount", bid.Amount+amount).
			Error; err != nil {
			return nil, err
		}
	}

	return s.FetchSale(saleID)
}

// SetWinner implements Service
func (s *service) SetWinner(saleID, bidderID string) (*models.Sale, error) {
	sale, err := s.FetchSale(saleID)
	if err != nil {
		return nil, err
	}

	if err := s.db.
		Where("id = ?", sale.ID).
		Update("sold_to", bidderID).Error; err != nil {
		return nil, err
	}

	return s.FetchSale(sale.ID)
}

// FetchSale implements Service
func (s *service) FetchSale(saleID string) (*models.Sale, error) {
	var sale models.Sale
	var bids []models.Bid

	if err := s.db.Where("sale_id =?", saleID).
		Select([]string{"amount", "created_at"}).
		Find(&bids).Error; err != nil {
		return nil, err
	}

	if err := s.db.
		Where("id = ?", saleID).
		Preload(clause.Associations).
		First(&sale).Error; err != nil {
		return nil, err
	}

	sale.Bids = bids

	return &sale, nil
}

// FetchSaleBids implements Service
func (s *service) FetchSaleBids(saleID string) ([]models.Bid, error) {
	var bids []models.Bid

	if err := s.db.
		Where("sale_id = ?", saleID).
		Find(&bids).Error; err != nil {
		return bids, err
	}

	return bids, nil
}

// FetchBid implements Service
func (s *service) FetchBid(userID, saleID string) (*models.Bid, error) {
	var bid models.Bid

	if err := s.db.
		First(&bid, "user_id = ? AND sale_id = ?", userID, saleID).Error; err != nil {
		return nil, err
	}

	return &bid, nil
}

// FetchPersonalBids implements Service
func (s *service) FetchPersonalBids(userID string) ([]models.Bid, error) {
	var bids []models.Bid

	if err := s.db.Where("user_id = ?", userID).Find(&bids).Error; err != nil {
		return bids, err
	}

	return bids, nil
}
