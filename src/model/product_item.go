package model

import "github.com/barrydev/api-3h-shop/src/constants"

type ProductItem struct {
	/** Response Field */
	Id        *int64   `json:"_id,omitempty"`
	ProductId *int64   `json:"product_id,omitempty"`
	Stock     *int64   `json:"stock,omitempty"`
	InPrice   *float64 `json:"in_price,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	UpdatedAt *string  `json:"updated_at,omitempty"`
	ExpiredAt *string  `json:"expired_at,omitempty"`
	/** Database Field */
	RawId        *int64   `json:"-"`
	RawProductId *int64   `json:"-"`
	RawStock     *int64   `json:"-"`
	RawInPrice   *float64 `json:"-"`
	RawCreatedAt *string  `json:"-"`
	RawUpdatedAt *string  `json:"-"`
	RawExpiredAt *string  `json:"-"`
}

func (productItem *ProductItem) FillResponse() {
	productItem.Id = productItem.RawId
	productItem.ProductId = productItem.RawProductId
	productItem.Stock = productItem.RawStock
	productItem.InPrice = productItem.RawInPrice
	productItem.CreatedAt = productItem.RawCreatedAt
	productItem.UpdatedAt = productItem.RawUpdatedAt
	productItem.ExpiredAt = productItem.RawExpiredAt
}

type BodyProductItem struct {
	Id        *int64   `json:"_id" binding:"omitempty,gt=0"`
	ProductId *int64   `json:"product_id" binding:"omitempty,gt=0"`
	Stock     *int64   `json:"stock" binding:"omitempty,gte=0"`
	InPrice   *float64 `json:"in_price" binding:"omitempty,gt=0"`
	ExpiredAt *string  `json:"expired_at" binding:"omitempty"`
}

func (body *BodyProductItem) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryProductItem struct {
	Id            *string  `form:"id" binding:"omitempty"`
	ProductId     *int64   `json:"product_id" binding:"omitempty"`
	Stock         *int64   `json:"stock" binding:"omitempty"`
	InPrice       *float64 `json:"in_price" binding:"omitempty"`
	CreatedAtFrom *string  `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo   *string  `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom *string  `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo   *string  `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	ExpiredAtFrom *string  `form:"expired_at_from" binding:"omitempty,required_with=ExpiredAtTo,datetime"`
	ExpiredAtTo   *string  `form:"expired_at_to" binding:"omitempty,required_with=ExpiredAtFrom,datetime"`
	Page          *int     `form:"page" binding:"omitempty,gte=0"`
	Limit         *int     `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryProductItem) ParsePaging() {
	if query.Page == nil {
		page := constants.DEFAULT_PAGE
		query.Page = &page
	}

	if query.Limit == nil {
		limit := constants.DEFAULT_LIMIT
		query.Limit = &limit
	}

	skip := (*query.Page - 1) * *query.Limit

	query.Offset = &skip
}
