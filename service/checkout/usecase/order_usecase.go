package usecase

import (
	"errors"

	"github.com/lonysutrisno/kuncie/pkg"
	"github.com/lonysutrisno/kuncie/service/checkout/repo"
)

func Checkout(args []pkg.FormOrder) (res pkg.Order, err error) {
	// var totalAmount float64
	var products []pkg.Product
	for _, item := range args {
		if item.ProductID == 0 || item.Qty == 0 {
			return res, errors.New("Invalid Request")
		}
		product, err := repo.GetProductDetail(item.ProductID)
		if err != nil {
			return res, err
		}
		if item.Qty > product.Qty {
			return res, errors.New("Insufficient Stock")
		}
		products = append(products, pkg.Product{ID: uint(item.ProductID), Price: product.Price, Qty: item.Qty, Name: product.Name, SKU: product.SKU})

	}
	order, err := ApplyPromo(products)
	if err != nil {
		return res, err
	}

	//store to db
	OrderID, err := repo.CreateOrder(pkg.Order{TotalPrice: order.TotalPrice, Reward: order.Reward})
	if err != nil {
		return res, err
	}
	for _, item := range order.OrderDetail {
		err = repo.CreateOrderDetail(pkg.OrderDetail{OrderID: OrderID, ProductID: int64(item.ID), Qty: item.Qty, Price: item.Price, Reward: item.Reward})
		if err != nil {
			return res, err
		}
		//decduct qty
		err = repo.DeductProductQty(pkg.Product{ID: item.ID, Qty: item.Qty})
		if err != nil {
			return res, err
		}
	}

	return order, err
}
