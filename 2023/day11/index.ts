import { print, printMatrix, readFile, sum } from '../utils';

function parseInput() {
  const input = readFile('day11/input.txt')!;
  return input.split('\n').map((row) => row.split(''));
}

function expand(input: string[][]) {
  // If a row contains only '.', then add a row full of '.' above it.
  for (let i = 0; i < input.length; i++) {
    const row = input[i];
    if (row.every((col) => col === '.')) {
      input.splice(
        i + 1,
        0,
        row.map(() => '.')
      );
      i += 1;
    }
  }

  // Now do the same for the columns.
  for (let i = 0; i < input[0].length; i++) {
    const col = input.map((row) => row[i]);
    if (col.every((row) => row === '.')) {
      input.forEach((row) => row.splice(i + 1, 0, '.'));
      i += 1;
    }
  }
}

function numberGalaxies(input: string[][]) {
  let count = 0;
  for (let i = 0; i < input.length; i++) {
    const row = input[i];
    for (let j = 0; j < row.length; j++) {
      const col = row[j];
      if (col === '#') {
        count += 1;
        row.splice(j, 1, count.toString());
      }
    }
  }
  return count;
}

class Point {
  x: number;
  y: number;
  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
}

class Cell {
  x: number;
  y: number;
  f: number;
  g: number;
  h: number;
  parent: Cell | undefined;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
    this.f = 0;
    this.g = 0;
    this.h = 0;
  }
}

function a_star(start: Point, end: Point, matrix: string[][]): Point[] {
  let openList: Cell[] = [],
    closedList: Cell[] = [],
    startCell = new Cell(start.x, start.y),
    endCell = new Cell(end.x, end.y);

  openList.push(startCell);

  while (openList.length > 0) {
    let lowInd = 0;
    for (let i = 0; i < openList.length; i++) {
      if (openList[i].f < openList[lowInd].f) lowInd = i;
    }
    let currentNode = openList[lowInd];

    if (currentNode.x == endCell.x && currentNode.y == endCell.y) {
      let curr = currentNode;
      let ret = [];
      while (curr.parent) {
        ret.push(new Point(curr.x, curr.y));
        curr = curr.parent;
      }
      return ret.reverse();
    }

    openList.splice(lowInd, 1);
    closedList.push(currentNode);

    let neighbors: Cell[] = [];
    const dirs = [
      { x: 0, y: -1 },
      { x: 1, y: 0 },
      { x: 0, y: 1 },
      { x: -1, y: 0 },
    ]; // Left, right, up, down
    for (let dir of dirs) {
      let x = currentNode.x + dir.x;
      let y = currentNode.y + dir.y;
      if (x < 0 || y < 0 || x >= matrix.length || y >= matrix[0].length)
        continue;
      neighbors.push(new Cell(x, y)); // Only cells up, down, left and right are considered.
    }

    for (let i = 0, il = neighbors.length; i < il; i++) {
      let neighbor = neighbors[i];
      if (
        closedList.findIndex(
          (cell) => cell.x === neighbor.x && cell.y === neighbor.y
        ) > -1
      )
        continue;

      let gScore = currentNode.g + 1,
        gScoreIsBest = false;

      if (
        openList.findIndex(
          (cell) => cell.x === neighbor.x && cell.y === neighbor.y
        ) === -1
      ) {
        gScoreIsBest = true;
        neighbor.h = heuristic(neighbor, endCell);
        openList.push(neighbor);
      } else if (gScore < neighbor.g) gScoreIsBest = true;

      if (gScoreIsBest) {
        neighbor.parent = currentNode;
        neighbor.g = gScore;
        neighbor.f = neighbor.g + neighbor.h;
      }
    }
  }
  return [];
}

function heuristic(pos0: Point, pos1: Point) {
  let d1 = Math.abs(pos1.x - pos0.x);
  let d2 = Math.abs(pos1.y - pos0.y);
  return d1 + d2;
}

function getCoords(input: string[][], value: string) {
  for (let i = 0; i < input.length; i++) {
    const row = input[i];
    for (let j = 0; j < row.length; j++) {
      const col = row[j];
      if (col === value) {
        return new Point(i, j);
      }
    }
  }
  return new Point(-1, -1);
}

function generatePairs(galaxiesCount: number) {
  const pairs: string[][] = [];
  for (let i = 1; i < galaxiesCount + 1; i++) {
    for (let j = i + 1; j < galaxiesCount + 1; j++) {
      pairs.push([i.toString(), j.toString()]);
    }
  }
  return pairs;
}

export function pt1() {
  const input = parseInput();
  print('Begin expanding of the universe...');
  expand(input);
  print('Universe expanded.');
  print('Begin numbering of galaxies...');
  const ct = numberGalaxies(input);
  print(`Number of galaxies: ${ct}`);
  print('Galaxies numbered.');
  print('Begin generating pairs...');
  const pairs = generatePairs(ct);
  print('Pairs generated.');
  print('Begin calculating distances...');
  const distances: number[] = [];
  for (const pair of pairs) {
    const start = getCoords(input, pair[0]);
    const end = getCoords(input, pair[1]);
    const path = a_star(start, end, input);
    distances.push(path.length);
  }
  print('Distances calculated.');
  return sum(distances);
}
