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

    def parse_object():
        if not consume_char('{'):
            return False
        if not consume_char('}'):
            return False
        return True

    return parse_object()
