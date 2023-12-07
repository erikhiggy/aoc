import { readFile, extract, print, sum } from '../utils';

interface Map {
	time: number;
	distance: number;
}

function parsePt1() {
	const input = readFile('day6/input.txt')!;
	const split = input.split('\n');
	const times = extract(split[0]);
	const distances = extract(split[1]);
	const maps: Map[] = [];

	for (let i = 0; i < times.length; i++) {
		maps.push({
			time: Number(times[i]),
			distance: Number(distances[i]),
		});
	}
	return maps;
}

function parsePt2() {
	const input = readFile('day6/input.txt')!;
	const split = input.split('\n');
	const times = extract(split[0]).join('');
	const distances = extract(split[1]).join('');
	const maps: Map[] = [];

	maps.push({
		time: Number(times),
		distance: Number(distances),
	});

	return maps;
}

export function pt1() {
	const maps = parsePt1();
	const combos: number[] = maps.map((map) => {
		const [root1, root2] = quad(-1 * map.time, 1, map.distance + 1);
		const start = Math.ceil(root2);
		const end = Math.floor(root1);
		return end - start + 1;
	});

	return combos.reduce((a, b) => a * b, 1);
}

export function pt2() {
	const maps = parsePt2();
	const combos: number[] = maps.map((map) => {
		const [root1, root2] = quad(-1 * map.time, 1, map.distance + 1);
		const start = Math.ceil(root2);
		const end = Math.floor(root1);
		return end - start + 1;
	});

	return combos.reduce((a, b) => a * b, 1);
}

function quad(b: number, a: number, c: number) {
	const discriminant = b * b - 4 * a * c;
	let root1,
		root2 = 0;

	if (discriminant > 0) {
		root1 = (-b + Math.sqrt(discriminant)) / (2 * a);
		root2 = (-b - Math.sqrt(discriminant)) / (2 * a);
		return [root1, root2];
	} else if (discriminant === 0) {
		root1 = root2 = -b / (2 * a);
		return [root1, root2];
	} else {
		// Don't think we'll need this but include it for completeness.
		let realPart = Number((-b / (2 * a)).toFixed(2));
		let imagPart = Number((Math.sqrt(-discriminant) / (2 * a)).toFixed(2));

		return [realPart, imagPart];
	}
}
