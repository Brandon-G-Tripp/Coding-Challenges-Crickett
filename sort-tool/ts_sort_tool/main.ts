async function sortFile(filePath: string, unique: boolean = false): Promise<string[]> {
    const fileContent = await Deno.readTextFile(filePath);
    const lines = fileContent.split("\n").filter(line => line.trim() !== "");
    lines.sort();
    return unique ? [...new Set(lines)] : lines;
}

if (import.meta.main) {
    const args = Deno.args;
    let unique = false;
    let filePath = "";

    for (let i = 0; i < args.length; i++){
        if (args[i] === "-u") {
            unique = true;
        } else {
            filePath = args[i];
        }
    }


    if (filePath === "") {
        console.error("Please provide a file path as an argument.");
        Deno.exit(1);
    }

    const sortedLines = await sortFile(filePath, unique);
    
    const writer = new TextEncoder();
    const stdout = Deno.stdout;

    try {
        for (let i = 0; i < sortedLines.length; i++) {
            await stdout.write(writer.encode(sortedLines[i]));
            if (i < sortedLines.length - 1) {
                await stdout.write(writer.encode("\n"));
            }
        }
    } catch (error) {
        if (error instanceof Deno.errors.BrokenPipe) {
            // ignore broken pipe error
        } else {
            console.error(`Error writing output: ${error}`);
            throw error;
        }
    }
}

export { sortFile };

