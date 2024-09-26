package node

type number interface {
	int | float64
}

type T interface {
	string | number
}
