package main

import (
	"fmt"
	"strings"
)

type point struct {
	x, y     int
	antenna  *rune
	antinode bool
}

func (p point) String() string {
	return fmt.Sprintf("{ (%d, %d) antenna: '%s', antinode: %t }\n", p.x, p.y, runePtrToString(p.antenna), p.antinode)
}

func runePtrToString(p *rune) string {
	if p != nil {
		return string(*p)
	}
	return "(nil)"
}

const sample = `
..........W......8..............................4.
............W...................F............1.L..
......o.................................A.........
................Z.........................A.......
................u.................................
.........8.......g...........................F....
.............2..8.......F......................1..
g....G............................................
.o..2.g...........Z..W.......................4b.1.
.................v...Z....c.....B...1.......f...b.
...uG.c............O.............Z................
..G..........c8....................5..4...........
...c.....G.g..........x..........B5............b..
.S.....o..v.......................................
............................BV....................
..........u............x.............B......0.....
...3....C..........I............V6..............f.
........S2......C.......................5.........
................a....v............................
..y...2..............i.......k.................4..
..........I........Yv......5...........f..........
.....3....o...................x...................
..........3.........S.........k...................
...y.......C.......d.I......X......fV.............
.....S............................d...............
.....O........I.......................iV.....A....
.y.................x.............k.ai0...F........
.......Y..............................a...........
........z.........................0...............
..O.3..............................0.......a......
.....z............................X...............
......z................................l..........
......L.......U........d.....X....................
......7z..........d...............................
O........7.....................K..................
.....................X...6........k...............
...L........Y...................s.................
.7...D................................l.9.........
..D...........................w..................i
......D.............6............s................
.................w..........6.9............s......
..............................s...................
........7.........................................
L...Y..........................U..................
.....................9..........l.........K.......
.......................9..............U...........
...D..............................................
.................................................K
..........y......................U.............l..
..................w...................K...........
`

func parseInput(in string) []string {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	return lines
}

func main() {

	input := parseInput(sample)
	// Part 1
	p, a := makePoints(input, false)
	antinodeCount := countAntinodes(p, a)
	fmt.Printf("Part 1: %d\n", antinodeCount)

	// Part 2
	p, a = makePoints(input, true)
	antinodeCount = countAntinodesPt2(p, a)
	fmt.Printf("Part 2: %d\n", antinodeCount)
}

func countAntinodes(p [][]*point, a map[rune][]*point) int {
	antinodeCount := 0
	for _, v := range a {
		if len(v) < 2 {
			// no antinodes
			continue
		}

		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				// find distance between two points
				dx, dy := v[j].x-v[i].x, v[j].y-v[i].y

				x2 := v[j].x + dx
				y2 := v[j].y + dy
				if x2 >= 0 && x2 <= len(p[0])-1 && y2 >= 0 && y2 <= len(p)-1 {
					if !p[y2][x2].antinode {
						p[y2][x2].antinode = true
						antinodeCount++
					}
				}

				x3 := v[i].x - dx
				y3 := v[i].y - dy
				if x3 >= 0 && x3 <= len(p[0])-1 && y3 >= 0 && y3 <= len(p)-1 {
					if !p[y3][x3].antinode {
						p[y3][x3].antinode = true
						antinodeCount++
					}
				}
			}
		}
	}
	return antinodeCount
}

func countAntinodesPt2(p [][]*point, a map[rune][]*point) int {
	antinodeCount := 0
	for _, v := range a {
		if len(v) < 2 {
			// no antinodes
			continue
		}
		antinodeCount += len(v)

		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				// find distance between two points
				dx, dy := v[j].x-v[i].x, v[j].y-v[i].y

				x2 := v[j].x + dx
				y2 := v[j].y + dy
				for x2 >= 0 && x2 <= len(p[0])-1 && y2 >= 0 && y2 <= len(p)-1 {
					if !p[y2][x2].antinode {
						p[y2][x2].antinode = true
						antinodeCount++
					}
					x2 = x2 + dx
					y2 = y2 + dy
				}

				x3 := v[i].x - dx
				y3 := v[i].y - dy
				for x3 >= 0 && x3 <= len(p[0])-1 && y3 >= 0 && y3 <= len(p)-1 {
					if !p[y3][x3].antinode {
						p[y3][x3].antinode = true
						antinodeCount++
					}
					x3 = x3 - dx
					y3 = y3 - dy
				}
			}
		}
	}
	return antinodeCount
}

func makePoints(in []string, antennaIsAntinode bool) ([][]*point, map[rune][]*point) {
	points := make([][]*point, len(in))
	antennas := map[rune][]*point{}
	for r, line := range in {
		for c, char := range line {
			p := &point{
				x: c,
				y: r,
			}
			if char != '.' {
				p.antenna = &char
				if antennaIsAntinode {
					p.antinode = true
				}
			}
			if p.antenna != nil {
				antennas[*p.antenna] = append(antennas[*p.antenna], p)
			}
			points[r] = append(points[r], p)
		}
	}
	return points, antennas
}
