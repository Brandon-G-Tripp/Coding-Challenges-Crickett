import fs from "node:fs";
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

function countChars(filePath) {
    try {
        const data = fs.readFileSync(filePath, 'utf8');
        return Array.from(data).length;
    } catch (error) {
        if (error.code === 'ENOENT') {
            return null;
        } 
        throw error;
    }
}

function countBytes(filePath) {
    try {
        const data = fs.readFileSync(filePath);
        return data.length;
    } catch (error) {
        if (error.code === 'ENOENT') {
            return null;
        } 
        throw error;
    } 
} 

function countLines(filePath) {
    try {
        const data = fs.readFileSync(filePath, 'utf8');
        const lines = data.split('\n');
        return lines[lines.length - 1] === '' ? lines.length - 1 : lines.length;
    } catch(error) {
        if (error.code === 'ENOENT') {
            return null;
        } 
        throw error;
    } 
}

function countWords(filePath) {
    try {
        const data = fs.readFileSync(filePath, 'utf8');
        const words = data.trim().split(/\s+/);
        return words.length;
    } catch (error) {
        if (error.code === 'ENOENT') {
            return null;
        }
        throw error;
    }
}

function runWordCount() {
    const argv = yargs(hideBin(process.argv))
        .scriptName('wc')
        .usage('Usage: $0 [options] <file>')
        .option('c', {
            alias: 'bytes',
            describe: 'Count bytes',
            type: 'boolean',
        })
        .option('l', {
            alias: 'lines',
            describe: 'Count lines',
            type: 'boolean',
        })
        .option('w', {
            alias: 'words',
            describe: 'Count words',
            type: 'boolean',
        })
        .option('m', {
            alias: 'chars',
            describe: 'Count characters',
            type: 'boolean',
        })
        .demandCommand(1, 'Please provide a file')
        .help('h')
        .alias('h', 'help')
        .argv;

    const filePath = argv._[0];

    if (argv.chars) {
        const count = countChars(filePath);
        if (count === null) {
            console.error(`Error: Could open file '${filePath}'`);
            process.exit(1);
        } 
        console.log(`${count} ${filePath}`);
    } else if (argv.bytes) {
        const count = countBytes(filePath);
        if (count === null) {
            console.error(`Error: Could not open file '${filePath}'`);
            process.exit(1);
        } 
        console.log(`${count} ${filePath}`);
    } else if (argv.lines) {
        const count = countLines(filePath);
        if (count === null) {
            console.error(`Error: Could not open file '${filePath}'`);
            process.exit(1);
        }
        console.log(`${count} ${filePath}`);
    } else if (argv.words) {
        const count = countWords(filePath);
        if (count === null) {
            console.error(`Error: Could not open file '${filePath}'`);
            process.exit(1);
        }
        console.log(`${count} ${filePath}`);
    } else {
        console.error('Error: Missing -c or -l flag');
        process.exit(1);
    } 
}

if (import.meta.url === `file://${process.argv[1]}`) {
    runWordCount();
}


export {
    countBytes,
    countLines,
    countWords,
    countChars,
}; 
