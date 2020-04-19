package model

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/helpers"
)

const (
	DEFAULT_PAGE  = 1
	DEFAULT_LIMIT = 20
)

type Category struct {
	/** Response Field */
	Id        *int64  `json:"_id"`
	Name      *string `json:"name"`
	ParentId  *int64  `json:"parent_id,omitempty"`
	Status    *string `json:"status"`
	UpdatedAt *string `json:"updated_at"`
	/** Database Field */
	RawId        *int64         `json:"-"`
	RawName      *string        `json:"-"`
	RawParentId  *sql.NullInt64 `json:"-"`
	RawStatus    *string        `json:"-"`
	RawUpdatedAt *string        `json:"-"`
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

}

type BodyCategory struct {
	Name     *string `json:"name" binding:"required"`
	ParentId *int64  `json:"parent_id" binding:"omitempty,gt=0"`
}

func (body *BodyCategory) Normalize() error {
	*body.Name = helpers.SanitizeString(*body.Name)
	if body.ParentId != nil && *body.ParentId < 1 {
		return errors.New("invalid parent_id")
	}

	return nil
}

type QueryCategory struct {
	Id            *string `form:"id" binding:"omitempty,gt=0"`
	Name          *string `form:"name" binding:"omitempty"`
	ParentId      *int64  `form:"parent_id" binding:"omitempty,gt=0"`
	Status        *string `form:"status" binding:""`
	UpdatedAtFrom *string `form:"updated_at_from" binding:"required_with=UpdatedAtTo"`
	UpdatedAtTo   *string `form:"updated_at_to" binding:"required_with=UpdatedAtFrom"`
	Page          *int    `form:"page" binding:"omitempty,gte=0"`
	Limit         *int    `form:"limit" binding:"omitempty,gte=0"`
	Offset        *int
}

func (query *QueryCategory) GetQueryList() (*connect.QueryMySQL, error) {
	Where := ""
	var Args []interface{}

	if query.Id != nil {
		Where += " _id=?"
		Args = append(Args, query.Id)
	}
	if query.Name != nil {
		Where += " name LIKE ?"
		Args = append(Args, "%"+ *query.Name + "%")
	}
	if query.ParentId != nil {
		Where += " parent_id=?"
		Args = append(Args, query.ParentId)
	}
	if query.Status != nil {
		Where += " status=?"
		Args = append(Args, query.Status)
	}
	if query.UpdatedAtFrom != nil && query.UpdatedAtTo != nil {
		Where += " updated_at BETWEEN ? AND ?"
		Args = append(Args, query.UpdatedAtFrom, query.UpdatedAtTo)
	}

	OrderBy := " _id ASC"

	query.ParsePaging()
	Offset := " ?"
	Args = append(Args, query.Offset)
	Limit := " ?"
	Args = append(Args, query.Limit)

	return &connect.QueryMySQL{
		Where:   &Where,
		OrderBy: &OrderBy,
		Limit:   &Limit,
		Offset:  &Offset,
		Args:    Args,
	}, nil
}

func (query *QueryCategory) GetQueryCountList() (*connect.QueryMySQL, error) {
	Where := ""
	var Args []interface{}

	if query.Id != nil {
		Where += " _id=?"
		Args = append(Args, query.Id)
	}
	if query.Name != nil {
		Where += " name LIKE ?"
		Args = append(Args, "%"+ *query.Name + "%")
	}
	if query.ParentId != nil {
		Where += " parent_id=?"
		Args = append(Args, query.ParentId)
	}
	if query.Status != nil {
		Where += " status=?"
		Args = append(Args, query.Status)
	}
	if query.UpdatedAtFrom != nil && query.UpdatedAtTo != nil {
		Where += " updated_at BETWEEN ? AND ?"
		Args = append(Args, query.UpdatedAtFrom, query.UpdatedAtTo)
	}

	return &connect.QueryMySQL{
		Where:   &Where,
		Args:    Args,
	}, nil
}

func (query *QueryCategory) ParsePaging() {
	if query.Page == nil {
		page := DEFAULT_PAGE
		query.Page = &page
	}

	if query.Limit == nil {
		limit := DEFAULT_LIMIT
		query.Limit = &limit
	}

	skip := (*query.Page - 1) * *query.Limit

	query.Offset = &skip
}

type CategoryTree struct {
	*Category
	Children []*CategoryTree `json:"children,omitempty"`
}
