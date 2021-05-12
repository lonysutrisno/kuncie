package repo

import "github.com/lonysutrisno/kuncie/pkg"

func CreateProduct(args pkg.Product) (err error) {
	stmt, err := pkg.DB.Prepare("INSERT INTO products (name, qty, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(args.Name, args.Qty, args.Price)

	if err != nil {
		return err
	}

	return
}
func GetProductDetail(productID int64) (res pkg.Product, err error) {
	rows, err := pkg.DB.Queryx("SELECT * FROM products where id = ?", productID)
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
func DeductProductQty(args pkg.Product) (err error) {
	stmt, err := pkg.DB.Prepare("UPDATE products SET  qty = qty - ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(args.Qty, args.ID)

	if err != nil {
		return err
	}
	return
}
