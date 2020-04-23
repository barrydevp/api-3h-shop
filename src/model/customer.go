package model

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/constants"
)

type Customer struct {
	/** Response Field */
	Id        *int64  `json:"_id,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	FullName  *string `json:"full_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	/** Database Field */
	RawId        *int64          `json:"-"`
	RawPhone     *string         `json:"-"`
	RawAddress   *string         `json:"-"`
	RawFullName  *sql.NullString `json:"-"`
	RawEmail     *sql.NullString `json:"-"`
	RawUpdatedAt *string         `json:"-"`
}

func (cus *Customer) FillResponse() {
	cus.Id = cus.RawId
	cus.Phone = cus.RawPhone
	cus.Address = cus.RawAddress
	if cus.RawFullName != nil {
		if cus.RawFullName.Valid {
			cus.FullName = &cus.RawFullName.String
		}
	}
	if cus.RawEmail != nil {
		if cus.RawEmail.Valid {
			cus.Email = &cus.RawEmail.String
		}
	}
	cus.UpdatedAt = cus.RawUpdatedAt

}

type BodyCustomer struct {
	Id       *string `json:"id" binding:"omitempty,gt=0"`
	Phone    *string `json:"phone" binding:"omitempty"`
	Address  *string `json:"address" binding:"omitempty"`
	FullName *string `json:"full_name" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
}

func (body *BodyCustomer) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryCustomer struct {
	Id            *int64  `form:"id" binding:"omitempty"`
	Phone         *string `form:"phone" binding:"omitempty"`
	Address       *string `form:"address" binding:"omitempty"`
	FullName      *string `form:"full_name" binding:"omitempty"`
	Email         *string `form:"email" binding:"omitempty"`
	UpdatedAtFrom *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo   *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	Page          *int    `form:"page" binding:"omitempty,gte=0"`
	Limit         *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryCustomer) ParsePaging() {
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
