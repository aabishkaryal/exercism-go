package scale

var sharp_scales []string = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

func Scale(tonic, interval string) []string {
	var i int
	if tonic == "C" {
	}
	i := indexSharpScale(tonic)
	return append(scales[i:], scales[:i]...)
}

// Function indexScale returns the index of the scale on the scales list
func indexSharpScale(s string) int {
	for i, scale := range scales {
		if s == scale {
			return i
		}
	}
	return -1
}
