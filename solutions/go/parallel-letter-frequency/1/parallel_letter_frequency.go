package letter

import (
    "sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
    var wg sync.WaitGroup
    ch := make(chan FreqMap, len(texts))
    
	for _, text := range texts {
        wg.Add(1)
        go func(t string) {
            defer wg.Done()
            ch <- Frequency(t)
        }(text)
    } 

    go func() {
        wg.Wait()
        close(ch)
    }()

    result := FreqMap{}
    for freqMap := range ch {
        for r, count := range freqMap {
            result[r] += count
        }
    }

    return result;
}
