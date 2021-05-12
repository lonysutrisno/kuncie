package repo

import "github.com/lonysutrisno/kuncie/pkg"

func GetPromoByProduct(productID, qty int64) (res pkg.Promo, err error) {
	rows, err := pkg.DB.Queryx("SELECT * FROM promos where product_id = ? and minimum_purchase <= ?", productID, qty)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			return res, err
		}
	}
	return
}
