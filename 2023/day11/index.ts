import { print, printMatrix, readFile, sum } from '../utils';

function parseInput() {
  const input = readFile('day11/input.txt')!;
  return input.split('\n').map((row) => row.split(''));
}

function expand(input: string[][], factor: number) {
  // If a row contains only '.', then add a row full of '.' above it.
  for (let i = 0; i < input.length; i++) {
    const row = input[i];
    if (row.every((col) => col === '.')) {
      // add factor rows above it.
      for (let j = 0; j < factor; j++) {
        input.splice(
          i,
          0,
          row.map(() => '.')
        );
        i += 1;
      }
    }
  }

  // Now do the same for the columns.
  for (let i = 0; i < input[0].length; i++) {
    const col = input.map((row) => row[i]);
    if (col.every((row) => row === '.')) {
      // add factor columns to the left of it.
      for (let j = 0; j < factor; j++) {
        input.forEach((row) => row.splice(i + 1, 0, '.'));
        i += 1;
      }
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

function minkowskiDistance(p1: Point, p2: Point, m: number) {
  return (Math.abs(p1.x - p2.x) ** m + Math.abs(p1.y - p2.y) ** m) ** (1 / m);
}

function manhattanDistance(p1: Point, p2: Point) {
  return minkowskiDistance(p1, p2, 1);
}

export function pt1() {
  const input = parseInput();
  // expand(input, 99);
  const ct = numberGalaxies(input);
  const pairs = generatePairs(ct);
  const distances: number[] = [];
  for (const pair of pairs) {
    const start = getCoords(input, pair[0]);
    const end = getCoords(input, pair[1]);
    const md = manhattanDistance(start, end);
    distances.push(md);
  }
  return sum(distances);
}

export function pt2() {
  const input = parseInput();
  const ct = numberGalaxies(input);
  const pairs = generatePairs(ct);
  const distances: number[] = [];
  for (const pair of pairs) {
    const start = getCoords(input, pair[0]);
    const end = getCoords(input, pair[1]);
    const md = manhattanDistance(start, end);
    distances.push(md);
  }
  return sum(distances);
}

// This is part 2
// 1 -> 8,870,229
// 10 -> 12,139,794 + 3,269,565
// 100 -> 44,835,444 + 32,695,650
// 1,000 -> 371,791,944 + 326,956,500
// 10,000 -> 3,641,356,944 + 3,269,565,000
// 100,000 -> 36,337,006,944 + 32,695,650,000
// 1,000,000 -> 363,293,506,944 + 326,956,500,000
