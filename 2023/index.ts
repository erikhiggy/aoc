import { pt1 as day1pt1, pt2 as day1Pt2 } from './day1';
import { pt1 as day2pt1, pt2 as day2pt2 } from './day2';
import { pt1 as day3pt1, pt2 as day3pt2 } from './day3';
import { pt1 as day4pt1, pt2 as day4pt2 } from './day4';

const solutions = {
  day1Pt1: `Day 1 Part 1: ${day1pt1()}`,
  day1Pt2: `Day 1 Part 2: ${day1Pt2()}`,
  day2Pt1: `Day 2 Part 1: ${day2pt1()}`,
  day2pt2: `Day 2 Part 2: ${day2pt2()}`,
  day3pt1: `Day 3 Part 1: ${day3pt1()}`,
  day3pt2: `Day 3 Part 2: ${day3pt2()}`,
  day4pt1: `Day 4 Part 1: ${day4pt1()}`,
  day4pt2: `Day 4 Part 2: ${day4pt2()}`,
};

for (const value of Object.values(solutions)) {
  console.log(value);
}
