package senml

// Unit is the unit for a measurement value.
type Unit string

const (
	Meter     Unit = "m"
	Kilogram  Unit = "kg"
	Second    Unit = "s"
	Ampere    Unit = "A"
	Kelvin    Unit = "K"
	Candela   Unit = "cd"
	Mole      Unit = "mol"
	Hertz     Unit = "Hz"
	Radian    Unit = "rad"
	Steradian Unit = "sr"
	Newton    Unit = "N"
	Pascal    Unit = "Pa"
	Joule     Unit = "J"
	Watt      Unit = "W"
	Coulomb   Unit = "C"
	Volt      Unit = "V"
	Farad     Unit = "F"
	Ohm       Unit = "Ohm"
	Siemens   Unit = "S"
	Weber     Unit = "Wb"
	Tesla     Unit = "T"
	Henry     Unit = "H"
	// degrees Celsius
	Celsius   Unit = "Cel"
	Lumen     Unit = "lm"
	Lux       Unit = "lux"
	Becquerel Unit = "Bq"
	Gray      Unit = "Gy"
	Sievert   Unit = "Sv"
	Katal     Unit = "kat"
	// square meter (area)
	SquareMeter Unit = "m2"
	// cubic meter (volume)
	CubicMeter Unit = "m3"
	// liter (volume)
	Liter Unit = "l"
	// meter per second (velocity)
	MeterPerSecond Unit = "m/s"
	// meter per square second (acceleration)
	MeterPerSquareSecond Unit = "m/s2"
	// cubic meter per second (flow rate)
	CubicMeterPerSecond Unit = "m3/s"
	// liter per second (flow rate)
	LiterPerSecond Unit = "l/s"
	// watt per square meter (irradiance)
	WattPerSquareMeter Unit = "W/m2"
	// candela per square meter (luminance)
	CandelaPerSquareMeter Unit = "cd/m2"
	// bit (information content)
	Bit Unit = "bit"
	// bit per second (data rate)
	BitPerSecond Unit = "bit/s"
	// degrees latitude
	DegreesLatitude Unit = "lat"
	// degrees longitude
	DegreesLongitude Unit = "lon"
	// pH value (acidity; logarithmic quantity)
	PH Unit = "pH"
	// decibel (logarithmic quantity)
	Decibel Unit = "dB"
	// decibel relative to 1 W (power level)
	Decibel1W Unit = "dBW"
	// bel (sound pressure level; logarithmic quantity)
	Bel Unit = "Bspl"
	// 1 (counter value)
	Count Unit = "count"
	// 1 (Ratio e.g., value of a switch)
	Switch Unit = "/"
	// percentage
	Percentage Unit = "%"
	// percentage (Relative Humidity)
	RelativeHumidity Unit = "%RH"
	// percentage (remaining battery energy level)
	EnergyLevel Unit = "%EL"
	// seconds (remaining battery energy level)
	EnergyRemaining Unit = "EL"
	// 1 per second (event rate)
	EventRate Unit = "1/s"
	// 1 per minute (Heart rate in beats per minute)
	BeatsPerMinute Unit = "beat/min"
	// 1 (Cumulative number of heart beats)
	Beats Unit = "beats"
	// siemens per meter (conductivity)
	SiemensPerMeter Unit = "S/m"
)
