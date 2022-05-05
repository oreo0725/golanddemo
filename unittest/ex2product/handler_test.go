package ex2product

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//TODO revise this test
func TestPurchase(t *testing.T) {

	u := &User{
		ID: "tommy@kkbox.com",
	}
	p := &Product{
		ID:       "toy1",
		Price:    100,
		Currency: "TWD",
	}

	err := Purchase(u, p)
	require.NoError(t, err)

}
