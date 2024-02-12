// tempconv パッケージは摂氏と華氏と絶対温度の温度計算を行います

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	KelvinOffset  Celsius = 273.15
)

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + KelvinOffset) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
