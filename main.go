package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Operation int

const (
	Equal Operation = iota
	Insert
	Delete
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

type EditOperation struct {
	Operation Operation
	Line      string
}

func genEditSequence(trace [][]int, old, new []string, offset int) []EditOperation {
	editSeq := []EditOperation{}
	x, y := len(old), len(new)

	for d := len(trace) - 1; d > 0; d-- {
		v := trace[d]
		k := x - y
		kIdx := k + offset

		var prevK int
		if k == -d || (k != d && v[kIdx-1] < v[kIdx+1]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX := v[prevK+offset]
		prevY := prevX - prevK

		for x > prevX && y > prevY {
			editSeq = append(editSeq, EditOperation{Operation: Equal, Line: old[x-1]})
			x--
			y--
		}

		if x > prevX {
			editSeq = append(editSeq, EditOperation{Operation: Delete, Line: old[x-1]})
		} else if y > prevY {
			editSeq = append(editSeq, EditOperation{Operation: Insert, Line: new[y-1]})
		}

		x, y = prevX, prevY
	}

	for x > 0 && y > 0 && old[x-1] == new[y-1] {
		editSeq = append(editSeq, EditOperation{Operation: Equal, Line: old[x-1]})
		x--
		y--
	}

	slices.Reverse(editSeq)
	return editSeq
}

func MyersAlgo(old, new []string) []EditOperation {
	n, m := len(old), len(new)
	offset := n + m
	v := make([]int, 2*offset+1)
	var trace [][]int

	for d := 0; d <= n+m; d++ {

		vCopy := make([]int, len(v))
		copy(vCopy, v)
		trace = append(trace, vCopy)

		for k := -d; k <= d; k += 2 {
			kIdx := k + offset
			var x int
			if k == -d || (k != d && v[kIdx-1] < v[kIdx+1]) {
				x = v[kIdx+1]
			} else {
				x = v[kIdx-1] + 1
			}
			y := x - k

			for x < n && y < m && old[x] == new[y] {
				x++
				y++
			}
			v[kIdx] = x

			if x >= n && y >= m {
				return genEditSequence(trace, old, new, offset)
			}
		}
	}
	return nil
}

func ComputeFileDiff(oldContent, newContent string) string {

	oldLines := strings.Split(strings.ReplaceAll(oldContent, "\r\n", "\n"), "\n")
	newLines := strings.Split(strings.ReplaceAll(newContent, "\r\n", "\n"), "\n")

	editSeq := MyersAlgo(oldLines, newLines)

	var sb strings.Builder
	for _, op := range editSeq {
		switch op.Operation {
		case Insert:

			sb.WriteString(Green + "+ " + op.Line + Reset + "\n")
		case Delete:

			sb.WriteString(Red + "- " + op.Line + Reset + "\n")
		case Equal:

			sb.WriteString("  " + op.Line + "\n")
		}
	}
	return sb.String()
}

func main() {
	oldFileName := "old.txt"
	newFileName := "new.txt"
	oldContent, err := os.ReadFile(oldFileName)
	if err != nil {
		fmt.Printf("Error while reading file: %s", err.Error())
	}

	newContent, err := os.ReadFile(newFileName)
	if err != nil {
		fmt.Printf("Error while reading file: %s", err.Error())
	}

	diffContent := ComputeFileDiff(string(oldContent), string(newContent))
	print(diffContent)
}
