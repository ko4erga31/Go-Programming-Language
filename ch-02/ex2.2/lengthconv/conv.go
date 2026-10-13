package lengthconv

const (
	mPerFeet = 0.3048
)

func MToF(m Meters) Feet { return Feet(m / mPerFeet) }
func FToM(f Feet) Meters { return Meters(f * mPerFeet) }
