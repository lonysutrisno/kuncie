package repo

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lonysutrisno/kuncie/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderDetail(t *testing.T) {
	db, mock, err := sqlmock.New()
	DBx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	pkg.DB = DBx
	mock.ExpectPrepare(regexp.QuoteMeta(`
	INSERT INTO order_details (order_id, product_id, price, qty, reward) VALUES (?, ?, ?, ?, ?)
	`)).ExpectExec().WillReturnResult(driver.RowsAffected(1))

	var data = pkg.OrderDetail{
		ID:      1,
		OrderID: 1,
		Qty:     10000,
		Price:   100,
	}

	//checking repo function
	err = CreateOrderDetail(data)
	if err != nil {
		t.Fatalf("an error '%s' ", err)
	}
	//compare result

	// t.Error(res)
	assert.Equal(t, nil, err)
}
func TestCreateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	DBx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	pkg.DB = DBx
	rows := sqlmock.NewRows([]string{"last_id"}).AddRow(1)
	mock.ExpectPrepare(regexp.QuoteMeta(`
	INSERT INTO orders (total_price, reward) VALUES (?, ?)
	`)).ExpectExec().WillReturnResult(driver.RowsAffected(1))

	mock.ExpectQuery(regexp.QuoteMeta(`
	SELECT LAST_INSERT_ID();
	`)).WillReturnRows(rows).RowsWillBeClosed()
	var data = pkg.Order{
		TotalPrice: 1000,
		Reward:     "buy 1 get 1",
	}

	//checking repo function
	id, err := CreateOrder(data)
	if err != nil {
		t.Fatalf("an error '%s' ", err)
	}
	//compare result

	// t.Error(res)
	assert.Equal(t, id, int64(1))
	assert.Equal(t, nil, err)
}
