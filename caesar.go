package main

const (
	power           = 26 // Мощность алфавита
	unicityDistance = 28 // Расстояние единственности
)

var freqsTable = [power]float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153,
	0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056,
	0.02758, 0.00978, 0.0236, 0.0015, 0.01974, 0.00074,
}

func caesar(in input) output {
	// Переводим ключ в [0; power)
	key := in.Shift
	key %= power
	if key < 0 {
		key = power + key
	}

	if !in.Encode {
		key = power - key
	}

	var (
		text = make([]byte, len(in.Text))

		counts       [power]int
		totalLetters = 0
	)

	// В цикле производим сдвиг текста и подсчет символов.
loop:
	for i := 0; i < len(in.Text); i++ {
		var (
			ch       = in.Text[i]
			zeroSing byte
		)

		switch {
		case byte('a') <= ch && ch <= byte('z'):
			zeroSing = 'a'
		case byte('A') <= ch && ch <= byte('Z'):
			zeroSing = 'A'
		default:
			text[i] = ch
			continue loop
		}

		idx := int(ch - zeroSing)
		text[i] = byte((idx+key)%power) + zeroSing

		totalLetters++
		counts[idx]++
	}

	if totalLetters == 0 {
		return output{
			Text:       string(text),
			IsAnalyzed: false,
		}
	}

	// Переводим количества символов в их частоты.
	var freqs [power]float64
	for i, num := range counts {
		freqs[i] = float64(num) / float64(totalLetters)
	}

	if totalLetters < unicityDistance {
		return output{
			Text:       string(text),
			Freqs:      freqs,
			IsAnalyzed: false,
		}
	}

	// Криптоанализ сообщения.
	// Для всех возможных значений сдвига выбираем
	// минимальное значение оценочной функции.
	var (
		shift = 0
		eMin  = estimate(&freqs, shift)
	)
	for sh := 1; sh < power; sh++ {
		e := estimate(&freqs, sh)
		if e < eMin {
			eMin = e
			shift = sh
		}
	}

	return output{
		Text:       string(text),
		Freqs:      freqs,
		Shift:      shift,
		IsAnalyzed: true,
	}
}

// Оценочная функция. \sum_{i=0}^{power} (E_i - R_i)^2 ,
// где E - эталонный ряд, R - рассматриваемый ряд.
func estimate(row *[power]float64, shift int) (e float64) {
	for i := 0; i < power; i++ {
		diff := freqsTable[i] - row[(i+shift)%power]
		e += diff * diff
	}
	return
}
