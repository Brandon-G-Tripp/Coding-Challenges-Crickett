import fs from 'node:fs';

import { countBytes, countLines } from './wc.js';

describe('countBytes', () => {
    afterEach(() => {
        // clean up the temp file after each test
        if (existsSync('test.txt')) {
            unlinkSync('test.txt');
        } 
    });

    test('counts the number of lines in a file', () => {
        fs.writeFileSync('test.txt', 'Line 1\nLine 2\nLine 3\n');

        const count = countLines('test.txt');

        expect(count).toBe(3);
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
});
