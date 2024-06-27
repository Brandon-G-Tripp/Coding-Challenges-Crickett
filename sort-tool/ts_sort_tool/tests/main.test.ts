import { assertEquals } from "https://deno.land/std@0.224.0/testing/asserts.ts";

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
CHERRY`);


Deno.test("Sort file and limit output to 5 lines", async () => {
    let cmd: Deno.Process | undefined;

    try {
        cmd = Deno.run({
            cmd: ["deno", "run", "--allow-read", "main.ts", wordsFilePath],
            stdout: "piped",
        });
        
        const output = await cmd.output();
        const lines = new TextDecoder().decode(output).trim().split("\n");
        const firstFiveLines = lines.slice(0, 5).join("\n");
        const expectedOutput = "A\nACTUAL\nAGREE\nAGREEMENT\nAND";


        assertEquals(firstFiveLines, expectedOutput);
    } catch(error) {
        console.error("Error:", error);
    } finally {
        // Close the command
        cmd?.close();
    }
});
