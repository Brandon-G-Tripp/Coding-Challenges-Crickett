def parse_json(input_string):
    index = 0

    def consume_whitespace():
        nonlocal index
        while index < len(input_string) and input_string[index].isspace():
            index += 1

    def consume_char(char):
        nonlocal index
        consume_whitespace()
        if index < len(input_string) and input_string[index] == char:
            index += 1
            return True
        return False

    def parse_string():
        nonlocal index
        if not consume_char('"'):
            return False
        while index < len(input_string) and input_string[index] != '"':
            if input_string[index] =='\\':
                index += 1
            index += 1
        if not consume_char('"'):
            return False
        return True

    def parse_key_value_pair():
        if not parse_string():
            return False
        if not consume_char(':'):
            return False
        if not parse_value():
            return False
        return True

    def parse_value():
        consume_whitespace()
        if parse_string():
            return True
        return False

    def parse_object():
        if not consume_char('{'):
            return False
        consume_whitespace()
        if consume_char('}'):
            consume_whitespace()
            return index == len(input_string)
        if not parse_key_value_pair():
            return False
        consume_whitespace()
        while consume_char(','):
            consume_whitespace()
            if not parse_key_value_pair():
                return False
            consume_whitespace()
        consume_whitespace()
        if not consume_char('}'):
            return False
        consume_whitespace()
        return index == len(input_string)

    return parse_object()
