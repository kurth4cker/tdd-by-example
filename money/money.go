package money

type Sum struct {
	Augend, Addend Expression
}

func (s Sum) Add(addend Expression) Expression {
	return Sum{
		Augend: s,
		Addend: addend,
	}
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.Augend.Reduce(bank, to).Amount +
		s.Addend.Reduce(bank, to).Amount
	return Money{
		Amount:   amount,
		currency: to,
	}
}

func (s Sum) Times(multiplier int) Expression {
	return Sum{
		Augend: s.Augend.Times(multiplier),
		Addend: s.Addend.Times(multiplier),
	}
}

type Bank struct {
	rates map[pair]int
}

func NewBank() Bank {
	return Bank{
		rates: make(map[pair]int),
	}
}

func (b Bank) AddRate(from, to string, rate int) {
	b.rates[pair{from, to}] = rate
}

func (b Bank) Reduce(expression Expression, currency string) Money {
	return expression.Reduce(b, currency)
}

func (b Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	rate := b.rates[pair{from, to}]
	return rate
}

type Expression interface {
	Add(addend Expression) Expression
	Reduce(bank Bank, currency string) Money
	Times(multiplier int) Expression
}

type Money struct {
	currency string
	Amount   int
}

func (m Money) Add(addend Expression) Expression {
	return Sum{
		Augend: m,
		Addend: addend,
	}
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{
		Amount:   m.Amount / rate,
		currency: to,
	}
}

func (m Money) Times(multiplier int) Expression {
	return Money{
		currency: m.currency,
		Amount:   m.Amount * multiplier,
	}
}

func Currency(money Money) string {
	return money.Currency()
}

func Dollar(amount int) Money {
	return Money{
		currency: "USD",
		Amount:   amount,
	}
}

func Franc(amount int) Money {
	return Money{
		currency: "CHF",
		Amount:   amount,
	}
}

func IsEqual(m1, m2 Money) bool {
	return m1 == m2
}

type pair struct {
	from, to string
}
