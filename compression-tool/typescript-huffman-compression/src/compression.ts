export function countFrequencies(text: string): { [char: string]: number } {
    const frequencies: { [char: string]: number } = {};
    for (const char of text) {
        frequencies[char] = (frequencies[char] || 0) + 1;
    }
    return frequencies;
}
