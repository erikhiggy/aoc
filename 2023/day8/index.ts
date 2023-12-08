import { readFile, print } from '../utils';

interface DirectionMap {
  [key: string]: string[];
}

function parseInput() {
  const input = readFile('day8/input.txt')!;
  const [instructions, dirs] = input.split('\n\n');
  const directionMap: DirectionMap = {};
  const directions = dirs.split('\n');
  for (const direction of directions) {
    const [from, to] = direction.split(' = ');
    const paths = to
      .replace(/\(|\)/g, '')
      .split(', ')
      .map((path) => path.trim());
    directionMap[from] = paths;
  }
  return { instructions, directionMap };
}

function traverse(instructions: string, directionMap: DirectionMap) {
  const start = 'AAA';
  const end = 'ZZZ';
  let queue = directionMap[start];
  let steps = 1;
  let i = 0;
  while (queue.length) {
    if (i === instructions.length) {
      i = 0;
    }
    const current =
      instructions.charAt(i) === 'L' ? queue.shift()! : queue.pop()!;
    queue = []; // we want a fresh queue each time
    // print(`Instruction: ${instructions.charAt(i)}`);
    // print(`Current: ${current}`);
    // print(`Queue: ${queue}`);
    if (current === end) {
      return steps;
    }
    const paths = directionMap[current];
    for (const path of paths) {
      queue.push(path);
    }
    steps++;
    i++;
  }
  return -1;
}

export function pt1() {
  const { instructions, directionMap } = parseInput();
  return traverse(instructions, directionMap);
}

export function pt2() {
  //   const [instructions, directionmap] = parseInput();
  return -1;
}
