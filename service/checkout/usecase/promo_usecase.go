package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lonysutrisno/kuncie/pkg"
	"github.com/lonysutrisno/kuncie/service/checkout/repo"
)

func ApplyPromo(args []pkg.Product) (res pkg.Order, err error) {
	var totalAmount float64
	var orderDetails []pkg.OrderDetail
	var reward []string
	for _, item := range args {
		promo, err := repo.GetPromoByProduct(int64(item.ID), item.Qty)
		if err != nil {
			return res, err
		}
		if promo.ID != 0 {
			switch promo.PromoType {
			case "free_item":
				totalProduct, err := strconv.Atoi(promo.Reward)
				if err != nil {
					return res, err
				}
				product, err := repo.GetProductDetail(int64(promo.ProductID))
				if err != nil {
					return res, err
				}
				rewardstr := fmt.Sprintf("Get Price for %+v pcs on %+v ", item.Qty-int64(totalProduct), product.Name)
				price := (item.Price * float64(item.Qty-int64(totalProduct)))
				reward = append(reward, rewardstr)
				totalAmount = totalAmount + price

				orderDetails = append(orderDetails, pkg.OrderDetail{ProductID: int64(item.ID), Qty: item.Qty, Price: price, Reward: rewardstr, SKU: item.SKU, ProductName: item.Name})
			case "bonus":
				productReward, err := strconv.Atoi(promo.Reward)
				if err != nil {
					return res, err
				}
				product, err := repo.GetProductDetail(int64(productReward))
				if err != nil {
					return res, err
				}
				rewardstr := "Bonus 1 pcs " + product.Name
				price := (item.Price * float64(item.Qty))

				reward = append(reward, rewardstr)
				totalAmount = totalAmount + price

				orderDetails = append(orderDetails, pkg.OrderDetail{ProductID: int64(item.ID), Qty: item.Qty, Price: price, Reward: rewardstr, SKU: item.SKU, ProductName: item.Name})
			case "discount":
				discount, err := strconv.Atoi(promo.Reward)
				if err != nil {
					return res, err
				}
				product, err := repo.GetProductDetail(int64(promo.ProductID))
				if err != nil {
					return res, err
				}
				rewardstr := "Discount on " + product.Name + " " + promo.Reward + "%"
				price := ((item.Price * float64(item.Qty)) - (item.Price * float64(item.Qty) * float64(discount) / 100))

				reward = append(reward, rewardstr)
				totalAmount = totalAmount + price
				orderDetails = append(orderDetails, pkg.OrderDetail{ProductID: int64(item.ID), Qty: item.Qty, Price: price, Reward: rewardstr, SKU: item.SKU, ProductName: item.Name})

			}
		} else {
			totalAmount = totalAmount + (item.Price * float64(item.Qty))
			orderDetails = append(orderDetails, pkg.OrderDetail{ProductID: int64(item.ID), Qty: item.Qty, Price: (item.Price * float64(item.Qty)), SKU: item.SKU, ProductName: item.Name})

		}
	}
	res.TotalPrice = totalAmount
	res.Reward = strings.Join(reward, ",")
	res.OrderDetail = orderDetails
	return
}
