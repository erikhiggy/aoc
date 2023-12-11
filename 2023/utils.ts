import fs from 'node:fs';

export function readFile(path: string) {
  try {
    const data = fs.readFileSync(path, {
      encoding: 'utf-8',
    });
    return data;
  } catch (err) {
    console.error('Could not read file', err);
  }
}

export function writeFile(path: string, data: string) {
  fs.writeFile(path, data, () => {
    console.log('file writen');
  });
}

export function sum(list: number[]) {
  return list.reduce((a, b) => a + b, 0);
}

export function extract(str: string) {
  return str.match(/\d+/g)!.map((s) => s);
}

export function print<T>(x: T) {
  console.log(x);
}

export function printMatrix(matrix: string[][]) {
  matrix.forEach((row) => {
    row.forEach((col) => {
      process.stdout.write(col);
    });
    process.stdout.write('\n');
  });
}
