import { readFile, print, printMatrix } from '../utils';

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
      if (grid[i][j] === 'S') {
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

  return { result, visited };
}

export function pt1() {
  const grid = parseInput();
  const { result } = traverse(grid);
  return result.length / 2;
}

function countLines(
  i: number,
  j: number,
  grid: string[],
  visited: Set<string>
) {
  const line = grid[i];
  let count = 0;
  for (let k = 0; k < j; k++) {
    if (!visited.has(`${i},${k}`)) {
      continue;
    }
    if (
      line[k] === '|' ||
      line[k] === 'L' ||
      line[k] === 'J' ||
      line[k] === 'S'
    ) {
      count += 1;
    }
  }
  return count;
}

export function pt2() {
  const grid = parseInput();
  const { visited } = traverse(grid);
  let ans = 0;
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      const coords = `${i},${j}`;
      if (!visited.has(coords)) {
        const count = countLines(i, j, grid, visited);
        if (count % 2 === 1) {
          ans += 1;
        }
      }
    }
  }
  return ans;
}
