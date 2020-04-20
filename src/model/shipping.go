package model

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/constants"
)

type Shipping struct {
	/** Response Field */
	Id          *int64  `json:"_id"`
	Carrier     *string `json:"carrier"`
	Status      *string `json:"status"`
	OrderId     *int64  `json:"order_id"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	DeliveredAt *string `json:"delivered_at"`
	/** Database Field */
	RawId          *int64          `json:"-"`
	RawCarrier     *string         `json:"-"`
	RawStatus      *string         `json:"-"`
	RawOrderId     *int64          `json:"-"`
	RawCreatedAt   *string         `json:"-"`
	RawUpdatedAt   *string         `json:"-"`
	RawDeliveredAt *sql.NullString `json:"-"`
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
}

type BodyShipping struct {
	Id          *int64  `json:"_id" binding:"omitempty,gt=0"`
	Carrier     *string `json:"carrier" binding:"omitempty"`
	Status      *string `json:"status" binding:"omitempty"`
	OrderId     *int64  `json:"order_id" binding:"omitempty,gt=0"`
	DeliveredAt *string `json:"delivered_at" binding:"omitempty"`
}

func (body *BodyShipping) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryShipping struct {
	Id              *string `form:"id" binding:"omitempty"`
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
