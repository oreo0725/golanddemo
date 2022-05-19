package ex2product

import (
	"testing"

	"golandtest/unittest/ex2product/credit"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//TODO revise this test
func TestPurchase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCreditService := credit.NewMockService(ctrl)
	mockCreditService.EXPECT().GetByUserID("tommy@kkbox.com").Return(float64(101), nil)
	mockCreditService.EXPECT().AddByUserID("tommy@kkbox.com", float64(-100)).Return(nil)

	u := &User{
		ID: "tommy@kkbox.com",
	}
	p := &Product{
		ID:       "toy1",
		Price:    100,
		Currency: "TWD",
	}

	err := Purchase(u, p, mockCreditService)

	require.NoError(t, err)

}
