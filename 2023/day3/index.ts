import { readFile } from '../utils';

export function pt1() {
  const input = readFile('day3/input.txt');
  const rows = input?.split('\n') || [];
  console.log(rows);
  let res: number = 0;
  let possiblePartNumber = '';
  let validPartNumber: boolean[] = [];
  for (let i = 0; i < rows.length; i++) {
    for (let j = 0; j < rows[i].length; j++) {
      const item = rows[i][j];
      // Just move on if we encounter a symbol
      // or a period with no possiblePartNumber contents.
      if ((isSymbol(item) || isPeriod(item)) && possiblePartNumber === '') {
        continue;
      }
      // If the item is not a period and it's not a symbol
      // add the digit to the possiblePartNumber string
      // and check all around it to see if it's a part number.
      if (!isPeriod(item) && !isSymbol(item)) {
        possiblePartNumber += item;
        validPartNumber.push(isPartNumber(rows, i, j));
      } else if (
        (isPeriod(item) || isSymbol(item)) &&
        possiblePartNumber !== ''
      ) {
        // If we hit a period, AND we have contents in the possiblePartNumber
        // string, check if one of the entries in the validPartNumber array
        // is true, and if so, add the possibleNumberString to the result list.
        // Regardless, we want to clear both the string and array.
        if (validPartNumber.some((v) => v === true)) {
          res += parseInt(possiblePartNumber, 10);
          console.log(possiblePartNumber);
          console.log(validPartNumber);
        }
        possiblePartNumber = '';
        validPartNumber = [];
      }
    }
  }

  return res;
}

function isSymbol(str: string) {
  return !isPeriod(str) && isNaN(parseInt(str, 10));
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
