import { parseJson } from '../src/parser';
import fs from "node:fs";
import path from "node:path";

describe('JSON Parser', () => {
    it('should parser a valid simple JSON object', () => {
        const validJson = fs.readFileSync(path.join(__dirname, 'step1', 'valid.json'), 'utf8');
        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

    it('should reject an invalid JSON file', () => {
        const invalidJson = fs.readFileSync(path.join(__dirname, 'step1', 'invalid.json'), 'utf8');
        const result = parseJson(invalidJson);
        expect(result).toBe(false);
    });
});
