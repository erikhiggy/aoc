import { pt1 as day1pt1, pt2 as day1Pt2 } from './day1';
import { pt1 as day2pt1, pt2 as day2pt2 } from './day2';
import { pt1 as day3pt1, pt2 as day3pt2 } from './day3';
import { pt1 as day4pt1, pt2 as day4pt2 } from './day4';
import { pt1 as day5pt1, pt2 as day5pt2 } from './day5';
import { pt1 as day6pt1, pt2 as day6pt2 } from './day6';
import { pt1 as day7pt1, pt2 as day7pt2 } from './day7';
import { pt1 as day8pt1, pt2 as day8pt2 } from './day8';
import { pt1 as day9pt1, pt2 as day9pt2 } from './day9';
import { pt1 as day10pt1, pt2 as day10pt2 } from './day10';

const solutions = {
  day1Pt1: `Day 1 Part 1: ${day1pt1()}`,
  day1Pt2: `Day 1 Part 2: ${day1Pt2()}`,
  day2Pt1: `Day 2 Part 1: ${day2pt1()}`,
  day2pt2: `Day 2 Part 2: ${day2pt2()}`,
  day3pt1: `Day 3 Part 1: ${day3pt1()}`,
  day3pt2: `Day 3 Part 2: ${day3pt2()}`,
  day4pt1: `Day 4 Part 1: ${day4pt1()}`,
  day4pt2: `Day 4 Part 2: ${day4pt2()}`,
  day5pt1: `Day 5 Part 1: ${day5pt1()}`,
  // day5pt2: `Day 5 Part 2: ${day5pt2()}`, Uncomment for hell
  day6pt1: `Day 6 Part 1: ${day6pt1()}`,
  day6pt2: `Day 6 Part 2: ${day6pt2()}`,
  day7pt1: `Day 7 Part 1: ${day7pt1()}`,
  day7pt2: `Day 7 Part 2: ${day7pt2()}`,
  day8pt1: `Day 8 Part 1: ${day8pt1()}`,
  day8pt2: `Day 8 Part 2: ${day8pt2()}`,
  day9pt1: `Day 9 Part 1: ${day9pt1()}`,
  day9pt2: `Day 9 Part 2: ${day9pt2()}`,
  day10pt1: `Day 10 Part 1: ${day10pt1()}`,
  // day10pt2: `Day 10 Part 2: ${day10pt2()}`, // Come back to this one.
};

for (const value of Object.values(solutions)) {
  console.log(value);
}
