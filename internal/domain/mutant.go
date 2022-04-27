package domain

import (
	"strings"
)

type (
	DNAType    string
	DNA        []string
	HumanError struct{}
)

func (h HumanError) Error() string {
	return "dna is human"
}

const (
	HumanDNA  DNAType = "Human"
	MutantDNA DNAType = "Mutant"
)

type Mutant struct {
	ID   string  `json:"id"`
	DNA  DNA     `json:"dna"`
	Type DNAType `json:"type"`
}

type Stats struct {
	CountMutantDNA int64   `json:"count_mutant_dna"`
	CountHumanDNA  int64   `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

func IsMutant(dna DNA) bool {
	mutantChan := make(chan int)
	humanChan := make(chan int)

	go verifyDNA(dna, mutantChan, humanChan)

	select {
	case <-mutantChan:
		return true
	case <-humanChan:
		return false
	}
}

func verifyDNA(dnas DNA, mutantChan, humanChan chan int) {
	horizontals := make(map[int]string)
	diagonals := make(map[int]string)
	dnaLength := 5

	for line, dna := range dnas {
		verifyRepetitionInString(dna, mutantChan)
		lettersDNA := strings.Split(dna, "")
		for column, letterDNA := range lettersDNA {
			horizontals[column] += letterDNA
			if line == column {
				diagonals[0] += letterDNA
			}
			if column == dnaLength {
				diagonals[1] += lettersDNA[dnaLength]
			}
		}
		dnaLength--
	}

	verifyRepetitionInMap(horizontals, mutantChan)
	verifyRepetitionInMap(diagonals, mutantChan)

	humanChan <- 1
}

func verifyRepetitionInMap(dnas map[int]string, mutantChan chan int) {
	for _, dna := range dnas {
		verifyRepetitionInString(dna, mutantChan)
	}
}

func verifyRepetitionInString(dna string, mutantChan chan int) {
	const maxMutantCount = 3

	dnaArray := strings.Split(dna, "")
	prevLetter := ""
	mutantCount := 0

	for i := 0; i < len(dnaArray); i++ {
		actualLetter := dnaArray[i]
		if prevLetter == actualLetter {
			mutantCount++
		} else {
			mutantCount = 0
		}

		if mutantCount == maxMutantCount {
			mutantChan <- 1
		}

		prevLetter = actualLetter
	}
}
