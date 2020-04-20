package model

import "github.com/barrydev/api-3h-shop/src/constants"

type User struct {
	/** Response Field */
	Id        *int64  `json:"_id,omitempty"`
	Email     *string `json:"email,omitempty"`
	Name      *string `json:"name,omitempty"`
	Password  *string `json:"password,omitempty"`
	Address   *string `json:"address,omitempty"`
	Status    *string `json:"status,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	/** Database Field */
	RawId        *int64  `json:"-"`
	RawEmail     *string `json:"-"`
	RawName      *string `json:"-"`
	RawPassword  *string `json:"-"`
	RawAddress   *string `json:"-"`
	RawStatus    *string `json:"-"`
	RawCreatedAt *string `json:"-"`
	RawUpdatedAt *string `json:"-"`
}

func (user *User) FillResponse() {
	user.Id = user.RawId
	user.Email = user.RawEmail
	user.Password = user.RawPassword
	user.Address = user.RawAddress
	user.Status = user.RawStatus
	user.CreatedAt = user.RawCreatedAt
	user.UpdatedAt = user.RawUpdatedAt
}

type BodyUser struct {
	Id       *int64  `json:"_id" binding:"omitempty,gt=0"`
	Email    *string `json:"email" binding::omitempty,email`
	Name     *string `json:"name" binding::omitempty`
	Password *string `json:"password" binding::omitempty`
	Address  *string `json:"address" binding::omitempty`
	Status   *string `json:"status" binding::omitempty`
}

func (body *BodyUser) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryUser struct {
	Id            *string `form:"id" binding:"omitempty"`
	Email         *string `json:"email" binding::omitempty`
	Name          *string `json:"name" binding::omitempty`
	Password      *string `json:"password" binding::omitempty`
	Address       *string `json:"address" binding::omitempty`
	Status        *string `json:"status" binding::omitempty`
	CreatedAtFrom *string `form:"created_at_from" binding:"omitempty,required_with=CreatedAtTo,datetime"`
	CreatedAtTo   *string `form:"created_at_to" binding:"omitempty,required_with=CreatedAtFrom,datetime"`
	UpdatedAtFrom *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo   *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	Page          *int    `form:"page" binding:"omitempty,gte=0"`
	Limit         *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryUser) ParsePaging() {
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
