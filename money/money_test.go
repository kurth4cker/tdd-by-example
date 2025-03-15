package money_test

import (
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
	"github.com/kurth4cker/tdd-by-example/money"
)

func TestAdd(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		five := money.Dollar(5)
		sum := five.Add(five)
		bank := money.Bank{}

		got := bank.Reduce(sum, "USD")
		want := money.Dollar(10)
		assert.Equal(t, got, want)
	})

	t.Run("Sum", func(t *testing.T) {
		five := money.Dollar(5)
		result := five.Add(five)
		sum, ok := result.(money.Sum)
		assert.True(t, ok)
		assert.Equal(t, sum.Augend, money.Expression(five))
		assert.Equal(t, sum.Addend, money.Expression(five))
	})

	t.Run("Reduce Sum", func(t *testing.T) {
		sum := money.Sum{
			Augend: money.Dollar(3),
			Addend: money.Dollar(4),
		}
		bank := money.Bank{}

		got := bank.Reduce(sum, "USD")
		want := money.Dollar(7)
		assert.Equal(t, got, want)
	})

	t.Run("mixed currencies", func(t *testing.T) {
		fiveDollars := money.Expression(money.Dollar(5))
		tenFrancs := money.Expression(money.Franc(10))
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)

		got := bank.Reduce(fiveDollars.Add(tenFrancs), "USD")
		want := money.Dollar(10)
		assert.Equal(t, got, want)
	})
}

func TestBank_Rate(t *testing.T) {
	t.Run("identity rate", func(t *testing.T) {
		got := money.NewBank().Rate("USD", "USD")
		want := 1
		assert.Equal(t, got, want)
	})
}

func TestBank_Reduce(t *testing.T) {
	bank := money.Bank{}
	got := bank.Reduce(money.Dollar(1), "USD")
	want := money.Dollar(1)
	assert.Equal(t, got, want)
}

func TestMoney_Reduce(t *testing.T) {
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	got := bank.Reduce(money.Franc(2), "USD")
	want := money.Dollar(1)
	assert.Equal(t, got, want)
}

func TestSum_Add(t *testing.T) {
	fiveBucks := money.Dollar(5)
	tenFrancs := money.Franc(10)
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := money.Sum{fiveBucks, tenFrancs}.Add(fiveBucks)

	got := bank.Reduce(sum, "USD")
	want := money.Dollar(15)
	assert.Equal(t, got, want)
}

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
		assertEqualExpression(t, got, want)

		got = five.Times(3)
		want = money.Dollar(15)
		assertEqualExpression(t, got, want)
	})

	t.Run("Franc", func(t *testing.T) {
		five := money.Franc(5)

		got := five.Times(2)
		want := money.Franc(10)
		assertEqualExpression(t, got, want)

		got = five.Times(3)
		want = money.Franc(15)
		assertEqualExpression(t, got, want)
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

func assertEqualExpression(t *testing.T, got, want money.Expression) {
	t.Helper()
	assert.Equal(t, got, want)
}
