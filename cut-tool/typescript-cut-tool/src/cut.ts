import * as fs from 'fs';
import * as readline from 'readline';

export async function cutSecondField(filePath: string): Promise<void> {
    const fileStream = fs.createReadStream(filePath);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });

    for await (const line of rl) {
        const fields = line.split('\t');
        if (fields.length >= 2) {
            console.log(fields[1]);
        }
    }
}
