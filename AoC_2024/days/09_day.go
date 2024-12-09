package days

import (
	"os"
	"strconv"
	"strings"
)

type Block struct {
	pos  int
	size int
}

type Day9 struct {
	disk       []int
	freeIdx    []int
	diskBlocks map[int]Block
	blanks     []Block
}

func (d *Day9) Part1() (int, error) {
	res := d.disk
	for _, i := range d.freeIdx {
		for len(res) >= 1 && res[len(res)-1] == -1 {
			res = res[:len(res)-1]
		}
		if len(res) <= i {
			break
		}
		res[i] = res[len(res)-1]
		// Pop
		res = res[:len(res)-1]
	}
	return d.checksumPart1(res), nil
}

func (d *Day9) Part2() (int, error) {
	fid := len(d.diskBlocks)
	for fid > 0 {
		fid--
		for i, free := range d.blanks {
			// Don't check if on the right of block to move
			if free.pos >= d.diskBlocks[fid].pos {
				d.blanks = d.blanks[:i]
				break
			}

			if free.size >= d.diskBlocks[fid].size {
				// Move block
				if b, ok := d.diskBlocks[fid]; ok {
					b.pos = free.pos
					d.diskBlocks[fid] = b
				}

				if free.size == d.diskBlocks[fid].size {
					// Remove blank element
					d.blanks = append(d.blanks[:i], d.blanks[i+1:]...)
				} else {
					// Update blank
					d.blanks[i].size = free.size - d.diskBlocks[fid].size
					d.blanks[i].pos += d.diskBlocks[fid].size
				}

				break
			}
		}
	}
	return d.checksumPart2(d.diskBlocks), nil
}

func (d *Day9) Parse() error {
	data, err := os.ReadFile("./data/09_day.txt")
	if err != nil {
		return err
	}

	line := strings.TrimSpace(string(data))
	d.disk = make([]int, 0)
	d.freeIdx = make([]int, 0)
	idx := 0

	d.diskBlocks = make(map[int]Block)
	d.blanks = make([]Block, 0)
	for i, c := range line {
		nb, _ := strconv.Atoi(string(c))

		// Part 1
		for j := range nb {
			if i%2 == 0 {
				fileId := i / 2
				d.disk = append(d.disk, fileId)
			} else {
				// Free space
				d.disk = append(d.disk, -1)
				d.freeIdx = append(d.freeIdx, idx+j)
			}
		}

		// Part 2
		if i%2 == 0 {
			d.diskBlocks[i/2] = Block{size: nb, pos: idx}
		} else {
			d.blanks = append(d.blanks, Block{size: nb, pos: idx})
		}
		idx += nb
	}

	return nil
}

func (d *Day9) checksumPart1(disk []int) (tot int) {
	for i, nb := range disk {
		tot += i * nb
	}
	return
}

func (d *Day9) checksumPart2(disk map[int]Block) (tot int) {
	for fid, b := range disk {
		for j := b.pos; j < b.pos+b.size; j++ {
			tot += j * fid
		}
	}
	return
}
