package alleles

import (
	"math/rand"
	"strconv"
)

// Support for handling alleles of code values where the allele
// takes a random code value and supports mutation to another
// random code value.

// CodeAllele is an allele made over a set of code values
type CodeAllele struct {
	Value        byte
	Translations map[byte]string
}

func (ca CodeAllele) String() string {
	str, ok := ca.Translations[ca.Value]
	if !ok {
		return strconv.FormatUint(uint64(ca.Value), 10)
	}

	return str
}

// CodeFactory selects a code based on the relative frequency of
// the codes.  For example, codes 1, 2, 3 might have a relative
// frequency of 5.0, 3.0, 2.0 which means 1 is chosen about 1/2
// the time, 2, 30% of the time and 2.0 20% of the time.
type CodeFactory struct {
	Codes        []byte
	Translations map[byte]string
	Frequencies  []float64
}

func pickBasedOnFrequency(codes []byte, freqs []float64) byte {
	result := byte(0)

	sumFreq := 0.0
	for _, f := range freqs {
		sumFreq += f
	}

	rand := rand.Float64() * sumFreq

	for i, f := range freqs {
		if rand < f {
			result = codes[i]
			break
		} else {
			rand -= f
		}
	}

	return result

}

// Random produce a random code allele.  Note that there should
// be as many codes as there are frequencies, if frequencies are
// being used.
func (cf CodeFactory) Random() Allele {
	result := CodeAllele{}

	if len(cf.Frequencies) != 0 {
		result.Value = pickBasedOnFrequency(cf.Codes, cf.Frequencies)
	} else {
		rand := rand.Float64() * float64(len(cf.Codes))
		result.Value = cf.Codes[int(rand)]
	}

	return result
}

// CodeMutator mutates a code to another code based on the given
// set of frequencies. Similar to the random function above.
type CodeMutator struct {
	Codes       []byte
	Frequencies []float64
}

// Mutate a given coded allele to another coded allele
func (cm CodeMutator) Mutate(a Allele) Allele {
	result := CodeAllele{}

	result.Translations = a.(CodeAllele).Translations

	if len(cm.Frequencies) != 0 {
		result.Value = pickBasedOnFrequency(cm.Codes, cm.Frequencies)
	} else {
		rand := rand.Float64() * float64(len(cm.Codes))
		result.Value = cm.Codes[int(rand)]
	}

	return result
}
