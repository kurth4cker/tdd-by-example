package money_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/kurth4cker/tdd-by-example/money"
)

func TestMultiplication(t *testing.T) {
	five := money.Dollar(5)

	got := five.Times(2)
	want := money.Dollar(10)
	assert.Equal(t, want, got)

	got = five.Times(3)
	want = money.Dollar(15)
	assert.Equal(t, want, got)
}
