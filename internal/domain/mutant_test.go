package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutant(t *testing.T) {
	mutants := []struct{
		dna    DNASequence
		result bool
	}{
		{[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}, true}, //vertical
		{[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCTTA", "TCACTG"}, true}, //left diagonal
		{[]string{"TTGCGA", "TAGTGC", "TTATGT", "TGAGAG", "CCCTTA", "TCACTG"}, true}, //horizontal
		{[]string{"ATGCGA", "CAGTAC", "TTAAGT", "AGAGAG", "CCCTTA", "TCACTG"}, true}, //right diagonal
	}
	for _, mutant := range mutants {
		t.Run(fmt.Sprintf("test with %v is mutant expected %v", mutant.dna, mutant.result), func(t *testing.T) {
			m := IsMutant(mutant.dna)

			assert.Equal(t, mutant.result, m)
		})
	}

	humans := []struct{
		dna    DNASequence
		result bool
	}{
		{[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAGAG", "CCCTTA", "TCACTG"}, false},
		{[]string{"ATGCGA", "CAGTGC", "GTATGT", "AGAGAG", "CCCTTA", "TCACTG"}, false},
		{[]string{"ACGCGA", "CAGTGC", "TTATGT", "AGAGAG", "CCCTTA", "TCACTG"}, false},
	}
	for _, human := range humans {
		t.Run(fmt.Sprintf("test with %v is human expected %v", human.dna, human.result), func(t *testing.T) {
			m := IsMutant(human.dna)
			assert.Equal(t, human.result, m)
		})
	}
}
