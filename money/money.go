package money

type Dollar int

func (d Dollar) Times(multiplier int) Dollar {
	return d * Dollar(multiplier)
}
