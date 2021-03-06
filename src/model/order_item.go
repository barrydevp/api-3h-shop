package model

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/constants"
)

type OrderItem struct {
	/** Response Field */
	Id            *int64  `json:"_id,omitempty"`
	ProductId     *int64  `json:"product_id,omitempty"`
	ProductItemId *int64  `json:"product_item_id,omitempty"`
	OrderId       *int64  `json:"order_id,omitempty"`
	Quantity      *int64  `json:"quantity,omitempty"`
	Status        *string `json:"status,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
	UpdatedAt     *string `json:"updated_at,omitempty"`
	WarrantyId    *int64  `json:"warranty_id,omitempty"`
	/** Database Field */
	RawId            *int64         `json:"-"`
	RawProductId     *int64         `json:"-"`
	RawProductItemId *sql.NullInt64 `json:"-"`
	RawOrderId       *int64         `json:"-"`
	RawQuantity      *int64         `json:"-"`
	RawStatus        *string        `json:"-"`
	RawCreatedAt     *string        `json:"-"`
	RawUpdatedAt     *string        `json:"-"`
	RawWarrantyId    *sql.NullInt64 `json:"-"`
}

func (orderItem *OrderItem) FillResponse() {
	orderItem.Id = orderItem.RawId
	orderItem.ProductId = orderItem.RawProductId
	if orderItem.RawProductItemId != nil {
		if orderItem.RawProductItemId.Valid {
			orderItem.ProductItemId = &orderItem.RawProductItemId.Int64
		}
	}
	if orderItem.RawWarrantyId != nil {
		if orderItem.RawWarrantyId.Valid {
			orderItem.WarrantyId = &orderItem.RawWarrantyId.Int64
		}
	}
	orderItem.OrderId = orderItem.RawOrderId
	orderItem.Status = orderItem.RawStatus
	orderItem.Quantity = orderItem.RawQuantity
	orderItem.CreatedAt = orderItem.RawCreatedAt
	orderItem.UpdatedAt = orderItem.RawUpdatedAt
}

type BodyOrderItem struct {
	Id            *int64  `json:"_id" binding:"omitempty,gt=0"`
	ProductId     *int64  `json:"product_id" binding:"omitempty,gt=0"`
	ProductItemId *int64  `json:"product_item_id" binding:"omitempty,gt=0"`
	OrderId       *int64  `json:"order_id" binding:"omitempty,gt=0"`
	Quantity      *int64  `json:"quantity" binding:"omitempty,gte=0"`
	Status        *string `json:"status"`
	WarrantyId    *int64  `json:"warranty_id" binding:"omitempty,gt=0"`
}

func (body *BodyOrderItem) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryOrderItem struct {
	Id            *int64  `form:"id" binding:"omitempty"`
	ProductId     *int64  `form:"product_id" binding:"omitempty"`
	ProductItemId *int64  `form:"product_item_id" binding:"omitempty"`
	OrderId       *int64  `form:"order_id" binding:"omitempty"`
	Quantity      *int64  `form:"quantity" binding:"omitempty"`
	WarrantyId    *int64  `form:"warranty_id" binding:"omitempty"`
	Status        *string `form:"status"`
	CreatedAtFrom *string `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo   *string `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo   *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	Page          *int    `form:"page" binding:"omitempty,gte=0"`
	Limit         *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryOrderItem) ParsePaging() {
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

type OrderItemJoinProduct struct {
	*OrderItem `json:"order_item"`
	*Product   `json:"product"`
	*Warranty  `json:"warranty"`
}

func (orderItemJoinProduct *OrderItemJoinProduct) FillResponse() {
	if orderItemJoinProduct.OrderItem != nil {
		orderItemJoinProduct.OrderItem.FillResponse()
	}
	if orderItemJoinProduct.Product != nil {
		orderItemJoinProduct.Product.FillResponse()
	}
	if orderItemJoinProduct.Warranty != nil {
		orderItemJoinProduct.Warranty.FillResponse()
	}
}
