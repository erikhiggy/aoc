import { print, readFile } from '../utils';

function parseInput() {
  const input = readFile('day12/input.txt')!
    .split('\n')
    .map((line) => line.split(' '))
    .map((line) => {
      return {
        pattern: line[0].split(''),
        groups: line[1].split(',').map((group) => Number(group)),
      };
    });
  return input;
}

export function pt1() {
  const input = parseInput();
  return -1;
}

export function pt2() {
  return -1;
}
