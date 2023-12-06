import { readFile, sum } from '../utils';

export function pt1() {
  const cards = readFile('day4/input.txt')!.split('\n');
  let res: string[][] = [];
  cards.forEach((card) => {
    const numsWithSep = card.split(':')[1];
    let [winningNums, myNums] = numsWithSep.split('|');
    winningNums = winningNums.trim();
    myNums = myNums.trim();
    const winnersAsList = winningNums.split(' ').filter((s) => s !== '');
    const myNumsAsList = myNums.split(' ').filter((s) => s !== '');
    const tmp: string[] = [];
    myNumsAsList.forEach((num) => {
      if (winnersAsList.includes(num)) {
        tmp.push(num);
      }
    });
    res.push(tmp);
  });
  const powers = res.map((r) => {
    return r.length > 0 ? 2 ** (r.length - 1) : 0;
  });
  return sum(powers);
}

export function pt2() {
  const cards = readFile('day4/input.txt')!.split('\n');
  const cardInstances = cards.map((_) => 1);
  const matchingNumbersCt: number[] = [];
  cards.forEach((card) => {
    const numsWithSep = card.split(':')[1];
    let [winningNums, myNums] = numsWithSep.split('|');
    winningNums = winningNums.trim();
    myNums = myNums.trim();
    const winnersAsList = winningNums.split(' ').filter((s) => s !== '');
    const myNumsAsList = myNums.split(' ').filter((s) => s !== '');
    let tmpCt = 0;
    myNumsAsList.forEach((num) => {
      if (winnersAsList.includes(num)) {
        tmpCt += 1;
      }
    });
    matchingNumbersCt.push(tmpCt);
  });

  matchingNumbersCt.forEach((ct, i) => {
    let cardsToprocess = cardInstances[i];
    while (cardsToprocess !== 0) {
      let start = i + 1;
      let tmpCt = ct;
      while (tmpCt !== 0) {
        cardInstances[start] += 1;
        start++;
        tmpCt--;
      }
      cardsToprocess -= 1;
    }
  });

  return sum(cardInstances);
}
