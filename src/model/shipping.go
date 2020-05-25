package model

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/constants"
)

type Shipping struct {
	/** Response Field */
	Id          *int64   `json:"_id,omitempty"`
	Carrier     *string  `json:"carrier,omitempty"`
	Status      *string  `json:"status,omitempty"`
	OrderId     *int64   `json:"order_id,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
	DeliveredAt *string  `json:"delivered_at,omitempty"`
	Note        *string  `json:"note,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	/** Database Field */
	RawId          *int64          `json:"-"`
	RawCarrier     *string         `json:"-"`
	RawStatus      *string         `json:"-"`
	RawOrderId     *int64          `json:"-"`
	RawCreatedAt   *string         `json:"-"`
	RawUpdatedAt   *string         `json:"-"`
	RawDeliveredAt *sql.NullString `json:"-"`
	RawNote        *sql.NullString `json:"-"`
	RawPrice       float64         `json:"-"`
}

func (shipping *Shipping) FillResponse() {
	shipping.Id = shipping.RawId
	shipping.Carrier = shipping.RawCarrier
	shipping.Status = shipping.RawStatus
	shipping.OrderId = shipping.RawOrderId
	shipping.CreatedAt = shipping.RawCreatedAt
	shipping.UpdatedAt = shipping.RawUpdatedAt
	if shipping.RawDeliveredAt != nil {
		if shipping.RawDeliveredAt.Valid {
			shipping.DeliveredAt = &shipping.RawDeliveredAt.String
		}
	}
	if shipping.RawNote != nil {
		if shipping.RawNote.Valid {
			shipping.Note = &shipping.RawNote.String
		}
	}
}

type BodyShipping struct {
	Id          *int64   `json:"_id" binding:"omitempty,gt=0"`
	Carrier     *string  `json:"carrier" binding:"omitempty"`
	Status      *string  `json:"status" binding:"omitempty"`
	OrderId     *int64   `json:"order_id" binding:"omitempty,gt=0"`
	DeliveredAt *string  `json:"delivered_at" binding:"omitempty"`
	Note        *string  `json:"note" binding:"omitempty"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
}

func (body *BodyShipping) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryShipping struct {
	Id              *int64  `form:"id" binding:"omitempty"`
	Carrier         *string `json:"carrier" binding:"omitempty"`
	Status          *string `json:"status" binding:"omitempty"`
	OrderId         *int64  `json:"order_id" binding:"omitempty"`
	CreatedAtFrom   *string `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo     *string `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom   *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo     *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	DeliveredAtFrom *string `form:"delivered_at_from" binding:"omitempty,required_with=DeliveredTo,datetime"`
	DeliveredAtTo   *string `form:"delivered_at_to" binding:"omitempty,required_with=DeliveredFrom,datetime"`
	Page            *int    `form:"page" binding:"omitempty,gte=0"`
	Limit           *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset          *int
}

func (query *QueryShipping) ParsePaging() {
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
