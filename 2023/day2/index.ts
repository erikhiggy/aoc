import { readFile } from '../utils';

let cubes: { [key: string]: number } = {
	red: 0,
	green: 0,
	blue: 0,
};

export function pt1() {
	const input = readFile('day2/input.txt');
	const games = input?.split('\n');
	const res: boolean[] = [];
	games?.forEach((game) => {
		const colorSequences: { [key: string]: number }[] = [];
		const sequences = game.split(':')[1].split(';');
		sequences.forEach((seq) => {
			const nums = seq.split(',').map((n) => n.trim());
			nums.forEach((n) => {
				const [number, color] = n.split(' ');
				cubes[color] = Number(number);
			});
			colorSequences.push(cubes);
			cubes = resetCubes();
		});
		let isValid = false;
		// console.log(`Game ${idx} sequences`, colorSequences);
		for (const seq of colorSequences) {
			if (seq.red > 12 || seq.green > 13 || seq.blue > 14) {
				isValid = false;
				break;
			}
			isValid = true;
		}
		res.push(isValid);
	});

	return res
		.map((r, idx) => {
			if (r === true) {
				return idx + 1;
			}
			return 0;
		})
		.reduce((a, b) => a + b, 0);
}

export function pt2() {
	const input = readFile('day2/input1.txt');
	const games = input?.split('\n');
	const powers: number[] = [];
	games?.forEach((game) => {
		const colorSequences: { [key: string]: number }[] = [];
		const sequences = game.split(':')[1].split(';');
		sequences.forEach((seq) => {
			const nums = seq.split(',').map((n) => n.trim());
			nums.forEach((n) => {
				const [number, color] = n.split(' ');
				cubes[color] = Math.max(Number(number), cubes[color]);
			});
			colorSequences.push(cubes);
		});
		cubes = resetCubes();
		let power = 1;
		for (const value of Object.values(colorSequences[0])) {
			power *= value;
		}
		powers.push(power);
	});
	return powers.reduce((a, b) => a + b, 0);
}

function resetCubes() {
	return {
		red: 0,
		blue: 0,
		green: 0,
	};
}
