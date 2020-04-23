package model

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/constants"
)

type Product struct {
	/** Response Field */
	Id          *int64   `json:"_id,omitempty"`
	CategoryId  *int64   `json:"category_id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	OutPrice    *float64 `json:"out_price,omitempty"`
	Discount    *float64 `json:"discount,omitempty"`
	ImagePath   *string  `json:"image_path,omitempty"`
	Description *string  `json:"description,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
	/** Database Field */
	RawId          *int64          `json:"-"`
	RawCategoryId  *int64          `json:"-"`
	RawName        *string         `json:"-"`
	RawOutPrice    *float64        `json:"-"`
	RawDiscount    *float64        `json:"-"`
	RawImagePath   *sql.NullString `json:"-"`
	RawDescription *sql.NullString `json:"-"`
	RawCreatedAt   *string         `json:"-"`
	RawUpdatedAt   *string         `json:"-"`
}

func (product *Product) FillResponse() {
	product.Id = product.RawId
	product.CategoryId = product.RawCategoryId
	product.Name = product.RawName
	product.OutPrice = product.RawOutPrice
	product.Discount = product.RawDiscount
	if product.RawImagePath != nil {
		if product.RawImagePath.Valid {
			product.ImagePath = &product.RawImagePath.String
		}
	}
	if product.RawDescription != nil {
		if product.RawDescription.Valid {
			product.Description = &product.RawDescription.String
		}
	}
	product.CreatedAt = product.RawCreatedAt
	product.UpdatedAt = product.RawUpdatedAt
}

type BodyProduct struct {
	Id          *int64   `json:"_id" binding:"omitempty,gt=0"`
	CategoryId  *int64   `json:"category_id" binding:"omitempty,gt=0"`
	Name        *string  `json:"name" binding:"omitempty"`
	OutPrice    *float64 `json:"out_price" binding:"omitempty,gt=0"`
	Discount    *float64 `json:"discount" binding:"omitempty,gte=0"`
	ImagePath   *string  `json:"image_path" binding:"omitempty"`
	Description *string  `json:"description" binding:"omitempty"`
}

func (body *BodyProduct) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryProduct struct {
	Id               *int64   `form:"id" binding:"omitempty"`
	CategoryId       *int64   `form:"category_id" binding:"omitempty"`
	CategoryParentId *int64   `form:"category_parent_id" binding:"omitempty"`
	Name             *string  `form:"name" binding:"omitempty"`
	OutPrice         *float64 `form:"out_price" binding:"omitempty"`
	Discount         *float64 `form:"discount" binding:"omitempty"`
	ImagePath        *string  `form:"image_path" binding:"omitempty"`
	Description      *string  `form:"description" binding:"omitempty"`
	CreatedAtFrom    *string  `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo      *string  `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom    *string  `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo      *string  `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	Page             *int     `form:"page" binding:"omitempty,gte=0"`
	Limit            *int     `form:"limit" binding:"omitempty,gte=0"`
	Offset           *int
}

func (query *QueryProduct) ParsePaging() {
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

type SliceBodyProduct struct {
	Data []*BodyProduct `json:"data"`
}
