package cipher
//import "fmt"
// Define the shift and vigenere types here.
type shift int
type vigenere string
// Both types should satisfy the Cipher interface.
// type Cipher interface {
//     Encode(input string) string
//     Decode(input string) string
// }

func NewCaesar() Cipher {
	return shift(3)
}

func NewShift(distance int) Cipher {
    if distance < -25 || distance == 0 || distance > 25 {
        return nil
    }
	return shift(distance)
}

func (c shift) Encode(input string) string {
	data := []byte{}
    for i := 0; i < len(input); i++ {
        if input[i] >= 'A' && input[i] <= 'Z' {
        	data = append(data, (input[i] - 'A' + byte(c) + 26) % 26 + 'a')
        } else if input[i] >= 'a' && input[i] <= 'z' {
        	data = append(data, (input[i] - 'a' + byte(c) + 26) % 26 + 'a')
        }
    }
    return string(data)
}

func (c shift) Decode(input string) string {
	data := []byte{}
    for i := 0; i < len(input); i++ {
        if input[i] >= 'A' && input[i] <= 'Z' {
        	data = append(data, (input[i] - 'A' - byte(c) + 26) % 26 + 'a')
        } else if input[i] >= 'a' && input[i] <= 'z' {
        	data = append(data, (input[i] - 'a' - byte(c) + 26) % 26 + 'a')
        }
    }
    return string(data)
}

func NewVigenere(key string) Cipher {
    //key = strings.ReplaceAll(key, " ", "")
    if key == "" || len(key) < 5 {
        return nil
    }
    for _, x := range key {
        if x < 'a' || x > 'z' {
            return nil
        }
    }
    return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	data := []byte{}
    key := string(v) 
    
    keyLen := len(key)
    j := 0
    for i := 0; i < len(input); i++ {
        if (input[i] >= 'A' && input[i] <= 'Z') || (input[i] >= 'a' && input[i] <= 'z') {
            c := key[j%keyLen] - 'a'
            if input[i] <= 'Z' {
                data = append(data, (input[i]-'A'+c)%26+'a')
            } else {
                data = append(data, (input[i]-'a'+c)%26+'a')
            }
            j++
        }
    }
    return string(data)
}

func (v vigenere) Decode(input string) string {
	data := []byte{}
    key := string(v) 
    keyLen := len(key)
    j := 0
    for i := 0; i < len(input); i++ {
        c := key[j % keyLen] - 'a'
    	if input[i] >= 'A' && input[i] <= 'Z' {
        	data = append(data, (input[i] - 'A' - c + 26) % 26 + 'a')
        } else if input[i] >= 'a' && input[i] <= 'z' {
        	data = append(data, (input[i] - 'a' - c + 26) % 26 + 'a')
        }
        j++
    }
    return string(data)
}
