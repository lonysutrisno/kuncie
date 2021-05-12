package pkg

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Up() {
	dbCon := os.Getenv("DB_CONNECTION")
	if dbCon == "" {
		dbCon = "root:123456@tcp(127.0.0.1:33061)/kuncie"
	}
	db, err := gorm.Open("mysql", dbCon)

	if err != nil {
		fmt.Println(err)
	}

	if !db.HasTable("order_details") || !db.HasTable("orders") || !db.HasTable("products") || !db.HasTable("promos") {

		db.CreateTable(&Product{})
		db.CreateTable(&Order{})
		db.CreateTable(&OrderDetail{})
		db.CreateTable(&Promo{})
		var Products = []Product{
			Product{ID: 1, SKU: "GH1", Name: "Google Home", Qty: 100, Price: 49.99},
			Product{ID: 2, SKU: "MBP", Name: "MacbookPro", Qty: 50, Price: 5399.99},
			Product{ID: 3, SKU: "ALX", Name: "Alexa Speaker", Qty: 50, Price: 109.5},
			Product{ID: 4, SKU: "RSPI", Name: "Raspberry Pi b", Qty: 50, Price: 30},
		}
		for _, x := range Products {
			db.Create(&x)
		}
		var Promos = []Promo{
			Promo{PromoType: "bonus", ProductID: 2, Reward: "4", MinimumPurchase: 1},
			Promo{PromoType: "free_item", ProductID: 1, Reward: "1", MinimumPurchase: 3},
			Promo{PromoType: "discount", ProductID: 3, Reward: "10", MinimumPurchase: 4},
		}
		for _, x := range Promos {
			db.Create(&x)
		}
	}
	defer db.Close()
}
