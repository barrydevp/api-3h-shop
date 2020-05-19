package model

import "github.com/barrydev/api-3h-shop/src/constants"

type Role struct {
	/** Response Field */
	Id   *int64  `json:"_id,omitempty"`
	Name *string `json:"name,omitempty"`
	/** Database Field */
	RawId   *int64  `json:"-"`
	RawName *string `json:"-"`
}

func (role *Role) FillResponse() {
	role.Id = role.RawId
	role.Name = role.RawName
}

type BodyRole struct {
	Id   *int64  `json:"_id" binding:"omitempty,gt=0"`
	Name *string `json:"name" binding::omitempty`
}

func (body *BodyRole) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryRole struct {
	Id     *int64  `form:"id" binding:"omitempty"`
	Name   *string `json:"name" binding::omitempty`
	Page   *int    `form:"page" binding:"omitempty,gte=0"`
	Limit  *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset *int
}

func (query *QueryRole) ParsePaging() {
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
