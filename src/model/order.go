package model

import (
	"database/sql"
	"strings"

	"github.com/barrydev/api-3h-shop/src/constants"
)

var ValidFulfillmentStatus = []string{"pending", "fulfilled", "in-production", "shipped", "cancelled"}
var ValidPaymentStatus = []string{"pending", "paid"}

type Order struct {
	/** Response Field */
	Id                *int64   `json:"_id,omitempty"`
	Session           *string  `json:"session,omitempty"`
	CustomerId        *int64   `json:"customer_id,omitempty"`
	Status            *string  `json:"status,omitempty"`
	TotalPrice        *float64 `json:"total_price,omitempty"`
	PaymentStatus     *string  `json:"payment_status,omitempty"`
	FulfillmentStatus *string  `json:"fulfillment_status,omitempty"`
	Note              *string  `json:"note,omitempty"`
	CreatedAt         *string  `json:"created_at,omitempty"`
	UpdatedAt         *string  `json:"updated_at,omitempty"`
	PaidAt            *string  `json:"paid_at,omitempty"`
	FulfilledAt       *string  `json:"fulfilled_at,omitempty"`
	CancelledAt       *string  `json:"cancelled_at,omitempty"`
	CouponId          *int64   `json:"coupon_id,omitempty"`
	/** Database Field */
	RawId                *int64           `json:"-"`
	RawSession           *string          `json:"-"`
	RawCustomerId        *sql.NullInt64   `json:"-"`
	RawStatus            *string          `json:"-"`
	RawTotalPrice        *sql.NullFloat64 `json:"-"`
	RawPaymentStatus     *string          `json:"-"`
	RawFulfillmentStatus *string          `json:"-"`
	RawNote              *sql.NullString  `json:"-"`
	RawCreatedAt         *string          `json:"-"`
	RawUpdatedAt         *string          `json:"-"`
	RawPaidAt            *sql.NullString  `json:"-"`
	RawFulfilledAt       *sql.NullString  `json:"-"`
	RawCancelledAt       *sql.NullString  `json:"-"`
	RawCouponId          *sql.NullInt64   `json:"-"`
}

func (order *Order) FillResponse() {
	order.Id = order.RawId
	order.Session = order.RawSession
	if order.RawCustomerId != nil {
		if order.RawCustomerId.Valid {
			order.CustomerId = &order.RawCustomerId.Int64
		}
	}
	if order.RawCouponId != nil {
		if order.RawCouponId.Valid {
			order.CouponId = &order.RawCouponId.Int64
		}
	}
	order.Status = order.RawStatus
	if order.RawTotalPrice != nil {
		if order.RawTotalPrice.Valid {
			order.TotalPrice = &order.RawTotalPrice.Float64
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
	Id                *int64   `json:"_id" binding:"omitempty,gt=0"`
	Session           *string  `json:"session" binding:"omitempty"`
	CustomerId        *int64   `json:"customer_id" binding:"omitempty,gt=0"`
	CouponId          *int64   `json:"coupon_id" binding:"omitempty,gt=0"`
	Status            *string  `json:"status" binding:"omitempty"`
	TotalPrice        *float64 `json:"total_price" binding:"omitempty,gte=0"`
	PaymentStatus     *string  `json:"payment_status" binding:"omitempty"`
	FulfillmentStatus *string  `json:"fulfillment_status" binding:"omitempty"`
	Note              *string  `json:"note" binding:"omitempty"`
	PaidAt            *string  `json:"paid_at" binding:"omitempty"`
	FulfilledAt       *string  `json:"fulfilled_at" binding:"omitempty"`
	CancelledAt       *string  `json:"cancelled_at" binding:"omitempty"`
}

func (body *BodyOrder) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryOrder struct {
	Id                *int64   `form:"_id" binding:"omitempty"`
	Session           *string  `form:"session" binding:"omitempty"`
	CustomerId        *int64   `form:"customer_id" binding:"omitempty"`
	CouponId          *int64   `form:"coupon_id" binding:"omitempty"`
	Status            *string  `json:"status" binding:"omitempty"`
	PaymentStatus     *string  `form:"payment_status" binding:"omitempty"`
	FulfillmentStatus *string  `form:"fulfillment_status" binding:"omitempty"`
	Note              *string  `form:"note" binding:"omitempty"`
	CreatedAtFrom     *string  `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo"`
	CreatedAtTo       *string  `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom"`
	UpdatedAtFrom     *string  `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo"`
	UpdatedAtTo       *string  `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom"`
	PaidAtFrom        *string  `form:"paid_at_from" binding:"omitempty,required_with=PaidAtTo"`
	PaidAtTo          *string  `form:"paid_at_to" binding:"omitempty,required_with=PaidAtFrom"`
	FulfilledAtFrom   *string  `form:"fulfilled_at_from" binding:"omitempty,required_with=FulfilledAtTo"`
	FulfilledAtTo     *string  `form:"fulfilled_at_to" binding:"omitempty,required_with=FulfilledAtFrom"`
	CancelledAtFrom   *string  `form:"cancelled_at_from" binding:"omitempty,required_with=CancelledAtTo"`
	CancelledAtTo     *string  `form:"cancelled_at_to" binding:"omitempty,required_with=CancelledAtFrom"`
	Page              *int     `form:"page" binding:"omitempty,gte=0"`
	Limit             *int     `form:"limit" binding:"omitempty,gte=0"`
	Sort              []string `form:"sort[]" binding:"omitempty"`
	Offset            *int
	OrderBy           *string
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

func (query *QueryOrder) ParseSort() {
	var orderBy []string
	if query.Sort != nil {
		if len(query.Sort) > 0 {
			for _, sort := range query.Sort {
				sortArr := strings.Split(sort, " ")
				if len(sortArr) > 0 {
					if len(sortArr) == 1 {
						subOrderBy := sortArr[0] + " DESC"
						orderBy = append(orderBy, subOrderBy)
					} else {
						subOrderBy := sortArr[0]
						typeOrder := strings.ToLower(sortArr[1])
						if typeOrder != "asc" && typeOrder != "desc" {
							typeOrder = "DESC"
						}
						subOrderBy += " " + typeOrder
						orderBy = append(orderBy, subOrderBy)
					}
				}
			}
		}
	}

	orderByString := "updated_at DESC"

	if len(orderBy) > 0 {
		orderByString = strings.Join(orderBy, ", ")
	}

	query.OrderBy = &orderByString
}

type BodyCheckoutOrder struct {
	*Customer   `json:"customer" binding:"omitempty"`
	Note        *string `json:"note" binding:"omitempty"`
	PaymentType *int    `json:"payment_type" binding:"omitempty"`
}

type QueryStatisticOrder struct {
	CreatedAtFrom *string `form:"created_at_from" binding:"omitempty"`
	CreatedAtTo   *string `form:"created_at_to" binding:"omitempty"`
}

type StatisticOrder struct {
	Pending      int64 `json:"pending"`
	Paid         int64 `json:"paid"`
	InProduction int64 `json:"in_production"`
	Shipped      int64 `json:"shipped"`
	Cancelled    int64 `json:"cancelled"`
	Fulfilled    int64 `json:"fulfilled"`
	TotalOrder   int64 `json:"total_order"`
}

func IsValidFulfillmentStatus(status string) bool {
	for _, validStatus := range ValidFulfillmentStatus {
		if validStatus == status {
			return true
		}
	}

	return false
}

func IsValidPaymentStatus(status string) bool {
	for _, validStatus := range ValidPaymentStatus {
		if validStatus == status {
			return true
		}
	}

	return false
}
