package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ApplyCoupon(orderId int64, body *model.BodyCoupon) (*model.Coupon, error) {

	if body.Code == nil {
		return nil, errors.New("coupon's code is required")
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		existCoupon, err := factories.FindOneCoupon(&connect.QueryMySQL{
			QueryString: "WHERE code=? AND expires_at>=NOW()",
			Args:        []interface{}{body.Code},
		})

		if err != nil {
			rejectChan <- err

			return
		}
		if existCoupon == nil {
			rejectChan <- errors.New("coupon does not exists or has been expires")

			return
		}

		resolveChan <- existCoupon
	}()

	go func() {
		existOrder, err := factories.FindOneOrder(&connect.QueryMySQL{
			QueryString: "WHERE _id=? AND payment_status='pending'",
			Args:        []interface{}{&orderId},
		})

		if err != nil {
			rejectChan <- err

			return
		}

		if existOrder == nil {
			rejectChan <- errors.New("order does not exists or has been checkout")

			return
		}

		resolveChan <- existOrder
	}()

	var coupon *model.Coupon

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case *model.Order:
			case *model.Coupon:
				coupon = val
			}

		case err := <-rejectChan:
			return nil, err
		}
	}

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
		QueryString: `SET coupon_id=?, status="apply_coupon" WHERE _id=?`,
		Args:        []interface{}{coupon.Id, orderId},
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return coupon, nil
}
