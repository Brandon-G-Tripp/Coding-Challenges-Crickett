import fs from "node:fs";
import process from "node:process";
import yargs from 'yargs';

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

function countLinesFromInput(inputData) {
    const lines = inputData.split('\n');
    return lines[lines.length - 1] === '' ? lines.length - 1 : lines.length;
}

function countWordsFromInput(inputData) {
    const words = inputData.trim().split(/\s+/);
    return words.length;
}

function countBytesFromInput(inputData) {
    return Buffer.byteLength(inputData, 'utf8');
} 

async function runWordCount(argv = process.argv) {
    const args = yargs(argv.slice(2))
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
        .help('h')
        .alias('h', 'help')
        .argv;

    const filePath = args._[0];
    let inputData = '';

    if (!filePath) {
        process.stdin.setEncoding('utf8');
        for await (const chunk of process.stdin) {
            inputData += chunk;
        }
    }

    if(!args.chars && !args.bytes && !args.lines && !args.words) {
        const lineCount = filePath ? countLines(filePath) : countLinesFromInput(inputData);
        const wordCount = filePath ? countWords(filePath) : countWordsFromInput(inputData);
        const byteCount = filePath ? countBytes(filePath) : countBytesFromInput(inputData);

        if (lineCount === null || wordCount === null || byteCount === null) {
            console.error(`Error: Could not '${filePath ? `open file '${filePath}'` : 'read from standard input'}`); 
            process.exit(1);
        } 

        const result = `${lineCount} ${wordCount} ${byteCount} ${filePath ? `${filePath}` : ''}`.trim();
        console.log(result);
        return result;
    } 

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
    countBytesFromInput,
    countWordsFromInput,
    countLinesFromInput,
    runWordCount,
}; 
