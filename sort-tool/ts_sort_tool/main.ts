async function sortFile(filePath: string): Promise<string[]> {
    const fileContent = await Deno.readTextFile(filePath);
    const lines = fileContent.split("\n");
    lines.sort();
    return lines.filter((line) => line.trim() !== "");
}

if (import.meta.main) {
    const args = Deno.args;
    if (args.length < 1) {
        console.error("Please provide a file path as an argument.");
        Deno.exit(1);
    }

    const filePath = args[0];
    const sortedLines = await sortFile(filePath);

    const writer = new TextEncoder();
    const stdout = Deno.stdout;

    try {
        for (const line of sortedLines) {
            await stdout.write(writer.encode(line + "\n"));
        }
    } catch (error) {
        if (error instanceof Deno.errors.BrokenPipe) {
            // ignore broken pipe error
        } else {
            throw error;
        }
    }
}
