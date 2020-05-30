package model

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/constants"
)

type Coupon struct {
	/** Response Field */
	Id          *int64  `json:"_id,omitempty"`
	Code        *string `json:"code,omitempty"`
	Discount    *int64  `json:"discount,omitempty"`
	Description *string `json:"description,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	ExpiresAt   *string `json:"expires_at,omitempty"`
	/** Database Field */
	RawId          *int64  `json:"-"`
	RawCode        *string `json:"-"`
	RawDiscount    *int64  `json:"-"`
	RawDescription *string `json:"-"`
	RawUpdatedAt   *string `json:"-"`
	RawExpiresAt   *string `json:"-"`
}

func (coupon *Coupon) FillResponse() {
	coupon.Id = coupon.RawId
	coupon.Code = coupon.RawCode
	coupon.Description = coupon.RawDescription
	coupon.Discount = coupon.RawDiscount
	coupon.ExpiresAt = coupon.RawExpiresAt
	coupon.UpdatedAt = coupon.RawUpdatedAt
}

type BodyCoupon struct {
	Id          *int64  `json:"_id" binding:"omitempty,gt=0"`
	Code        *string `json:"code" binding:"omitempty"`
	Discount    *int64  `json:"discount" binding:"omitempty,gte=0"`
	Description *string `json:"description" binding:"omitempty"`
	ExpiresAt   *string `json:"expires_at" binding:"omitempty"`
}

func (body *BodyCoupon) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryCoupon struct {
	Id            *int64   `form:"id" binding:"omitempty"`
	Code          *string  `form:"code" binding:"omitempty"`
	Discount      *int64   `form:"discount" binding:"omitempty"`
	Description   *string  `json:"description" binding:"omitempty"`
	UpdatedAtFrom *string  `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo"`
	UpdatedAtTo   *string  `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom"`
	ExpiresAtFrom *string  `form:"expires_at_from" binding:"omitempty,required_with=ExpiresAtTo"`
	ExpiresAtTo   *string  `form:"expires_at_to" binding:"omitempty,required_with=ExpiresAtFrom"`
	Page          *int     `form:"page" binding:"omitempty,gte=0"`
	Limit         *int     `form:"limit" binding:"omitempty,gte=0"`
	Sort          []string `form:"sort[]" binding:"omitempty"`
	Offset        *int
	OrderBy       *string
}

func (query *QueryCoupon) ParsePaging() {
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

func (query *QueryCoupon) ParseSort() {
	var couponBy []string
	if query.Sort != nil {
		if len(query.Sort) > 0 {
			for _, sort := range query.Sort {
				sortArr := strings.Split(sort, " ")
				if len(sortArr) > 0 {
					if len(sortArr) == 1 {
						subCouponBy := sortArr[0] + " DESC"
						couponBy = append(couponBy, subCouponBy)
					} else {
						subCouponBy := sortArr[0]
						typeCoupon := strings.ToLower(sortArr[1])
						if typeCoupon != "asc" && typeCoupon != "desc" {
							typeCoupon = "DESC"
						}
						subCouponBy += " " + typeCoupon
						couponBy = append(couponBy, subCouponBy)
					}
				}
			}
		}
	}

	couponByString := "_id DESC"

	if len(couponBy) > 0 {
		couponByString = strings.Join(couponBy, ", ")
	}

	query.OrderBy = &couponByString
}
