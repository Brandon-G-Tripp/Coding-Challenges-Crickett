import fs from 'node:fs';
import { parseJson } from './parser';
import process from 'node:process';

const filePath = process.argv[2];

if (!filePath) {
    console.error('Please provide a JSON file path as a command line argument.');
    process.exit(1);
}

try {
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const result = parseJson(fileContent);

    if (result) {
        console.log('Valid JSON');
        process.exit(0);
    } else {
        console.log('Invalid JSON');
        process.exit(1);
    }
} catch (error) {
    console.error('Error reading the JSON file: ', error);
    process.exit(1);
}
