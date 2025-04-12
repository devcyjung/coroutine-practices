package float

import (
	"github.com/cockroachdb/apd"
)

type Temperature interface {
	float64 | float32 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func FahrToCels[T Temperature](fahrenheit T) (celsius T, err error) {
	ctx := apd.BaseContext
	ctx.Precision = 20
	errDecimal := apd.MakeErrDecimal(&ctx)
	dec := apd.New(0, 0)
	dec, err = dec.SetFloat64(float64(fahrenheit))
	if err != nil {
		return
	}
	errDecimal.Sub(dec, dec, apd.New(32, 0))
	errDecimal.Mul(dec, dec, apd.New(5, 0))
	errDecimal.Quo(dec, dec, apd.New(9, 0))
	if err = errDecimal.Err(); err != nil {
		return
	}
	var cels float64
	cels, err = dec.Float64()
	if err != nil {
		return
	}
	celsius = T(cels)
	return
}

func CelsToFahr[T Temperature](celsius T) (fahrenheit T, err error) {
	ctx := apd.BaseContext
	ctx.Precision = 20
	errDecimal := apd.MakeErrDecimal(&ctx)
	dec := apd.New(0, 0)
	dec, err = dec.SetFloat64(float64(celsius))
	if err != nil {
		return
	}
	errDecimal.Mul(dec, dec, apd.New(9, 0))
	errDecimal.Quo(dec, dec, apd.New(5, 0))
	errDecimal.Add(dec, dec, apd.New(32, 0))
	if err = errDecimal.Err(); err != nil {
		return
	}
	var fahr float64
	fahr, err = dec.Float64()
	if err != nil {
		return
	}
	fahrenheit = T(fahr)
	return
}
