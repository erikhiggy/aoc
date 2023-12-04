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