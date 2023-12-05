import { readFile, writeFile } from '../utils';

export function pt1() {
	const input = readFile('day3/input.txt');
	const rows = input?.split('\n') || [];
	let res: number = 0;
	let possiblePartNumber = '';
	let validPartNumber: boolean[] = [];
	for (let i = 0; i < rows.length; i++) {
		for (let j = 0; j < rows[i].length; j++) {
			const item = rows[i][j].trim();
			const endOfRow = j === rows[i].length - 1;
			// Just move on if we encounter a symbol
			// or a period with no possiblePartNumber contents.
			if (isNaN(Number(item)) && possiblePartNumber === '') {
				continue;
			}
			// If the item is not a period and it's not a symbol
			// add the digit to the possiblePartNumber string
			// and check all around it to see if it's a part number.
			if (!isNaN(Number(item))) {
				possiblePartNumber += item;
				validPartNumber.push(isPartNumber(rows, i, j));
			}

			if ((isNaN(Number(item)) || endOfRow) && possiblePartNumber !== '') {
				// If we hit a period, AND we have contents in the possiblePartNumber
				// string, check if one of the entries in the validPartNumber array
				// is true, and if so, add the possibleNumberString to the result list.
				// Regardless, we want to clear both the string and array.
				if (validPartNumber.some((v) => v === true)) {
					res += Number(possiblePartNumber);
				}
				possiblePartNumber = '';
				validPartNumber = [];
			}
		}
		// Always clear at the end of the line.
		possiblePartNumber = '';
		validPartNumber = [];
	}
	return res;
}

export function pt2() {
	const input = readFile('day3/input.txt');
	const grid = input?.split(/\n/g).map((line) => line.split(''))!;
	let gearNums: { [key: string]: number[] } = {
		'': [],
	};

	for (let i = 0; i < grid.length; i++) {
		let currNumber = '',
			checking = false,
			gearLoc = '';
		for (let j = 0; j < grid[i].length; j++) {
			const char = grid[i][j];
			const endOfRow = grid[i].length - 1;
			// if we are at a number, but not checking yet, start to check
			if (isDigit(char) && !checking) {
				checking = true;
				currNumber = '';
				gearLoc = '';
			}

			// If we are at the end of a row, or we encounter a non-digit and we are checking, stop checking
			// and add to the sum if needed
			if ((j === endOfRow || !isDigit(char)) && checking) {
				if (gearLoc !== '') {
					gearNums[gearLoc].push(
						parseInt(currNumber + (isDigit(char) ? char : ''), 10)
					);
				}
				checking = false;
			}

			// If we are checking, add the current char to the number and check for '*' around it
			if (checking) {
				currNumber += char;

				for (let x = -1; x <= 1; x++) {
					for (let y = -1; y <= 1; y++) {
						if (x === 0 && y === 0) {
							continue;
						}

						if (
							i + x < 0 ||
							i + x >= grid.length ||
							j + y < 0 ||
							j + y >= grid[i].length
						) {
							continue;
						}

						const token = grid[i + x][j + y];
						if (isGear(token)) {
							gearLoc = `${j + y},${i + x}`;
							if (!gearNums[gearLoc]) gearNums[gearLoc] = [];
						}
					}
				}
			}
		}
	}

	return Object.values(gearNums).reduce((sum, arr) => {
		if (arr.length === 2) {
			sum += arr[0] * arr[1];
		}
		return sum;
	}, 0);
}

function isDigit(str: string) {
	return str.match(/[0-9]/);
}

function isGear(str: string) {
	return str === '*';
}

function isSymbol(str: string) {
	return !isPeriod(str) && isNaN(Number(str));
}

function isPeriod(str: string) {
	return str === '.';
}

function isPartNumber(arr: string[], row: number, col: number) {
	if (row === 0) {
		if (col === 0) {
			return (
				isSymbol(bottom(arr, row, col)) ||
				isSymbol(right(arr, row, col)) ||
				isSymbol(bottomRight(arr, row, col))
			);
		} else if (col === arr.length - 1) {
			return (
				isSymbol(bottom(arr, row, col)) ||
				isSymbol(left(arr, row, col)) ||
				isSymbol(bottomLeft(arr, row, col))
			);
		}
		return (
			isSymbol(bottom(arr, row, col)) ||
			isSymbol(left(arr, row, col)) ||
			isSymbol(right(arr, row, col)) ||
			isSymbol(bottomLeft(arr, row, col)) ||
			isSymbol(bottomRight(arr, row, col))
		);
	}

	if (row === arr.length - 1) {
		if (col === 0) {
			return (
				isSymbol(top(arr, row, col)) ||
				isSymbol(topRight(arr, row, col)) ||
				isSymbol(right(arr, row, col))
			);
		} else if (col === arr.length - 1) {
			return (
				isSymbol(top(arr, row, col)) ||
				isSymbol(left(arr, row, col)) ||
				isSymbol(topLeft(arr, row, col))
			);
		}
		return (
			isSymbol(top(arr, row, col)) ||
			isSymbol(left(arr, row, col)) ||
			isSymbol(right(arr, row, col)) ||
			isSymbol(topLeft(arr, row, col)) ||
			isSymbol(topRight(arr, row, col))
		);
	}

	if (col === 0) {
		return (
			isSymbol(top(arr, row, col)) ||
			isSymbol(topRight(arr, row, col)) ||
			isSymbol(right(arr, row, col)) ||
			isSymbol(bottomRight(arr, row, col)) ||
			isSymbol(bottom(arr, row, col))
		);
	}

	if (col === arr.length - 1) {
		return (
			isSymbol(top(arr, row, col)) ||
			isSymbol(topLeft(arr, row, col)) ||
			isSymbol(left(arr, row, col)) ||
			isSymbol(bottomLeft(arr, row, col)) ||
			isSymbol(bottom(arr, row, col))
		);
	}

	return (
		isSymbol(top(arr, row, col)) ||
		isSymbol(topLeft(arr, row, col)) ||
		isSymbol(left(arr, row, col)) ||
		isSymbol(bottomLeft(arr, row, col)) ||
		isSymbol(bottom(arr, row, col)) ||
		isSymbol(topRight(arr, row, col)) ||
		isSymbol(right(arr, row, col)) ||
		isSymbol(bottomRight(arr, row, col))
	);
}

function top(arr: string[], row: number, col: number) {
	return arr[row - 1].charAt(col);
}

function left(arr: string[], row: number, col: number) {
	return arr[row].charAt(col - 1);
}

function right(arr: string[], row: number, col: number) {
	return arr[row].charAt(col + 1);
}

function bottom(arr: string[], row: number, col: number) {
	return arr[row + 1].charAt(col);
}

function topLeft(arr: string[], row: number, col: number) {
	return arr[row - 1].charAt(col - 1);
}

function topRight(arr: string[], row: number, col: number) {
	return arr[row - 1].charAt(col + 1);
}

function bottomLeft(arr: string[], row: number, col: number) {
	return arr[row + 1].charAt(col - 1);
}

function bottomRight(arr: string[], row: number, col: number) {
	return arr[row + 1].charAt(col + 1);
}
