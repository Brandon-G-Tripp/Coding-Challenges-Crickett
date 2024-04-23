import fs from "node:fs";
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

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
        return lines.length;
    } catch(error) {
        if (error.code === 'ENOENT') {
            return null;
        } 
        throw error;
    } 
}


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
    .demandCommand(1, 'Please provide a file')
    .help('h')
    .alias('h', 'help')
    .argv;

const filePath = argv._[0];

if (argv.bytes) {
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
} else {
    console.error('Error: Missing -c or -l flag');
    process.exit(1);
} 


export {
    countBytes,
    countLines,
}; 
