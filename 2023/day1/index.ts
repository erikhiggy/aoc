import { readFile } from '../utils';

export function pt1() {
	const input = readFile('day1/input.txt');
	const split = input?.split('\n');
	let sum = 0;
	split?.forEach((item) => {
		let lineNumString: string = '';
		for (const letter of item) {
			const asNum = parseInt(letter, 10);
			if (!isNaN(asNum)) {
				lineNumString += letter;
			}
		}
		sum += parseInt(
			lineNumString[0] + lineNumString[lineNumString.length - 1],
			10
		);
	});

	return sum;
}

const numberStringsMap: { [key: string]: string } = {
	one: '1',
	two: '2',
	three: '3',
	four: '4',
	five: '5',
	six: '6',
	seven: '7',
	eight: '8',
	nine: '9',
	'1': '1',
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
};

const possible = [
	'one',
	'two',
	'three',
	'four',
	'five',
	'six',
	'seven',
	'eight',
	'nine',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
];

export function pt2() {
	const input = readFile('day1/input1.txt');
	const split = input?.split('\n');
	let sum = 0;
	split?.forEach((line) => {
		const map = new Map<string, { first: number; last: number }>();
		let numberString = '';
		possible.forEach((p) => {
			map.set(p, { first: line.indexOf(p), last: line.lastIndexOf(p) });
		});
		const filtered = [...map].filter(
			(m) => m[1].first !== -1 || m[1].last !== -1
		);
		const firstNum = filtered.sort((a, b) => a[1].first - b[1].first)[0];
		const lastNum = filtered.sort((a, b) => b[1].last - a[1].last)[0];
		numberString = numberStringsMap[firstNum[0]] + numberStringsMap[lastNum[0]];
		sum += parseInt(numberString, 10);
	});

	return sum;
}
