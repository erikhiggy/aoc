import { print, readFile } from '../utils';

function parseInput() {
  const input = readFile('day12/input.txt')!
    .split('\n')
    .map((line) => line.split(' '))
    .map((line) => {
      return {
        pattern: line[0],
        groups: line[1].split(',').map((group) => Number(group)),
      };
    });
  return input;
}

function valid(line: number[], groups: number[]): boolean {
  let n = line.length;
  let runs: number[] = [];

  let i = 0;
  while (i < n) {
    while (i < n && !line[i]) i += 1;
    if (i === n) break;
    let j = i;
    let c = 0;
    while (j < n && line[j]) {
      j += 1;
      c += 1;
    }
    runs.push(c);
    i = j;
  }
  return runs.length === groups.length && runs.every((x, i) => x === groups[i]);
}

function permute(pattern: string, groups: number[]): number {
  let line: number[] = [];
  let indexes: number[] = [];

  [...pattern].forEach((x, i) => {
    if (x === '.') {
      line.push(0);
    }
    if (x === '?') {
      line.push(-1);
      indexes.push(i);
    }
    if (x === '#') {
      line.push(1);
    }
  });

  let count = 0;
  for (let mask = 0; mask < 1 << indexes.length; mask++) {
    let lineCopy = [...line];
    for (let i = 0; i < indexes.length; i++) {
      if (mask & (1 << i)) {
        lineCopy[indexes[i]] = 0;
      } else {
        lineCopy[indexes[i]] = 1;
      }
    }
    if (valid(lineCopy, groups)) {
      count += 1;
    }
  }
  return count;
}

export function pt1() {
  const input = parseInput();
  let count = 0;
  input.forEach((line) => {
    count += permute(line.pattern, line.groups);
  });
  return count;
}

export function pt2() {
  return -1;
}
