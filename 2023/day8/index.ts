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

function traverse(
  instructions: string,
  directionMap: DirectionMap,
  start: string,
  end?: string
) {
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
    if (!end) {
      if (current.endsWith('Z')) {
        return steps;
      }
    } else if (current === end) {
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

function traversePt2(instructions: string, directionMap: DirectionMap) {
  const starts = Object.keys(directionMap).filter((key) => key.endsWith('A'));
  const steps = starts.map((start, i) => {
    return traverse(instructions, directionMap, start);
  });
  return steps.reduce((acc, val) => LCM(acc, val));
}

export function pt1() {
  const { instructions, directionMap } = parseInput();
  return traverse(instructions, directionMap, 'AAA', 'ZZZ');
}

export function pt2() {
  const { instructions, directionMap } = parseInput();
  return traversePt2(instructions, directionMap);
}

function LCM(a: number, b: number) {
  return (a * b) / GCD(a, b);
}
function GCD(a: number, b: number) {
  if (b === 0) {
    return a;
  }
  return GCD(b, a % b);
}
