import { readFile } from '../utils';

export function pt1() {
	const cards = readFile('day4/input.txt')!.split('\n');
	let res: string[][] = [];
	cards.forEach((card, i) => {
		const numsWithSep = card.split(':')[1];
		let [winningNums, myNums] = numsWithSep.split('|');
		winningNums = winningNums.trim();
		myNums = myNums.trim();
		const winnersAsList = new Set<string>(
			winningNums.split(' ').filter((s) => s !== '')
		);
		const myNumsAsList = new Set<string>(
			myNums.split(' ').filter((s) => s !== '')
		);
		const tmp: string[] = [];
		console.log(myNumsAsList);
		myNumsAsList.forEach((num) => {
			if (winnersAsList.has(num)) {
				tmp.push(num);
			}
		});
		res.push(tmp);
	});
	const powers = res.map((r) => {
		return r.length > 0 ? 2 ** (r.length - 1) : 0;
	});
	return powers.reduce((a, b) => {
		return a + b;
	}, 0);
}
