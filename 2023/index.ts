import { pt1, pt2 } from './day1';

const solutions = {
	day1Pt1: `Day 1 Part 1: ${pt1()}`,
	day1Pt2: `Day 1 Part 2: ${pt2()}`,
};

for (const value of Object.values(solutions)) {
	console.log(value);
}
