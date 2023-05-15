package bighello

type Random interface {
	RandomInt() int
}

type BigHello struct {
	random Random
}

func New(random Random) *BigHello {
	return &BigHello{random}
}

func (bigHello *BigHello) DoSomething(a int, b int) int {
	randomStuff:= bigHello.random.RandomInt()
	return a + b + randomStuff
}
