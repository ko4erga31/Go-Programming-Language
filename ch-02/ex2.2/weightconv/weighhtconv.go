package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }
