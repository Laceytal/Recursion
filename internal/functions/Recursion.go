package functions

import "math/big"

func RecursiveCounting(index int) *big.Int {
	f0 := big.NewInt(1)
	f1 := big.NewInt(3)
	oddValues := []*big.Int{}

	// Проверка начальных значений на нечётность
	if f0.Bit(0) == 1 {
		oddValues = append(oddValues, new(big.Int).Set(f0))
	}
	if f1.Bit(0) == 1 {
		oddValues = append(oddValues, new(big.Int).Set(f1))
	}

	for len(oddValues) <= index {
		fn := new(big.Int).Add(new(big.Int).Mul(big.NewInt(5), f1), f0)
		if fn.Bit(0) == 1 {
			oddValues = append(oddValues, new(big.Int).Set(fn))
		}
		f0, f1 = f1, fn
	}

	return oddValues[index]
}
