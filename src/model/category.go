package model

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/constants"
)

type Category struct {
	/** Response Field */
	Id        *int64  `json:"_id"`
	Name      *string `json:"name"`
	ImagePath *string `json:"image_path"`
	ParentId  *int64  `json:"parent_id,omitempty"`
	Status    *string `json:"status"`
	UpdatedAt *string `json:"updated_at"`
	/** Database Field */
	RawId        *int64          `json:"-"`
	RawName      *string         `json:"-"`
	RawImagePath *sql.NullString `json:"-"`
	RawParentId  *sql.NullInt64  `json:"-"`
	RawStatus    *string         `json:"-"`
	RawUpdatedAt *string         `json:"-"`
}

func (cat *Category) FillResponse() {
	cat.Id = cat.RawId
	cat.Name = cat.RawName
	if cat.RawParentId != nil {
		if cat.RawParentId.Valid {
			cat.ParentId = &cat.RawParentId.Int64
		}
	}
	cat.Status = cat.RawStatus
	cat.UpdatedAt = cat.RawUpdatedAt
	if cat.RawImagePath != nil {
		if cat.RawImagePath.Valid {
			cat.ImagePath = &cat.RawImagePath.String
		}
	}

}

type BodyCategory struct {
	Name      *string `json:"name" binding:"omitempty"`
	ImagePath *string `json:"image_path" binding:"omitempty"`
	ParentId  *int64  `json:"parent_id" binding:"omitempty,gte=0"`
}

func (body *BodyCategory) Normalize() error {
	// *body.Name = helpers.SanitizeString(*body.Name)
	// if body.ParentId != nil && *body.ParentId < 1 {
	// 	return errors.New("invalid parent_id")
	// }

	return nil
}

type QueryCategory struct {
	Id            *int64  `form:"id" binding:"omitempty"`
	Name          *string `form:"name" binding:"omitempty"`
	ParentId      *int64  `form:"parent_id" binding:"omitempty"`
	Status        *string `form:"status" binding:"omitempty"`
	UpdatedAtFrom *string `form:"updated_at_from" binding:"omitempty,required_with=UpdatedAtTo,datetime"`
	UpdatedAtTo   *string `form:"updated_at_to" binding:"omitempty,required_with=UpdatedAtFrom,datetime"`
	Page          *int    `form:"page" binding:"omitempty,gte=0"`
	Limit         *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryCategory) ParsePaging() {
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

type CategoryTree struct {
	*Category
	Children []*CategoryTree `json:"children,omitempty"`
}
