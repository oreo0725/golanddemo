package ex2product

import (
	"errors"
	"fmt"

	"golandtest/log"
	"golandtest/unittest/ex2product/credit"
)

type User struct {
	ID string
}

type Product struct {
	ID       string
	Price    float64
	Currency string
}

func Purchase(u *User, p *Product, cs credit.Service) error {
	creditService := cs

	ct, err := creditService.GetByUserID(u.ID)
	if err != nil {
		return fmt.Errorf("%w: credit service error", err)
	}
	if ct-p.Price < 0 {
		return errors.New("no enough credit")
	}
	if err := creditService.AddByUserID(u.ID, -p.Price); err != nil {
		return err
	}
	return addOrder(u, p)
}

func addOrder(u *User, p *Product) error {
	log.Info("add order").Str("user_id", u.ID).Interface("product", p).Send()
	return nil
}
