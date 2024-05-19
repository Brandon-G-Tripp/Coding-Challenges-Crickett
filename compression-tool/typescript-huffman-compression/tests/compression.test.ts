import { count } from "console";
import { countFrequencies } from "../src/compression";

describe('countFrequencies', () => {
    it('should count the frequencies of characters correctly', () => {
        const text = 'aabbccabc';
        const expectedFrequencies = { a: 3, b: 3, c: 3 };
        expect(countFrequencies(text)).toEqual(expectedFrequencies);
    });
});
