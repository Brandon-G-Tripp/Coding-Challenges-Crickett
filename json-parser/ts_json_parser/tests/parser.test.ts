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

    it('should parse a valid JSON object with different value types', () => {
        const validJson = `{
            "key1": true,
            "key2": false,
            "key3": null,
            "key4": "value",
            "key5": 101
        }`;

        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

    it('should parse a valid JSON object with an object value', () => {
        const validJson = `{
            "key": "value",
            "nested": {
                "key": "value"
            }
        }`;

        const result = parseJson(validJson);
        expect(result).toBe(true);
    });


    it('should parse a valid JSON object with an array value', () => {
        const validJson = `{
            "key": "value",
            "array": [1, 2, 3]
        }`;

        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

    it('should parse a valid JSON object with nested object and array values', () => {
        const validJson = `{
            "key": "value",
            "nested": {
                "key": "value",
                "array": [1, 2, 3]
            }
        }`;

        const result = parseJson(validJson);
        expect(result).toBe(true);
    });

    it('should parse a valid JSON object with empty object and array values', () => {
        const validJson = `{
            "key": "value",
            "empty_object": {},
            "empty_array": []
        }`;

        const result = parseJson(validJson);
        expect(result).toBe(true);
    });
});
