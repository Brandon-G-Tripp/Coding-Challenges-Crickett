export function parseJson(input: string): boolean {
    let index = 0;

    function consumeWhitespace() {
        while (index < input.length && /\s/.test(input[index])) {
            index++;
        }
    } 

    function consumeChar(char: string): boolean {
        consumeWhitespace();
        if (index < input.length && input[index] === char) {
            index++;
            return true;
        }
        return false;
    }

    function parseObject(): boolean {
        if (!consumeChar('{')) {
            return false;
        }
        if (!consumeChar('}')) {
            return false;
        }
        return true;
    }

    return parseObject();
} 
