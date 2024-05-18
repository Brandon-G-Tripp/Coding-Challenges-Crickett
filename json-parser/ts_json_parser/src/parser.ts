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

    function parseValue(): boolean {
        consumeWhitespace();

        if (input[index] === '"') {
            return parseString();
        } else if (input[index] === 't' || input[index] === 'f') {
            return parseBoolean();
        } else if (input[index] === 'n') {
            return parseNull();
        } else if (/[-0-9]/.test(input[index])) {
            return parseNumber();
        } else if (input[index] === '{') {
            return parseObject();
        } else if (input[index] === '[') {
            return parseArray();
        }

        return false;
    }

    function parseBoolean(): boolean {
        if (input.slice(index, index + 4) === 'true') {
            index += 4;
            return true;
        } else if (input.slice(index, index + 5) === 'false') {
            index += 5;
            return true;
        }
        return false;
    }

    function parseNull(): boolean {
        if (input.slice(index, index + 4) === 'null') {
            index += 4;
            return true;
        }
        return false;
    }

    function parseNumber(): boolean {
        const start = index;
        if (input[index] === '-') {
            index++;
        }
        while (index < input.length && /[0-9]/.test(input[index])) {
            index++;
        }
        if (input[index] === '.') {
            index++;
            while (index < input.length && /[0-9]/.test(input[index])) {
                index++;
            }
        }
        return index > start;
    }

    function parseString(): boolean {
        if (!consumeChar('"')) {
            return false;
        }
        while (index < input.length && input[index] !== '"') {
            if (input[index] === '\\') {
                index++;
            }
            index++;
        }
        if (!consumeChar('"')) {
            return false;
        }
        return true;
    }

    function parseArray(): boolean {
        if (!consumeChar('[')) {
            return false;
        }

        consumeWhitespace();

        if (consumeChar(']')) {
            return true;
        }

        while (true) {
            if (!parseValue()) {
                return false;
            }

            consumeWhitespace();

            if (!consumeChar(',')) {
                break;
            }

            consumeWhitespace();
        }

        if (!consumeChar(']')) {
            return false;
        }

        return true;
    }

    function parseKeyValuePair(): boolean {
        if (!parseString()) {
            return false;
        }

        consumeWhitespace();

        if (!consumeChar(':')) {
            return false;
        }

        consumeWhitespace();

        if (!parseValue()) {
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

        while (true) {
            if (!parseKeyValuePair()) {
                return false;
            }

            consumeWhitespace();

            if (!consumeChar(',')) {
                break;
            }

            consumeWhitespace();
        }

        if (!consumeChar('}')) {
            return false;
        }

        return true;
    }

    return parseObject();
}
