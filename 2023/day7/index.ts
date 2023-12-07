import { readFile, print, extract } from '../utils';

interface HandMap {
  hand: string;
  bid: number;
}

function parseInput() {
  const input = readFile('day7/input.txt')!.split('\n');
  let mapped = input.map((i): HandMap => {
    const [hand, bid] = i.split(' ');
    return { hand, bid: Number(bid) };
  });
  return mapped;
}

export function pt1() {
  const mapped = parseInput();

  const res = mapped.map((map) => {
    return { strength: checkHand(map.hand), bid: map.bid };
  });
  res.sort((a, b) => {
    if (a.strength[0] !== b.strength[0]) return b.strength[0] - a.strength[0];

    let i = 0;
    while (a.strength[1][i] === b.strength[1][i]) {
      i++;
    }
    return CARD_STRENGTH[b.strength[1][i]] - CARD_STRENGTH[a.strength[1][i]];
  });

  const ret = res.map((r, i) => {
    return { ...r, rank: res.length - i };
  });
  console.log('TOP', ret[0]);
  console.log('BOTTOM', ret[ret.length - 1]);
  let sum = 0;
  for (let i = 0; i < ret.length; i++) {
    sum += ret[i].bid * ret[i].rank;
  }
  return sum;
}

const CARD_STRENGTH: { [key: string]: number } = {
  '2': 1,
  '3': 2,
  '4': 3,
  '5': 4,
  '6': 5,
  '7': 6,
  '8': 7,
  '9': 8,
  T: 9,
  J: 10,
  Q: 11,
  K: 12,
  A: 13,
};

const HAND_TO_STRENGTH = {
  HIGH_CARD: 1,
  PAIR: 2,
  TWO_PAIR: 3,
  THREE_OF_A_KIND: 4,
  FULL_HOUSE: 5,
  FOUR_OF_A_KIND: 6,
  FIVE_OF_A_KIND: 7,
};

function checkHand(hand: string): [number, string] {
  const cardMap = new Map<string, number>();
  for (const card of hand) {
    if (cardMap.has(card)) {
      cardMap.set(card, cardMap.get(card)! + 1);
    } else {
      cardMap.set(card, 1);
    }
  }
  //   print(cardMap);

  let handType: number = 0;
  if (highCard(cardMap)) handType = HAND_TO_STRENGTH.HIGH_CARD;
  else if (pair(cardMap)) handType = HAND_TO_STRENGTH.PAIR;
  else if (twoPair(cardMap)) handType = HAND_TO_STRENGTH.TWO_PAIR;
  else if (threeOfAKind(cardMap)) handType = HAND_TO_STRENGTH.THREE_OF_A_KIND;
  else if (fullHouse(cardMap)) handType = HAND_TO_STRENGTH.FULL_HOUSE;
  else if (fourOfAKind(cardMap)) handType = HAND_TO_STRENGTH.FOUR_OF_A_KIND;
  else if (fiveOfAKind(cardMap)) handType = HAND_TO_STRENGTH.FIVE_OF_A_KIND;

  return [handType, hand];
}

function highCard(cardMap: Map<string, number>) {
  return [...cardMap.values()].every((v) => v === 1);
}

function pair(cardMap: Map<string, number>) {
  let oneTwoValue = false;
  for (const val of cardMap.values()) {
    if (oneTwoValue && val === 2) return false;
    if (val === 2) {
      oneTwoValue = true;
    }
  }
  return oneTwoValue;
}

function twoPair(cardMap: Map<string, number>) {
  let oneTwoValue = false;
  let twoTwoValues = false;
  for (const val of cardMap.values()) {
    if (oneTwoValue && twoTwoValues && val === 2) return false;
    if (oneTwoValue && val === 2) {
      twoTwoValues = true;
      continue;
    }
    if (val === 2) {
      oneTwoValue = true;
    }
  }
  return oneTwoValue && twoTwoValues;
}

function threeOfAKind(cardMap: Map<string, number>) {
  let hasThree = false;
  for (const val of cardMap.values()) {
    if (hasThree && val > 1) return false;
    if (val === 3) hasThree = true;
  }
  return hasThree;
}

function fullHouse(cardMap: Map<string, number>) {
  return cardMap.size === 2 && [...cardMap.values()].includes(3);
}

function fourOfAKind(cardMap: Map<string, number>) {
  return cardMap.size === 2 && [...cardMap.values()].includes(4);
}

function fiveOfAKind(cardMap: Map<string, number>) {
  return cardMap.size === 1;
}
