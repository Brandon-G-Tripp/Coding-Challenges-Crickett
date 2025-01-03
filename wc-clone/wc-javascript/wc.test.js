import fs from 'node:fs';

import { jest } from '@jest/globals';

import { Readable } from 'stream';

import { runWordCount, countBytes, countWords, countLines, countChars } from './wc.js';

describe('word count functions', () => {
    afterEach(() => {
        // clean up the temp file after each test
        if (fs.existsSync('test.txt')) {
            fs.unlinkSync('test.txt');
        } 
    });

    test('counts the number of lines in a file', () => {
        fs.writeFileSync('test.txt', 'Line 1\nLine 2\nLine 3\n');

        const count = countLines('test.txt');

        expect(count).toBe(3);
    });

    test('counts lines, words, and bytes when no option is provided', async () => {
        fs.writeFileSync('test.txt', 'Line 1\nLine 2\nLine 3\n');

        const result = await runWordCount(['node', 'wc.js', 'test.txt']);

        expect(result).toBe('3 6 21 test.txt')
    });

    test('counts the number of bytes in a file', () => {
        // Create a temp file with sample content
        fs.writeFileSync('test.txt', 'Sample content');

        // Call the countsBytes function
        const count = countBytes('test.txt');

        // Assert the expected byte count
        expect(count).toBe(14);
    }); 

    test('returns null for a non-existent file', () => {
        // Call the countBytes function with a non-existent file
        const count = countBytes('nonexistent.txt');

        // Assert that the count is null (indicating an error) 
        expect(count).toBeNull();
    });

    test('counts the number of words in a file', () => {
        fs.writeFileSync('test.txt', 'This is a sample file\nwith multiple words\non each line\n');

        const count = countWords('test.txt');

        expect(count).toBe(11);
    }); 

    test('counts the number of characters in a file', () => {
        fs.writeFileSync('test.txt', 'Sample content with 🚀 emoji');

        const charCount = countChars('test.txt');
        const byteCount = countBytes('test.txt');

        expect(charCount).toBe(27);
        expect(charCount).not.toBe(byteCount);
    }); 

    test('reads from standard input when no filename is provided', async () => {
        const stdin = 'Line 1\nLine 2\nLine 3\n';
        const expectedOutput = '3 6 21';

        const stdinStream = new Readable();
        stdinStream.push(stdin);
        stdinStream.push(null);

        // Mock standard input
        const stdinMock = jest.spyOn(process, 'stdin', 'get').mockReturnValue(stdinStream);

        // capture the console output
        const consoleLogMock = jest.spyOn(console, 'log').mockImplementation();

        await runWordCount(['node', 'wc.js']);

        expect(consoleLogMock).toHaveBeenCalledWith(expectedOutput);

        stdinMock.mockRestore();
        consoleLogMock.mockRestore()
    });
});
