package ex2product

import (
	"golandtest/log"
	"golandtest/unittest/ex2product/credit"
)

func main() {
	cs := credit.NewService()
	user := &User{ID: "tommy@kkbox.com"}
	product := &Product{
		ID:       "kktv-1-month",
		Price:    100,
		Currency: "TWD",
	}
	err := Purchase(user, product, cs)
	if err != nil {
		log.Fatal("fail to purchase").Err(err).Send()
	}
}
