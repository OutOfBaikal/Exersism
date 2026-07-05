package variablelengthquantity

import (
    "errors"
)

func EncodeVarint(input []uint32) []byte {
    out := make([]byte, 0)
    for _, x := range input {
		// VLQ: 0 -> 0x00
		if x == 0 {
			out = append(out, 0)
			continue
		}

		// Собираем 7-bit группы в обратном порядке
		var tmp [5]byte
		n := 0
		for x > 0 {
			tmp[n] = byte(x & 0x7F)
			n++
			x >>= 7
		}

		// tmp[0] - младшая группа. Для VLQ ставим continuation на все байты,
		// кроме последнего (самого старшего по значению).
		for i := n - 1; i >= 0; i-- {
			b := tmp[i]
			if i != 0 {
				b |= 0x80
			}
			out = append(out, b)
		}
	}
	return out
}

func DecodeVarint(input []byte) ([]uint32, error) {
    out := make([]uint32, 0)

	var (
		result uint32
		//shift  uint
		seen   bool
	)

	for _, b0 := range input {
		b := b0

		seen = true
		result = (result << 7) | uint32(b&0x7F)

		// continuation bit 0 => конец текущего числа
		if b&0x80 == 0 {
			out = append(out, result)
			result = 0
			//shift = 0
			seen = false
		}
	}

	// Если строка закончилась посреди числа (у всех байтов continuation=1)
	if seen {
		return []uint32{}, errors.New("Error!")
	}

    return out, nil
}
