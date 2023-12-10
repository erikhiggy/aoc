import { readFile, print } from '../utils';

function parseInput() {
	const grid = readFile('day10/input.txt')!.split('\n');
	return grid;
}

function traverse(grid: string[]) {
	let coords: number[] = [];
	for (let i = 0; i < grid.length; i++) {
		for (let j = 0; j < grid[i].length; j++) {
			if (grid[i].charAt(j) === 'S') {
				coords = [i, j];
				break;
			}
		}
	}

	// Look around S
	let [row, col] = coords;
	const top = grid[row - 1][col];
	const bot = grid[row + 1][col];
	const left = grid[row][col - 1];
	const right = grid[row][col + 1];
	print([top, right, bot, left]);

	let queue = [`${row},${col}`];
	const visited = new Set<string>(); // holds the coords
	const result: string[] = [];
	while (queue.length > 0) {
		const curr = queue.shift();
		if (!curr) continue;
		result.push(curr);
		visited.add(curr);

         [row, col] = curr.split(',').map(c => Number(c));
         const item = grid[row][col];
		if (item === '')
	}
}

export function pt1() {
	const grid = parseInput();
	traverse(grid);
	return -1;
}
