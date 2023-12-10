import { readFile, print, sum } from '../utils';

function parseInput() {
	const grid = readFile('day10/input.txt')!.split('\n');
	return grid;
}

const StartOpts: { [key: string]: string[] } = {
	top: ['|', '7', 'F'],
	right: ['-', 'J', '7'],
	left: ['-', 'F', 'L'],
	bottom: ['|', 'J', 'L'],
};

const VerticalPipeOpts: { [key: string]: string[] } = {
	top: ['|', '7', 'F'],
	bottom: ['|', 'J', 'L'],
};

const HorizontalPipeOpts: { [key: string]: string[] } = {
	right: ['-', 'J', '7'],
	left: ['-', 'L', 'F'],
};

const FPipeOpts: { [key: string]: string[] } = {
	right: ['J', '7', '-'],
	bottom: ['|', 'L', 'J'],
};

const JPipeOpts: { [key: string]: string[] } = {
	top: ['|', 'F', '7'],
	left: ['-', 'F', 'L'],
};

const LPipeOpts: { [key: string]: string[] } = {
	top: ['|', '7', 'F'],
	right: ['-', 'J', '7'],
};

const _7PipeOpts: { [key: string]: string[] } = {
	left: ['-', 'L', 'F'],
	bottom: ['|', 'L', 'J'],
};

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

	let queue = [`${row},${col}`];
	const visited = new Set<string>(); // holds the coords
	const result: string[] = [];
	while (queue.length > 0) {
		const curr = queue.shift();
		if (!curr) continue;
		[row, col] = curr.split(',').map((c) => Number(c));
		const item = grid[row][col];
		result.push(item);
		visited.add(curr);

		const top = row !== 0 ? grid[row - 1][col] : '';
		const bot = row !== grid.length - 1 ? grid[row + 1][col] : '';
		const left = col !== 0 ? grid[row][col - 1] : '';
		const right = col !== grid[row].length - 1 ? grid[row][col + 1] : '';

		let coordsStr = '';

		if (item === 'S') {
			if (StartOpts['top'].includes(top)) {
				coordsStr = `${row - 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			} else if (StartOpts['right'].includes(right)) {
				coordsStr = `${row},${col + 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			} else if (StartOpts['bottom'].includes(bot)) {
				coordsStr = `${row + 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			} else if (StartOpts['left'].includes(left)) {
				coordsStr = `${row},${col - 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === '|') {
			if (VerticalPipeOpts['top'].includes(top)) {
				coordsStr = `${row - 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (VerticalPipeOpts['bottom'].includes(bot)) {
				coordsStr = `${row + 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === '-') {
			if (HorizontalPipeOpts['right'].includes(right)) {
				coordsStr = `${row},${col + 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (HorizontalPipeOpts['left'].includes(left)) {
				coordsStr = `${row},${col - 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === 'F') {
			if (FPipeOpts['right'].includes(right)) {
				coordsStr = `${row},${col + 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (FPipeOpts['bottom'].includes(bot)) {
				coordsStr = `${row + 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === 'J') {
			if (JPipeOpts['left'].includes(left)) {
				coordsStr = `${row},${col - 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (JPipeOpts['top'].includes(top)) {
				coordsStr = `${row - 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === 'L') {
			if (LPipeOpts['top'].includes(top)) {
				coordsStr = `${row - 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (LPipeOpts['right'].includes(right)) {
				coordsStr = `${row},${col + 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		} else if (item === '7') {
			if (_7PipeOpts['left'].includes(left)) {
				coordsStr = `${row},${col - 1}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
			if (_7PipeOpts['bottom'].includes(bot)) {
				coordsStr = `${row + 1},${col}`;
				if (!visited.has(coordsStr)) {
					queue.push(coordsStr);
				}
			}
		}
	}

	return result;
}

export function pt1() {
	const grid = parseInput();
	const res = traverse(grid);
	return res.length / 2;
}
