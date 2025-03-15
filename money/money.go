package money

type Money struct {
	Currency string
	Amount   int
}

func (m Money) Times(multiplier int) Money {
	return Money{
		Currency: m.Currency,
		Amount:   m.Amount * multiplier,
	}
}

func Currency(money Money) string {
	return money.Currency
}

func Dollar(amount int) Money {
	return Money{
		Currency: "USD",
		Amount:   amount,
	}
}

func Franc(amount int) Money {
	return Money{
		Currency: "CHF",
		Amount:   amount,
	}
}

func IsEqual(m1, m2 Money) bool {
	return m1 == m2
}
