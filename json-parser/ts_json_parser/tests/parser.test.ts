import { parseJson } from '../src/parser';
import fs from "node:fs";
import path from "node:path";

describe('JSON Parser', () => {
    it('should parser a valid simple JSON object', () => {
        const validJson = fs.readFileSync(path.join(__dirname, 'step2', 'valid.json'), 'utf8');
        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

     it('should parse a valid JSON object with multiple key-value pairs', () => {
        const validJson = fs.readFileSync(path.join(__dirname, 'step2', 'valid2.json'), 'utf8');
        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

    it('should reject an invalid JSON object missing a key', () => {
        const invalidJson = fs.readFileSync(path.join(__dirname, 'step2', 'invalid.json'), 'utf8');
        const result = parseJson(invalidJson);
        expect(result).toBe(false);
    });

    it('should reject an invalid JSON object with an invalid key', () => {
        const invalidJson = fs.readFileSync(path.join(__dirname, 'step2', 'invalid2.json'), 'utf8');
        const result = parseJson(invalidJson);
        expect(result).toBe(false);
    });
});
