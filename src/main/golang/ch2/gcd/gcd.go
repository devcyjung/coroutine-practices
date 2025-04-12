package gcd

type UnsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func GCD[E UnsignedInt](a, b E) E {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
