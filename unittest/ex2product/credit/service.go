//go:generate mockgen -source service.go -destination service_mock.go -package credit
package credit

import (
	"errors"

	"golandtest/unittest/ex2product/credit/bank"
)

var creditStore = map[string]float64{
	"tommy@kkbox.com": 1000.0,
}

var (
	ErrServerTooBusy = errors.New("too busy")
)

type service struct {
}

type Service interface {
	GetByUserID(userID string) (float64, error)
	AddByUserID(userID string, interval float64) error
}

func NewService() Service {
	return &service{}
}

func (s *service) GetByUserID(userID string) (float64, error) {
	if bank.IsTooBusy() {
		return 0, ErrServerTooBusy
	}

	if credit, ok := creditStore[userID]; !ok {
		return 0, errors.New("user not found")
	} else {
		return credit, nil
	}
}

func (s *service) AddByUserID(userID string, interval float64) error {
	if bank.IsTooBusy() {
		return ErrServerTooBusy
	}
	if c, ok := creditStore[userID]; !ok {
		return errors.New("user not found")
	} else {
		creditStore[userID] = c + interval
	}
	return nil
}
