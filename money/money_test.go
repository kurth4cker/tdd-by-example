package money_test

import (
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
	"github.com/kurth4cker/tdd-by-example/money"
)

func TestCurrency(t *testing.T) {
	dollar := money.Dollar(1)
	dollarCurrency := money.Currency(dollar)
	assert.Equal(t, dollarCurrency, "USD")

	franc := money.Franc(1)
	francCurrency := money.Currency(franc)
	assert.Equal(t, francCurrency, "CHF")
}

func TestMultiplication(t *testing.T) {
	t.Run("Dollar", func(t *testing.T) {
		five := money.Dollar(5)

		got := five.Times(2)
		want := money.Dollar(10)
		assert.Equal(t, got, want)

		got = five.Times(3)
		want = money.Dollar(15)
		assert.Equal(t, got, want)
	})

	t.Run("Franc", func(t *testing.T) {
		five := money.Franc(5)

		got := five.Times(2)
		want := money.Franc(10)
		assert.Equal(t, got, want)

		got = five.Times(3)
		want = money.Franc(15)
		assert.Equal(t, got, want)
	})
}

func TestEquality(t *testing.T) {
	five := money.Dollar(5)
	six := money.Dollar(6)
	assert.True(t, money.IsEqual(five, five))
	assert.False(t, money.IsEqual(five, six))

	dollar := money.Dollar(5)
	franc := money.Franc(5)
	assert.False(t, money.IsEqual(dollar, franc))
}
