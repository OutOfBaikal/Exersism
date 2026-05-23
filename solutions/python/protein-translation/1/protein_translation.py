ErrStop = StopIteration
ErrInvalidBase = ValueError

def from_codon(codon):
    if len(codon) != 3:
        raise ErrInvalidBase("invalid codon length")
    if codon == "UAA" or codon == "UAG" or codon == "UGA":
        raise ErrStop()
    data = {}
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
    return data[codon]
    

def proteins(strand):
    result = []
    strand_len = len(strand)
    for i in range(0, strand_len, 3):
        codon = strand[i:i+3]
        try:
            data = from_codon(codon)
        except ErrInvalidBase:
            return []
        except ErrStop:
            break
        result.append(data)
    return result
