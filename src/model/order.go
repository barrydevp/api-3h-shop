package model

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/constants"
)

type Order struct {
	/** Response Field */
	Id                *int64  `json:"_id,omitempty"`
	Session           *string `json:"session,omitempty"`
	CustomerId        *int64  `json:"customer_id,omitempty"`
	PaymentStatus     *string `json:"payment_status,omitempty"`
	FulfillmentStatus *string `json:"fulfillment_status,omitempty"`
	Note              *string `json:"note,omitempty"`
	CreatedAt         *string `json:"created_at,omitempty"`
	UpdatedAt         *string `json:"updated_at,omitempty"`
	PaidAt            *string `json:"paid_at,omitempty"`
	FulfilledAt       *string `json:"fulfill_at,omitempty"`
	CancelledAt       *string `json:"cancelled_at,omitempty"`
	/** Database Field */
	RawId                *int64          `json:"-"`
	RawSession           *string         `json:"-"`
	RawCustomerId        *sql.NullInt64  `json:"-"`
	RawPaymentStatus     *string         `json:"-"`
	RawFulfillmentStatus *string         `json:"-"`
	RawNote              *sql.NullString `json:"-"`
	RawCreatedAt         *string         `json:"-"`
	RawUpdatedAt         *string         `json:"-"`
	RawPaidAt            *sql.NullString `json:"-"`
	RawFulfilledAt       *sql.NullString `json:"-"`
	RawCancelledAt       *sql.NullString `json:"-"`
}

func (order *Order) FillResponse() {
	order.Id = order.RawId
	order.Session = order.RawSession
	if order.RawCustomerId != nil {
		if order.RawCustomerId.Valid {
			order.CustomerId = &order.RawCustomerId.Int64
		}
	}
	order.PaymentStatus = order.RawPaymentStatus
	order.FulfillmentStatus = order.RawFulfillmentStatus
	if order.RawNote != nil {
		if order.RawNote.Valid {
			order.Note = &order.RawNote.String
		}
	}
	order.CreatedAt = order.RawCreatedAt
	order.UpdatedAt = order.RawUpdatedAt
	if order.RawFulfilledAt != nil {
		if order.RawFulfilledAt.Valid {
			order.FulfilledAt = &order.RawFulfilledAt.String
		}
	}
	if order.RawPaidAt != nil {
		if order.RawPaidAt.Valid {
			order.PaidAt = &order.RawPaidAt.String
		}
	}
	if order.RawCancelledAt != nil {
		if order.RawCancelledAt.Valid {
			order.CancelledAt = &order.RawCancelledAt.String
		}
	}
}

type BodyOrder struct {
	Id                *int64  `json:"_id" binding:"omitempty,gt=0"`
	Session           *string `json:"session" binding:"omitempty"`
	CustomerId        *int64  `json:"customer_id" binding:"omitempty,gt=0"`
	PaymentStatus     *string `json:"payment_status" binding:"omitempty"`
	FulfillmentStatus *string `json:"fulfillment_status" binding:"omitempty"`
	Note              *string `json:"note" binding:"omitempty"`
	PaidAt            *string `json:"paid_at" binding:"omitempty"`
	FulfilledAt       *string `json:"fulfill_at" binding:"omitempty"`
	CancelledAt       *string `json:"cancelled_at" binding:"omitempty"`
}

func (body *BodyOrder) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryOrder struct {
	Id                *string `form:"id" binding:"omitempty"`
	Session           *string `form:"session" binding:"omitempty"`
	CustomerId        *int64  `form:"customer_id" binding:"omitempty"`
	PaymentStatus     *string `form:"payment_status" binding:"omitempty"`
	FulfillmentStatus *string `form:"fulfillment_status" binding:"omitempty"`
	Note              *string `form:"note" binding:"omitempty"`
	CreatedAtFrom     *string `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo       *string `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom     *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo       *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	PaidAtFrom        *string `form:"paid_at_from" binding:"omitempty,required_with=PaidAtTo,datetime"`
	PaidAtTo          *string `form:"paid_at_to" binding:"omitempty,required_with=PaidAtFrom,datetime"`
	FulfilledAtFrom   *string `form:"fulfilled_at_from" binding:"omitempty,required_with=FulfilledAtTo,datetime"`
	FulfilledAtTo     *string `form:"fulfilled_at_to" binding:"omitempty,required_with=FulfilledAtFrom,datetime"`
	CancelledAtFrom   *string `form:"cancelled_at_from" binding:"omitempty,required_with=CancelledAtTo,datetime"`
	CancelledAtTo     *string `form:"cancelled_at_to" binding:"omitempty,required_with=CancelledAtFrom,datetime"`
	Page              *int    `form:"page" binding:"omitempty,gte=0"`
	Limit             *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset            *int
}

func (query *QueryOrder) ParsePaging() {
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
