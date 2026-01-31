package protein

import (
    "errors"
)

var ErrStop = errors.New("stop error")
var ErrInvalidBase = errors.New("invalid base error")

func FromRNA(rna string) ([]string, error) {
    result := make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
        data, err := FromCodon(rna[i:i+3])
        if err == ErrInvalidBase {
            return []string{}, err
        }
        if err == ErrStop {
            break
        }
        result = append(result, data)
    }

    return uniqueStrings(result), nil
}

func FromCodon(codon string) (string, error) {
	data := make(map[string]string)
    data["AUG"] = "Methionine"
    data["UUU"] = "Phenylalanine"
    data["UUC"] = "Phenylalanine"
    data["UUA"] = "Leucine"
    data["UUG"] = "Leucine"
    data["UCU"] = "Serine"
    data["UCC"] = "Serine"
    data["UCA"] = "Serine"
    data["UCG"] = "Serine"
    data["UAU"] = "Tyrosine"
    data["UAC"] = "Tyrosine"
    data["UGU"] = "Cysteine"
    data["UGC"] = "Cysteine"
    data["UGG"] = "Tryptophan"
    data["UAA"] = "STOP"
    data["UAG"] = "STOP"
    data["UGA"] = "STOP"
    v, e := data[codon]
    if e == false {
        return "", ErrInvalidBase
    }
    if v == "STOP" {
        return "", ErrStop
    }
    return v, nil
}

func uniqueStrings(input []string) []string {
    seen := make(map[string]struct{})
    uniqueSlice := make([]string, 0)

    for _, str := range input {
        if _, exists := seen[str]; !exists {
            seen[str] = struct{}{}
            uniqueSlice = append(uniqueSlice, str)
        }
    }

    return uniqueSlice
}