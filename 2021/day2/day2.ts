import { readFileSync } from 'fs';

const inputData = readFileSync('./day2/day2.txt').toString().split('\r\n');

const makeUsableData = (inputData: string[]) => {
    const input = inputData;
    const dirMap = input.map(dir => {
        const [direction, amount] = dir.split(' ');
        return {
            direction,
            amount: Number(amount)
        }
    });

    return dirMap;
};

export const day2SolutionPart1 = (): number => {
    const dirMap = makeUsableData(inputData);

    let x = 0;
    let y = 0;

    dirMap.forEach((entry) => {

        switch (entry.direction) {
            case 'forward':
                x += entry.amount;
                break;
            case 'down':
                y += entry.amount;
                break;
            case 'up':
                y -= entry.amount;
                break;
        }
    })

    return x*y
};

export const day2SolutionPart2 = (): number => {
    const dirMap = makeUsableData(inputData);

    let xPos = 0;
    let depth = 0;
    let aim = 0;

    dirMap.forEach((entry) => {
        switch (entry.direction) {
            case 'forward':
                xPos += entry.amount;
                depth += entry.amount * aim;
                break;
            case 'down':
                aim += entry.amount;
                break;
            case 'up':
                aim -= entry.amount;
                break;
        }
    })

    return xPos*depth;
};