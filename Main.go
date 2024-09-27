package main

const boilingKelvin = 373.0

func main() {
	temperatureKelvin := boilingKelvin
	temperatureCelsius := temperatureKelvin - 273.0

	println("A temperatura de ebulição da água em Kelvin é", temperatureKelvin)
	println("A temperatura de ebulição da água em Celsius é", temperatureCelsius)
}
