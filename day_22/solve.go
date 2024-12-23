package day22

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TotalSteps  = 2000
	SequenceLen = 4
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_22/input.txt"
	if isTest {
		f = "day_22/input-test.txt"
	}

	body, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	switch part {
	case 1:
		return partOne(string(body))
	case 2:
		return partTwo(string(body))
	}

	return nil, fmt.Errorf("part should be only 1 or 2")
}

func partOne(data string) (any, error) {
	initialSecretNums := parseInput(data)
	totalSum := 0
	for _, secret := range initialSecretNums {
		for i := 0; i < TotalSteps; i++ {
			secret = nextSecret(secret)
		}
		totalSum += secret
	}
	return totalSum, nil
}

func partTwo(data string) (any, error) {
	nums := parseInput(data)
	sequenceTracker := NewSequenceTracker()

	for _, n := range nums {
		buyer := NewBuyer(n)
		sequenceTracker.TrackSequences(buyer)
	}

	return sequenceTracker.MaxBananas(), nil
}

func parseInput(data string) []int {
	var result []int
	for _, line := range strings.Split(data, "\n") {
		v, _ := strconv.Atoi(line)
		result = append(result, v)
	}
	return result
}

func nextSecret(secret int) int {
	a := (secret ^ (secret * 64)) % 16777216
	b := (a ^ (a / 32)) % 16777216
	c := (b ^ (b * 2048)) % 16777216
	return c
}

// Buyer represents a buyer with a secret number
type Buyer struct {
	initialSecret int
	steps         []int
	priceChanges  []int
}

// NewBuyer creates a new Buyer and precomputes its steps and price changes
func NewBuyer(initialSecret int) *Buyer {
	b := &Buyer{
		initialSecret: initialSecret,
		steps:         make([]int, TotalSteps+1),
		priceChanges:  make([]int, TotalSteps),
	}
	b.generateSteps()
	b.calculatePriceChanges()
	return b
}

// generateSteps calculates the steps for the buyer
func (b *Buyer) generateSteps() {
	b.steps[0] = b.initialSecret
	for i := 1; i <= TotalSteps; i++ {
		b.steps[i] = nextSecret(b.steps[i-1])
	}
}

// calculatePriceChanges calculates the price changes (mod 10)
func (b *Buyer) calculatePriceChanges() {
	for i := 1; i <= TotalSteps; i++ {
		b.priceChanges[i-1] = (b.steps[i] % 10) - (b.steps[i-1] % 10)
	}
}

// SequenceTracker tracks sequences and their banana counts
type SequenceTracker struct {
	bananaCounts map[[SequenceLen]int]int
}

// NewSequenceTracker initializes a SequenceTracker
func NewSequenceTracker() *SequenceTracker {
	return &SequenceTracker{
		bananaCounts: make(map[[SequenceLen]int]int),
	}
}

// TrackSequences tracks all valid sequences in a buyer's price changes
func (st *SequenceTracker) TrackSequences(buyer *Buyer) {
	seen := make(map[[SequenceLen]int]bool)
	for i := 0; i <= len(buyer.priceChanges)-SequenceLen; i++ {
		var seq [SequenceLen]int
		copy(seq[:], buyer.priceChanges[i:i+SequenceLen])

		if seen[seq] {
			continue
		}
		seen[seq] = true

		// Add the price at the end of the sequence
		st.bananaCounts[seq] += buyer.steps[i+SequenceLen] % 10
	}
}

// MaxBananas returns the maximum bananas for any sequence
func (st *SequenceTracker) MaxBananas() int {
	maxBananas := 0
	for _, count := range st.bananaCounts {
		if count > maxBananas {
			maxBananas = count
		}
	}
	return maxBananas
}
