import { readFile, print } from '../utils';

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

type coordsItemMap = { [key: string]: string };

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
  const result: coordsItemMap[] = [];
  while (queue.length > 0) {
    const curr = queue.shift();
    if (!curr) continue;
    [row, col] = curr.split(',').map((c) => Number(c));
    const item = grid[row][col];
    result.push({ [curr]: item });
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

const LetterToSymbol: { [key: string]: string } = {
  F: '\u250F',
  J: '\u251B',
  L: '\u2517',
  '7': '\u2513',
  '|': '\u2503',
  '-': '\u2501',
};

export function pt2() {
  const grid = parseInput();
  const res = traverse(grid);
  print(res);
  const newMat: string[][] = [];
  for (let i = 0; i < grid.length; i++) {
    let tmp: string[] = [];
    for (let j = 0; j < grid[i].length; j++) {
      const pipe = res.find((r) => r[`${i},${j}`]);
      if (pipe) {
        tmp.push(LetterToSymbol[pipe[`${i},${j}`]] || pipe[`${i},${j}`]);
        // process.stdout.write(LetterToSymbol[pipe[`${i},${j}`]] || 'S');
      } else {
        tmp.push(' ');
        // process.stdout.write(' ');
      }
    }
    newMat.push(tmp);
    // process.stdout.write('\n');
  }

  let filled = floodFill(newMat, 59, 10, '.');
  filled.forEach((n) => {
    n.forEach((x) => {
      process.stdout.write(LetterToSymbol[x] || x);
    });
    process.stdout.write('\n');
  });

  return -1;
}

const floodFill = (
  image: string[][],
  sr: number,
  sc: number,
  newColor: string
) => {
  // Get the input which needs to be replaced.
  const current = image[sr][sc];

  // If the newColor is same as the existing
  // Then return the original image.
  if (current === newColor) {
    return image;
  }

  //Other wise call the fill function which will fill in the existing image.
  fill(image, sr, sc, newColor, current);

  //Return the image once it is filled
  return image;
};

let fillStack: any[][] = [];
// Write ma fill function iteratively using a stack.
const fill = (
  image: string[][],
  sr: number,
  sc: number,
  newColor: string,
  current: string
) => {
  // Push the starting element into the stack.
  fillStack.push([sr, sc]);

  // Iterate till the stack is not empty.
  while (fillStack.length) {
    // Pop the top element from the stack.
    const [row, col] = fillStack.pop()!;

    // If the current element is same as the starting element
    // Then replace it with the newColor.
    if (image[row][col] === current) {
      image[row][col] = newColor;
    }

    // Check if the top element is same as the starting element
    // If yes then push it into the stack.
    if (row - 1 >= 0 && image[row - 1][col] === current) {
      fillStack.push([row - 1, col]);
    }

    // Check if the bottom element is same as the starting element
    // If yes then push it into the stack.
    if (row + 1 < image.length && image[row + 1][col] === current) {
      fillStack.push([row + 1, col]);
    }

    // Check if the left element is same as the starting element
    // If yes then push it into the stack.
    if (col - 1 >= 0 && image[row][col - 1] === current) {
      fillStack.push([row, col - 1]);
    }

    // Check if the right element is same as the starting element
    // If yes then push it into the stack.
    if (col + 1 < image[row].length && image[row][col + 1] === current) {
      fillStack.push([row, col + 1]);
    }
  }
};
