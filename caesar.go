package main

const power = 26

var freqsTable = [power]float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153,
	0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056,
	0.02758, 0.00978, 0.0236, 0.0015, 0.01974, 0.00074}

func caesar(in input) output {
	key := in.Shift
	if !in.Encode {
		key = power - key
	}

	var counts [power]int
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

	var freqs [power]float64
	for i, num := range counts {
		freqs[i] = float64(num) / float64(total)
	}

	shift := 0
	eMin := estimate(&freqs, shift)
	for sh := 1; sh < power; sh++ {
		e := estimate(&freqs, sh)
		if e < eMin {
			eMin = e
			shift = sh
		}
	}

	return output{
		Text:  shiftText(in.Text, key),
		Freqs: freqs,
		Shift: shift,
	}
}

func shiftText(in string, n int) string {
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

func estimate(row *[power]float64, shift int) (e float64) {
	for i := 0; i < power; i++ {
		diff := freqsTable[i] - row[(i+shift)%power]
		e += diff * diff
	}
	return
}
