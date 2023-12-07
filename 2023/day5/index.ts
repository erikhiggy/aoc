import { extract, readFile, print } from '../utils';

export function pt1() {
	const input = readFile('day5/input.txt')!;
	const sections = input.split('\n\n');
	const seeds = extract(sections[0]);
	const maps: string[][][] = [];

	for (let i = 1; i < sections.length; i++) {
		const s = sections[i];

		const split = s.split('\n');
		const sec = [];
		for (let j = 1; j < split.length; j++) {
			const line = split[j];
			const extracted = extract(line)!;
			sec.push(extracted);
		}
		maps.push(sec);
	}

	function traverse(idx: number, val: number) {
		// base case
		if (idx >= maps.length) {
			return val;
		}

		for (let [dStart, sStart, len] of maps[idx]) {
			const dstStart = Number(dStart);
			const srcStart = Number(sStart);
			const rLen = Number(len);
			if (val >= srcStart && val < srcStart + rLen) {
				return traverse(idx + 1, dstStart + val - srcStart);
			}
		}
		return traverse(idx + 1, val);
	}

	let loc = -1;
	let max = Number.POSITIVE_INFINITY;
	for (const seed of seeds) {
		loc = traverse(0, Number(seed));
		if (loc < max) {
			max = loc;
		}
	}
	return max;
}

export function pt2() {
	const input = readFile('day5/input.txt')!;
	const sections = input.split('\n\n');
	const seeds = extract(sections[0]);
	const maps: Map[][] = [];

	for (let i = 1; i < sections.length; i++) {
		const s = sections[i];

		const split = s.split('\n');
		const sec = [];
		for (let j = 1; j < split.length; j++) {
			const line = split[j];
			const extracted = extract(line)!;
			const tmpMap: Map = {
				dst: Number(extracted[0]),
				src: Number(extracted[1]),
				len: Number(extracted[2]),
			};
			sec.push(tmpMap);
		}
		maps.push(sec);
	}

	let loc = Number.POSITIVE_INFINITY;
	for (let j = 0; j < seeds.length; j = j + 2) {
		const start = Number(seeds[j]);
		const count = Number(seeds[j + 1]);
		print(`Working on seed ${start}`);
		for (let i = start; i < start + count; i++) {
			loc = Math.min(loc, applyMaps(i, maps));
		}
	}
	return loc - 1; // -__-
}

interface Map {
	dst: number;
	src: number;
	len: number;
}

function apply(seed: number, ranges: Map[]) {
	const mapping = ranges.find((r) => seed >= r.src && seed <= r.src + r.len);
	if (!mapping) {
		return seed;
	}

	const offset = seed - mapping.src;
	return mapping.dst + offset;
}

function applyMaps(seed: number, maps: Map[][]) {
	return maps.reduce((seed, ranges) => apply(seed, ranges), seed);
}
