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
		fmt.Printf("Monkey %d:", mId)

		for len(m.itemWorries) > 0 {
			itm := m.popItem()
			inspectCounts[mId]++
			fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", itm)

			var rhs int
			if m.opRHS == "old" {
				rhs = itm
			} else {
				rhs, _ = strconv.Atoi(m.opRHS)
			}

			var operation string
			switch m.operator {
			case "+":
				operation = "is multiplied by"
				itm += rhs
			case "*":
				operation = "increases by"
				itm *= rhs
			}
			fmt.Printf("    Worry level %s %d to %d.\n", operation, rhs, itm)

			itm /= 3
			fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", itm)

			var nextMonkyId int
			if itm%m.divTest == 0 {
				fmt.Printf("    Current worry level is divisible by %d.\n", m.divTest)
				nextMonkyId = m.trueMonkey
			} else {
				fmt.Printf("    Current worry level is not divisible by %d.\n", m.divTest)
				nextMonkyId = m.falseMonkey
			}

			fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", itm, nextMonkyId)

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

	monkies := parseInput(string(file))
	inspectCounts := make([]int, len(monkies))
	for i := 0; i < 20; i++ {
		inspectLoop(monkies, inspectCounts)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))

	fmt.Printf("\n20 round stuff-slinging simian shenanigan level: %d\n", inspectCounts[0]*inspectCounts[1])

}
