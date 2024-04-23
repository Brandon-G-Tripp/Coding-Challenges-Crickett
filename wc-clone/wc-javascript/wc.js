const fs = require('fs');
const yargs = require('yargs');

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

const argv = yargs 
    .usage('Usage: $0 [options] <file>')
    .option('c', {
        alias: 'bytes',
        describe: 'Count bytes',
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
} else {
    console.error('Error: Missing -c flag');
    process.exit(1);
} 


module.exports = {
    countBytes,
}; 
