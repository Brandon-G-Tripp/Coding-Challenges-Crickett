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

    def parse_value():
        consume_whitespace()
        if input_string[index] == '"':
            return parse_string()
        elif input_string[index] == 't' or input_string[index] == 'f':
            return parse_boolean()
        elif input_string[index] == 'n':
            return parse_null()
        elif input_string[index].isdigit() or input_string[index] == '-':
            return parse_number()
        elif input_string[index] == '{':
            return parse_object()
        elif input_string[index] == '[':
            return parse_array()
        return False

    def parse_boolean():
        nonlocal index
        if input_string[index:index+4] == 'true':
            index += 4
            return True
        elif input_string[index:index+5] == 'false':
            index += 5
            return True
        return False

    def parse_null(): 
        nonlocal index 
        if input_string[index:index+4] == 'null':
            index += 4
            return True
        return False

    def parse_number():
        nonlocal index
        start = index 
        if input_string[index] == '-':
            index += 1
        while index < len(input_string) and input_string[index].isdigit():
            index +=1
        if index < len(input_string) and input_string[index] == '.':
            index += 1
            while index < len(input_string) and input_string[index].isdigit():
                index += 1
        return index > start

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

    def parse_array():
        if not consume_char('['):
            return False
        consume_whitespace()
        if consume_char(']'):
            return True
        while True:
            if not parse_value():
                return False
            consume_whitespace()
            if not consume_char(','):
                break
        if not consume_char(']'):
            return False
        return True
            

    def parse_key_value_pair():
        if not parse_string():
            return False
        consume_whitespace()
        if not consume_char(':'):
            return False
        consume_whitespace()
        if not parse_value():
            return False
        return True

    def parse_object():
        if not consume_char('{'):
            return False
        consume_whitespace()
        if consume_char('}'):
            return True
        while True:
            if not parse_key_value_pair():
                return False
            consume_whitespace()
            if not consume_char(','):
                break
            consume_whitespace()
        if not consume_char('}'):
            return False
        return True

    return parse_object()
