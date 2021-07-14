package scale

import "strings"

var sharpNotes []string = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatNotes []string = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

// Scale function passes test_cases in scale_generator_test.go.
func Scale(tonic, interval string) []string {
	tonic = ToTitle(tonic)
	switch tonic {
	case "G", "D", "A", "E", "B", "F#", "C", "e", "b", "f#", "c#", "g#", "d#":
		index := SharpNotesIndex(tonic)
		if index != -1 {
			if interval == "" {
				return append(sharpNotes[index:], sharpNotes[:index]...)
			}
			result := make([]string, 0, len(interval))
			for _, i := range interval {
				result = append(result, sharpNotes[index])
				if i == 'M' {
					index += 2
				} else if i == 'm' {
					index += 1
				}
				index %= len(sharpNotes)
			}
			return result
		}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		index := FlatNotesIndex(tonic)
		if index != -1 {
			if interval == "" {
				return append(flatNotes[index:], flatNotes[:index]...)
			}
			result := make([]string, 0, len(interval))
			for _, i := range interval {
				result = append(result, flatNotes[index])
				if i == 'M' {
					index += 2

				} else if i == 'm' {
					index += 1
				}
				index %= len(flatNotes)
			}
			return result
		}
	}
	return []string{}
}

// ToTitle converts the given word by uppercasing the first letter.
func ToTitle(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}

// SharpNotesIndex returns the index of the note on sharpNotes array.
func SharpNotesIndex(note string) int {
	for i, n := range sharpNotes {
		if n == note {
			return i
		}
	}
	return -1
}

// FlatNotesIndex returns the index of the note on the flatNotes array.
func FlatNotesIndex(note string) int {
	for i, n := range flatNotes {
		if n == note {
			return i
		}
	}
	return -1
}
