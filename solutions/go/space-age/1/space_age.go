package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    data := make(map[Planet]float64)
    data["Mercury"] = seconds / (31557600.0 * 0.2408467)
    data["Venus"] = seconds / (31557600.0 * 0.61519726)
    data["Earth"] = seconds / 31557600.0
    data["Mars"] = seconds / (31557600.0 * 1.8808158)
    data["Jupiter"] = seconds / (31557600.0 * 11.862615)
    data["Saturn"] = seconds / (31557600.0 * 29.447498)
    data["Uranus"] = seconds / (31557600.0 * 84.016846)
    data["Neptune"] = seconds / (31557600.0 * 164.79132)
    _, exists := data[planet]
    if exists {
		return data[planet]
    }
    return -1
}
