import process from "node:process";
import { cutSecondField } from "./cut";

async function main() {
    const args = process.argv.slice(2);

    if (args.length !== 2 || args[0] != '-f2') {
        console.log('Usage: cut -f2 <file>');
        process.exit(1);
    }

    const filePath = args[1];
    await cutSecondField(filePath);
}

main().catch((err) => {
    console.error('Error:', err);
    process.exit(1);
});
