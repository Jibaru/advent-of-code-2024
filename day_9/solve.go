package day9

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const GapID = -1

func Solve(part int, isTest bool) (any, error) {
	f := "day_9/input.txt"
	if isTest {
		f = "day_9/input-test.txt"
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
	diskMap := newDiskMapFromCompressed(data)

	gaps := diskMap.GapsIndexesQueue()
	noGaps := diskMap.NoGapsIndexesQueue()

	for diskMap.HasGapsBetweenBlocks() {
		gapIdx := gaps[0]
		gaps = gaps[1:]

		noGapIdx := noGaps[0]
		noGaps = noGaps[1:]

		diskMap.Swap(gapIdx, noGapIdx)
	}

	return diskMap.Checksum(), nil
}

func partTwo(data string) (any, error) {
	diskMap := newBlockDiskMapFromCompressed(data)

	noGaps := diskMap.NoGapsIndexesQueue()

	for len(noGaps) > 0 {
		noGapIdx := noGaps[0]
		noGaps = noGaps[1:]

		gaps := diskMap.GapsIndexesQueue()

		for len(gaps) > 0 {
			gapIdx := gaps[0]
			gaps = gaps[1:]

			if noGapIdx < gapIdx {
				break
			}

			if !diskMap.CanReplace(noGapIdx, gapIdx) {
				continue
			}

			diskMap.Replace(noGapIdx, gapIdx)
			noGaps = diskMap.NoGapsIndexesQueue()
			break
		}
	}

	return diskMap.Checksum(), nil
}

type DiskMap struct {
	values []int
}

func newDiskMapFromCompressed(data string) *DiskMap {
	diskMap := []int{}
	id := 0
	isBlock := true

	for _, c := range strings.Split(data, "") {
		total, _ := strconv.ParseInt(c, 10, 64)
		char := GapID // this is "."
		if isBlock {
			char = id
			id++
		}
		for i := 1; i <= int(total); i++ {
			diskMap = append(diskMap, char)
		}
		isBlock = !isBlock
	}

	return &DiskMap{
		values: diskMap,
	}
}

func (d *DiskMap) HasGapsBetweenBlocks() bool {
	canNotBeGapNow := false
	for i := len(d.values) - 1; i >= 0; i-- {
		if d.values[i] == GapID && canNotBeGapNow {
			return true
		}

		if d.values[i] != GapID {
			canNotBeGapNow = true
		}
	}

	return false
}

func (d DiskMap) Swap(i, j int) {
	aux := d.values[j]
	d.values[j] = d.values[i]
	d.values[i] = aux
}

func (d DiskMap) GapsIndexesQueue() []int {
	idxs := []int{}

	for i, v := range d.values {
		if v == GapID {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func (d DiskMap) NoGapsIndexesQueue() []int {
	idxs := []int{}

	for i := len(d.values) - 1; i >= 0; i-- {
		if d.values[i] != GapID {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func (d *DiskMap) Checksum() int {
	checksum := 0
	for i, v := range d.values {
		if v != GapID {
			checksum += (i * v)
		}
	}

	return checksum
}

type BlockDiskMap struct {
	values []Block
}

type Block struct {
	id     int
	spaces int
}

func (b Block) IsGap() bool {
	return b.id == GapID
}

func newBlockDiskMapFromCompressed(data string) *BlockDiskMap {
	diskMap := []Block{}
	id := 0
	isBlock := true

	for _, c := range strings.Split(data, "") {
		total, _ := strconv.ParseInt(c, 10, 64)
		char := GapID // this is "."
		if isBlock {
			char = id
			id++
		}
		diskMap = append(diskMap, Block{
			id:     char,
			spaces: int(total),
		})
		isBlock = !isBlock
	}

	return &BlockDiskMap{
		values: diskMap,
	}
}

func (b *BlockDiskMap) CanReplace(validIdx, gapIdx int) bool {
	gap := b.values[gapIdx]
	valid := b.values[validIdx]

	return gap.spaces >= valid.spaces
}

func (b *BlockDiskMap) Replace(validIdx, gapIdx int) {
	gap := b.values[gapIdx]
	valid := b.values[validIdx]

	if gap.spaces >= valid.spaces {
		gap.spaces -= valid.spaces
	}

	b.values[gapIdx] = valid
	b.values[validIdx] = Block{id: GapID, spaces: valid.spaces}

	if gap.spaces > 0 {
		b.values = slices.Insert(b.values, gapIdx+1, gap)
	}
}

func (d *BlockDiskMap) GapsIndexesQueue() []int {
	idxs := []int{}

	for i, v := range d.values {
		if v.IsGap() {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func (d *BlockDiskMap) NoGapsIndexesQueue() []int {
	idxs := []int{}

	for i := len(d.values) - 1; i >= 0; i-- {
		if !d.values[i].IsGap() {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func (d *BlockDiskMap) String() string {
	s := ""
	for _, v := range d.values {
		c := "."
		if !v.IsGap() {
			c = fmt.Sprintf("%v", v.id)
		}

		for i := 0; i < v.spaces; i++ {
			s += c
		}
	}

	return s
}

func (d *BlockDiskMap) Checksum() int {
	checksum := 0
	idx := 0
	for _, v := range d.values {
		if v.IsGap() {
			idx += v.spaces
			continue
		}

		for i := 0; i < v.spaces; i++ {
			checksum += (idx * v.id)
			idx++
		}
	}

	return checksum
}
