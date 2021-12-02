import { readFileSync } from 'fs';

const readInput = () => {
  return readFileSync('./day1/day1.txt').toString().split('\n');
};

export const day1SolutionPart1 = () => {
  let counter = 0;
  const dataArr = readInput();
  for (let i = 1; i < dataArr.length; i++) {
    if (Number(dataArr[i]) > Number(dataArr[i-1])) {
      counter++;
    }
  }

  return counter;
};

export const day1SolutionPart2 = () => {
  let counter = 0;
  const dataArr = readInput();
  let prevSum = 0;
  for (let i = 2; i < dataArr.length; i++) {
    const sumOfThree = Number(dataArr[i]) + Number(dataArr[i-1]) + Number(dataArr[i-2]);
    if (i !== 2 && sumOfThree > prevSum) {
      counter++;
    }
    prevSum = sumOfThree;
  }
  return counter;
};