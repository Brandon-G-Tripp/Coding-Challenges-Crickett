import { assertEquals } from "https://deno.land/std@0.224.0/testing/asserts.ts";
import { sortFile } from "../main.ts";    

const tempDir = Deno.makeTempDirSync();
const wordsFilePath = `${tempDir}/words.txt`;


Deno.writeTextFileSync(wordsFilePath, `ZEBRA
ACTUAL
AGREE
AGREEMENT
AND
A
APPLE
BANANA
CHERRY
APPLE
ZEBRA
A`);


Deno.test("Sort file and limit output to 5 lines", async () => {
    const sortedLines = await sortFile(wordsFilePath, false);
    const firstFiveLines = sortedLines.slice(0, 5).join("\n");
    const expectedOutput = "A\nA\nACTUAL\nAGREE\nAGREEMENT";

    assertEquals(firstFiveLines, expectedOutput);
});

Deno.test("Sort file with unique option", async () => {
    const sortedLines = await sortFile(wordsFilePath, true);
    const expectedOutput = "A\nACTUAL\nAGREE\nAGREEMENT\nAND\nAPPLE\nBANANA\nCHERRY\nZEBRA";

    assertEquals(sortedLines.join("\n"), expectedOutput);
});

// Add this test to verify the actual command-line output
Deno.test("Verify command-line output without -u flag", async () => {
    const cmd = Deno.run({
        cmd: ["deno", "run", "--allow-read", "main.ts", wordsFilePath],
        stdout: "piped",
    });

    const output = await cmd.output();
    cmd.close();

    const decodedOutput = new TextDecoder().decode(output);
    const lines = decodedOutput.trim().split("\n");
    const expectedOutput = "A\nA\nACTUAL\nAGREE\nAGREEMENT\nAND\nAPPLE\nAPPLE\nBANANA\nCHERRY\nZEBRA\nZEBRA";

    assertEquals(lines.join("\n"), expectedOutput);

});


Deno.test("Verify command-line output with -u flag", async () => {
    const cmd = Deno.run({
        cmd: ["deno", "run", "--allow-read", "main.ts", "-u", wordsFilePath],
        stdout: "piped",
    });

    const output = await cmd.output();
    cmd.close();

    const decodedOutput = new TextDecoder().decode(output);
    const lines = decodedOutput.trim().split("\n");
    const expectedOutput = "A\nACTUAL\nAGREE\nAGREEMENT\nAND\nAPPLE\nBANANA\nCHERRY\nZEBRA";

    assertEquals(lines.join("\n"), expectedOutput);

});
