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

    function parseString(): boolean {
        if (!consumeChar('"')) {
            return false;
        }
        while (index < input.length && input[index] !== '"') {
            index++;
        }
        if (!consumeChar('"')) {
            return false;
        }
        return true;
    }

    function parseKeyValuePair(): boolean {
        if (!parseString()) {
            return false;
        }

        if (!consumeChar(':')) {
            return false;
        }

        if (!parseString()) {
            return false;
        }

        return true;
    }

    function parseObject(): boolean {
        if (!consumeChar('{')) {
            return false;
        }

        consumeWhitespace();

        if (consumeChar('}')) {
            return true;
        }

        if (!parseKeyValuePair()) {
            return false;
        }
        while (consumeChar(',')) {
            if (!parseKeyValuePair()) {
                return false;
            }
        }
        if (!consumeChar('}')) {
            return false;
        }

        return true;
    }

    return parseObject();
} 
