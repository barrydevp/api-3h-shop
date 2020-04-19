package model

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/constants"
)

type Order struct {
	/** Response Field */
	Id                *int64  `json:"_id"`
	Session           *string `json:"session"`
	CustomerId        *int64  `json:"customer_id"`
	PaymentStatus     *string `json:"payment_status"`
	FulfillmentStatus *string `json:"fulfillment_status"`
	Note              *string `json:"note"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
	PaidAt            *string `json:"paid_at"`
	FulfilledAt       *string `json:"fulfill_at"`
	CancelledAt       *string `json:"cancelled_at"`
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

func (cus *Order) FillResponse() {
	cus.Id = cus.RawId
	cus.Session = cus.RawSession
	if cus.RawCustomerId != nil {
		if cus.RawCustomerId.Valid {
			cus.CustomerId = &cus.RawCustomerId.Int64
		}
	}
	cus.PaymentStatus = cus.RawPaymentStatus
	cus.FulfillmentStatus = cus.RawFulfillmentStatus
	if cus.RawNote != nil {
		if cus.RawNote.Valid {
			cus.Note = &cus.RawNote.String
		}
	}
	cus.CreatedAt = cus.RawCreatedAt
	cus.UpdatedAt = cus.RawUpdatedAt
	if cus.RawFulfilledAt != nil {
		if cus.RawFulfilledAt.Valid {
			cus.FulfilledAt = &cus.RawFulfilledAt.String
		}
	}
	if cus.RawPaidAt != nil {
		if cus.RawPaidAt.Valid {
			cus.PaidAt = &cus.RawPaidAt.String
		}
	}
	if cus.RawCancelledAt != nil {
		if cus.RawCancelledAt.Valid {
			cus.CancelledAt = &cus.RawCancelledAt.String
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
	CreatedAt         *string `json:"created_at" binding:"omitempty,datetime"`
	UpdatedAt         *string `json:"updated_at" binding:"omitempty"`
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
	Id                *string `form:"id" binding:"omitempty,gt=0"`
	Session           *string `json:"session" binding:"omitempty"`
	CustomerId        *int64  `json:"customer_id" binding:"omitempty,gt=0"`
	PaymentStatus     *string `json:"payment_status" binding:"omitempty"`
	FulfillmentStatus *string `json:"fulfillment_status" binding:"omitempty"`
	Note              *string `json:"note" binding:"omitempty"`
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
