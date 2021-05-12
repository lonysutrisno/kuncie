package pkg

type (
	Order struct {
		ID          uint    `json:"id,omitempty" db:"id"`
		TotalPrice  float64 `json:"price" db:"price"`
		Reward      string  `json:"reward" db:"reward"`
		OrderDetail []OrderDetail
	}
	OrderDetail struct {
		ID          uint    `json:"id,omitempty" db:"id"`
		OrderID     int64   `json:"order_id,omitempty" db:"order_id"`
		ProductID   int64   `json:"product_id" db:"product_id"`
		ProductName string  `json:"product_name" `
		SKU         string  `json:"sku" `
		Qty         int64   `json:"qty" db:"qty"`
		Price       float64 `json:"price" db:"price"`
		Reward      string  `json:"reward" db:"reward"`
	}
	Product struct {
		ID    uint    `json:"id" db:"id"`
		SKU   string  `json:"sku" db:"sku"`
		Name  string  `json:"name" db:"name"`
		Qty   int64   `json:"qty" db:"qty"`
		Price float64 `json:"price" db:"price"`
	}
	Promo struct {
		ID              uint   `json:"id" db:"id"`
		ProductID       int64  `json:"product_id" db:"product_id"`
		MinimumPurchase int64  `json:"minimum_purchase" db:"minimum_purchase"`
		PromoType       string `json:"promo_type" db:"promo_type"`
		Reward          string `json:"reward" db:"reward"`
	}
	FormOrder struct {
		ProductID int64 `json:"product_id"`
		Qty       int64 `json:"qty"`
	}
)
