package weightconv

const (
	kgPerPound = 0.45359237
)

func KToP(k Kilograms) Pounds { return Pounds(k / kgPerPound) }
func PToK(p Pounds) Kilograms { return Kilograms(p * kgPerPound) }
