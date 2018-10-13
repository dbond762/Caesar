package main

const power = 26

func caesar(in input) output {
	key := in.Shift
	if !in.Encode {
		key = power - key
	}

	counts := make([]int, power)
	total := 0
loop:
	for i := 0; i < len(in.Text); i++ {
		ch := in.Text[i]
		var idx int
		switch {
		case byte('a') <= ch && ch <= byte('z'):
			idx = int(ch) - 'a'
		case byte('A') <= ch && ch <= byte('Z'):
			idx = int(ch) - 'A'
		default:
			continue loop
		}
		total++
		counts[idx]++
	}

	freqs := make([]float64, power)
	for i, num := range counts {
		freqs[i] = float64(num) / float64(total)
	}

	return output{
		Text:  shift(in.Text, key),
		Freqs: freqs,
	}
}

func shift(in string, n int) string {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		ch := in[i]
		switch {
		case byte('a') <= ch && ch <= byte('z'):
			out[i] = byte((int(ch)-'a'+n)%power + 'a')
		case byte('A') <= ch && ch <= byte('Z'):
			out[i] = byte((int(ch)-'A'+n)%power + 'A')
		default:
			out[i] = ch
		}
	}
	return string(out)
}
