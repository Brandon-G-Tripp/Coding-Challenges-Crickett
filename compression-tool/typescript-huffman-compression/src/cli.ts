import { readFileSync } from "node:fs";
import { countFrequencies } from "./compression";
import process from "node:process";

const filePath = process.argv[2];

try {
    const text = readFileSync(filePath, 'utf-8');
    const frequencies = countFrequencies(text);
    console.log(frequencies);
} catch (error) {
    console.error('Error reading file:', (error as Error).message);
    process.exit(1);
}
