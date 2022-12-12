package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id          int
	itemWorries []int
	operator    string
	opRHS       string
	divTest     int
	trueMonkey  int
	falseMonkey int
}

func (m *monkey) popItem() int {
	i := m.itemWorries[0]
	m.itemWorries = m.itemWorries[1:]
	return i
}

func (m *monkey) pushItem(itm int) {
	m.itemWorries = append(m.itemWorries, itm)
}

func parseInput(input string) map[int]*monkey {
	parsed := strings.Split(input, "\n\n")

	var monkies = make(map[int]*monkey)
	for _, m := range parsed {
		pattern := regexp.MustCompile(
			`Monkey ([\d]+):\n` +
				`  Starting items: (.*)\n` +
				`  Operation: new = old (.) ([\w]+)\n` +
				`  Test: divisible by ([\d]+)\n` +
				`    If true: throw to monkey ([\d]+)\n` +
				`    If false: throw to monkey ([\d]+)`,
		)

		groups := pattern.FindStringSubmatch(m)

		id, _ := strconv.Atoi(groups[1])
		var items []int
		for _, itm := range strings.Split(groups[2], ", ") {
			worry, _ := strconv.Atoi(itm)
			items = append(items, worry)
		}
		op := groups[3]
		rhs := groups[4]
		test, _ := strconv.Atoi(groups[5])
		tm, _ := strconv.Atoi(groups[6])
		fm, _ := strconv.Atoi(groups[7])

		monkies[id] = &monkey{id, items, op, rhs, test, tm, fm}
	}

	return monkies
}

func inspectLoop(monkies map[int]*monkey, inspectCounts []int) {
	for mId := 0; mId < len(monkies); mId++ {
		m := monkies[mId]

		for len(m.itemWorries) > 0 {
			itm := m.popItem()
			inspectCounts[mId]++

			var rhs int
			if m.opRHS == "old" {
				rhs = itm
			} else {
				rhs, _ = strconv.Atoi(m.opRHS)
			}

			switch m.operator {
			case "+":
				itm += rhs
			case "*":
				itm *= rhs
			}

			itm /= 3

			var nextMonkyId int
			if itm%m.divTest == 0 {
				nextMonkyId = m.trueMonkey
			} else {
				nextMonkyId = m.falseMonkey
			}

			next := monkies[nextMonkyId]
			next.pushItem(itm)

		}
	}

}

func inspectLoopManagable(monkies map[int]*monkey, inspectCounts []int) {
	mod := 1
	for _, m := range monkies {
		mod *= m.divTest
	}

	for mId := 0; mId < len(monkies); mId++ {
		m := monkies[mId]

		for len(m.itemWorries) > 0 {
			itm := m.popItem() % mod
			inspectCounts[mId]++

			var rhs int
			if m.opRHS == "old" {
				rhs = itm
			} else {
				rhs, _ = strconv.Atoi(m.opRHS)
			}

			switch m.operator {
			case "+":
				itm += rhs
			case "*":
				itm *= rhs
			}

			var nextMonkyId int
			if itm%m.divTest == 0 {
				nextMonkyId = m.trueMonkey
			} else {
				nextMonkyId = m.falseMonkey
			}

			next := monkies[nextMonkyId]
			next.pushItem(itm)

		}
	}

}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	// part 1
	monkies := parseInput(string(file))
	inspectCounts := make([]int, len(monkies))
	for i := 0; i < 20; i++ {
		inspectLoop(monkies, inspectCounts)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))
	fmt.Printf("\n20 round stuff-slinging simian shenanigan level: %d\n", inspectCounts[0]*inspectCounts[1])

	// part 2
	monkies = parseInput(string(file))
	inspectCounts = make([]int, len(monkies))
	for i := 0; i < 10000; i++ {
		inspectLoopManagable(monkies, inspectCounts)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))

	fmt.Printf("10000 round stuff-slinging simian shenanigan level: %d\n", inspectCounts[0]*inspectCounts[1])

}
