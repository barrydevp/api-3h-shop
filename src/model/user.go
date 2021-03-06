package model

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/constants"
)

var SupportRole = []int64{1, 11, 21}

type User struct {
	/** Response Field */
	Id        *int64  `json:"_id,omitempty"`
	Email     *string `json:"email,omitempty"`
	Name      *string `json:"name,omitempty"`
	Password  *string `json:"password,omitempty"`
	Address   *string `json:"address,omitempty"`
	Role      *int64  `json:"role,omitempty"`
	Session   *string `json:"session,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Status    *string `json:"status,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	/** Database Field */
	RawId        *int64          `json:"-"`
	RawEmail     *string         `json:"-"`
	RawName      *string         `json:"-"`
	RawPassword  *string         `json:"-"`
	RawAddress   *string         `json:"-"`
	RawRole      *int64          `json:"-"`
	RawSession   *sql.NullString `json:"-"`
	RawPhone     *string         `json:"-"`
	RawStatus    *string         `json:"-"`
	RawCreatedAt *string         `json:"-"`
	RawUpdatedAt *string         `json:"-"`
}

func (user *User) FillResponse() {
	user.Id = user.RawId
	user.Email = user.RawEmail
	user.Name = user.RawName
	// user.Password = user.RawPassword
	user.Address = user.RawAddress
	user.Role = user.RawRole
	if user.RawSession != nil {
		if user.RawSession.Valid {
			user.Session = &user.RawSession.String
		}
	}
	user.Phone = user.RawPhone
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
	Session  *string `json:"session" binding:"omitempty"`
	Phone    *string `json:"phone" binding:"omitempty"`
	Role     *int64  `json:"role" binding:"omitempty,gt=0"`
	Status   *string `json:"status" binding::omitempty`
}

type BodyUserChangePassword struct {
	NewPassword *string `json:"new_password" binding::omitempty`
	OldPassword *string `json:"old_password" binding::omitempty`
}

func (body *BodyUser) Normalize() error {
	//*body.Name = helpers.SanitizeString(*body.Name)
	//if body.ParentId != nl && *body.ParentId < 1 {
	//	return errors.New("invalid parent_id")
	//}

	return nil
}

type QueryUser struct {
	Id    *int64  `form:"id" binding:"omitempty"`
	Email *string `json:"email" binding::omitempty`
	Name  *string `json:"name" binding::omitempty`
	Phone *string `json:"phone" binding::omitempty`
	// Password      *string `json:"password" binding::omitempty`
	Address       *string `json:"address" binding::omitempty`
	Role          *string `form:"role" binding:"omitempty"`
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

func IsSupportedRole(role int64) bool {
	for _, supportedRole := range SupportRole {
		if role == supportedRole {
			return true
		}
	}

	return false
}
