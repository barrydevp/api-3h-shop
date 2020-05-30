package model

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/constants"
)

type Warranty struct {
	/** Response Field */
	Id          *int64  `json:"_id,omitempty"`
	Code        *string `json:"code,omitempty"`
	Month       *int64  `json:"month,omitempty"`
	Trial       *int64  `json:"trial,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryId  *int64  `json:"category_id,omitempty"`
	/** Database Field */
	RawId          *int64  `json:"-"`
	RawCode        *string `json:"-"`
	RawMonth       *int64  `json:"-"`
	RawTrial       *int64  `json:"-"`
	RawStatus      *string `json:"-"`
	RawDescription *string `json:"-"`
	RawCategoryId  *int64  `json:"-"`
}

func (warranty *Warranty) FillResponse() {
	warranty.Id = warranty.RawId
	warranty.CategoryId = warranty.RawCategoryId
	warranty.Code = warranty.RawCode
	warranty.Description = warranty.RawDescription
	warranty.Month = warranty.RawMonth
	warranty.Trial = warranty.RawTrial
	warranty.Status = warranty.RawStatus
}

type BodyWarranty struct {
	Id          *int64  `json:"_id" binding:"omitempty,gt=0"`
	Code        *string `json:"code" binding:"omitempty"`
	Month       *int64  `json:"month" binding:"omitempty,gte=0"`
	Trial       *int64  `json:"trial" binding:"omitempty,gte=0"`
	Description *string `json:"description" binding:"omitempty"`
	Status      *string `json:"status" binding:"omitempty"`
	CategoryId  *int64  `json:"category_id" binding:"omitempty,gt=0"`
}

func (body *BodyWarranty) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nil && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryWarranty struct {
	Id          *int64   `form:"_id" binding:"omitempty"`
	Code        *string  `form:"code" binding:"omitempty"`
	Month       *int64   `form:"month" binding:"omitempty"`
	Trial       *int64   `form:"trial" binding:"omitempty"`
	Description *string  `json:"description" binding:"omitempty"`
	Status      *string  `json:"status" binding:"omitempty"`
	CategoryId  *int64   `form:"category_id" binding:"omitempty"`
	Page        *int     `form:"page" binding:"omitempty,gte=0"`
	Limit       *int     `form:"limit" binding:"omitempty,gte=0"`
	Sort        []string `form:"sort[]" binding:"omitempty"`
	Offset      *int
	OrderBy     *string
}

func (query *QueryWarranty) ParsePaging() {
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

func (query *QueryWarranty) ParseSort() {
	var warrantyBy []string
	if query.Sort != nil {
		if len(query.Sort) > 0 {
			for _, sort := range query.Sort {
				sortArr := strings.Split(sort, " ")
				if len(sortArr) > 0 {
					if len(sortArr) == 1 {
						subWarrantyBy := sortArr[0] + " DESC"
						warrantyBy = append(warrantyBy, subWarrantyBy)
					} else {
						subWarrantyBy := sortArr[0]
						typeWarranty := strings.ToLower(sortArr[1])
						if typeWarranty != "asc" && typeWarranty != "desc" {
							typeWarranty = "DESC"
						}
						subWarrantyBy += " " + typeWarranty
						warrantyBy = append(warrantyBy, subWarrantyBy)
					}
				}
			}
		}
	}

	warrantyByString := "_id DESC"

	if len(warrantyBy) > 0 {
		warrantyByString = strings.Join(warrantyBy, ", ")
	}

	query.OrderBy = &warrantyByString
}
