import { readFile, print, sum } from '../utils';

function parseInput() {
	const input = readFile('day9/input.txt')!;
	const lines = input.split('\n');

	return lines.map((line) => line.split(' ').map((l) => Number(l)));
}

function buildTree(sequence: number[]) {
	let seq: number[] = sequence;
	const tree: number[][] = [seq];
	while (!seq.every((s) => s === 0)) {
		const tmp: number[] = [];
		for (let i = 0; i < seq.length - 1; i++) {
			tmp.push(seq[i + 1] - seq[i]);
		}
		tree.push(tmp);
		seq = tmp;
	}

	return tree;
}

function predict(tree: number[][]) {
	// 1. Append a zero to the last sequence.
	tree[tree.length - 1].push(0);
	// 2. For each sequence moving up the tree,
	// Add the sum of the item below, to the last item in the
	// sequence to the end of the sequence.

	let predictedVals: number[] = [];
	for (let i = tree.length - 2; i >= 0; i--) {
		let itemBelow = tree[i + 1][tree[i + 1].length - 1];
		let itemToLeft = tree[i][tree[i].length - 1];
		tree[i].push(itemBelow + itemToLeft);
		predictedVals.push(itemBelow + itemToLeft);
	}

	return predictedVals;
}

function predictPt2(tree: number[][]) {
	// 1. Append a zero to the front of the last sequence.
	tree[tree.length - 1].unshift(0);
	// 2. For each sequence moving up the tree,
	// Add the sum of the item below, to the last item in the
	// sequence to the end of the sequence.

	let predictedVals: number[] = [];
	for (let i = tree.length - 2; i >= 0; i--) {
		let itemBelow = tree[i + 1][0];
		let itemToLeft = tree[i][0];
		tree[i].unshift(itemToLeft - itemBelow);
		predictedVals.push(itemToLeft - itemBelow);
	}

	return predictedVals;
}

export function pt1() {
	const input = parseInput();
	let sum = 0;
	for (let _in of input) {
		let tree = buildTree(_in);
		sum += predict(tree)[predict(tree).length - 1];
	}
	return sum;
}

export function pt2() {
	const input = parseInput();
	let sum = 0;
	for (let _in of input) {
		let tree = buildTree(_in);
		sum += predictPt2(tree)[predictPt2(tree).length - 1];
	}
	return sum;
}
