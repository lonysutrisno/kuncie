package repo

import (
	"github.com/lonysutrisno/kuncie/pkg"
)

func CreateOrder(args pkg.Order) (OrderID int64, err error) {
	stmt, err := pkg.DB.Prepare("INSERT INTO orders (total_price, reward) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(args.TotalPrice, args.Reward)
	if err != nil {
		return 0, err
	}
	rows, err := pkg.DB.Queryx("SELECT LAST_INSERT_ID();")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&OrderID)
		if err != nil {
			return 0, err
		}
	}
	return
}

func CreateOrderDetail(args pkg.OrderDetail) (err error) {
	stmt, err := pkg.DB.Prepare("INSERT INTO order_details (order_id, product_id, price, qty, reward) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(args.OrderID, args.ProductID, args.Price, args.Qty, args.Reward)

	if err != nil {
		return err
	}

	return
}
