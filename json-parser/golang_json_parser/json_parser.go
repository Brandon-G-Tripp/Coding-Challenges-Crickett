package main

import (
    "errors"
    "fmt"
    "strconv"
    "unicode"
)

type Parser struct {
    json  string
    index int
}

func (p *Parser) parseValue() (interface{}, error) {
    p.skipWhitespace()

    switch p.peek() {
    case '{':
        return p.parseObject()
    case '[':
        return p.parseArray()
    case '"':
        return p.parseString()
    case 't', 'f':
        return p.parseBoolean()
    case 'n':
        return p.parseNull()
    case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return p.parseNumber()
    default:
        return nil, fmt.Errorf("unexpected character: %c", p.peek())
    }
}

func (p *Parser) parseObject() (map[string]interface{}, error) {
    result := make(map[string]interface{})

    p.expect('{')
    p.skipWhitespace()

    if p.peek() == '}' {
        p.advance()
        return result, nil
    }

    for {
        key, err := p.parseString()
        if err != nil {
            return nil, err
        }

        p.skipWhitespace()

        if p.peek() != ':' {
            return nil, fmt.Errorf("expected ':', got %c", p.peek())
        }
        p.advance()

        p.skipWhitespace()

        value, err := p.parseValue()
        if err != nil {
            return nil, err
        }

        result[key] = value

        p.skipWhitespace()

        if p.peek() == ',' {
            p.advance()
            p.skipWhitespace()
        } else if p.peek() == '}' {
            p.advance()
            break
        } else {
            return nil, fmt.Errorf("unexpected character: %c", p.peek())
        }
    }

    return result, nil
}

func (p *Parser) parseArray() ([]interface{}, error) {
    var result []interface{}

    p.expect('[')
    p.skipWhitespace()

    if p.peek() == ']' {
        p.advance()
        return result, nil
    }

    for {
        value, err := p.parseValue()
        if err != nil {
            return nil, err
        }

        result = append(result, value)

        p.skipWhitespace()

        if p.peek() == ',' {
            p.advance()
            p.skipWhitespace()
        } else if p.peek() == ']' {
            p.advance()
            break
        } else {
            return nil, fmt.Errorf("unexpected character: %c", p.peek())
        }
    }

    return result, nil
}

func (p *Parser) parseString() (string, error) {
    var result string

    p.expect('"')

    for {
        ch := p.peek()

        if ch == '"' {
            p.advance()
            break
        } else if ch == '\\' {
            p.advance()
            ch = p.peek()

            switch ch {
            case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
                result += string(ch)
                p.advance()
            case 'u':
                p.advance()
                hexCode := p.json[p.index : p.index+4]
                p.index += 4

                value, err := strconv.ParseInt(hexCode, 16, 32)
                if err != nil {
                    return "", fmt.Errorf("invalid Unicode escape sequence: %s", hexCode)
                }

                result += string(rune(value))
            default:
                return "", fmt.Errorf("invalid escape sequence: \\%c", ch)
            }
        } else {
            result += string(ch)
            p.advance()
        }
    }

    return result, nil
}

func (p *Parser) parseNumber() (float64, error) {
    start := p.index

    if p.peek() == '-' {
        p.advance()
    }

    if p.peek() == '0' {
        p.advance()
    } else {
        if err := p.parseDigits(); err != nil {
            return 0, err
        }
    }

    if p.peek() == '.' {
        p.advance()
        if err := p.parseDigits(); err != nil {
            return 0, err
        }
    }

    if p.peek() == 'e' || p.peek() == 'E' {
        p.advance()

        if p.peek() == '+' || p.peek() == '-' {
            p.advance()
        }

        if err := p.parseDigits(); err != nil {
            return 0, err
        }
    }

    number, err := strconv.ParseFloat(p.json[start:p.index], 64)
    if err != nil {
        return 0, err
    }

    return number, nil
}

func (p *Parser) parseDigits() error {
    if !unicode.IsDigit(rune(p.peek())) {
        return errors.New("expected digits")
    }

    for unicode.IsDigit(rune(p.peek())) {
        p.advance()
    }

    return nil
}

func (p *Parser) parseBoolean() (bool, error) {
    if p.json[p.index:p.index+4] == "true" {
        p.index += 4
        return true, nil
    } else if p.json[p.index:p.index+5] == "false" {
        p.index += 5
        return false, nil
    }

    return false, fmt.Errorf("unexpected boolean value: %s", p.json[p.index:])
}

func (p *Parser) parseNull() (interface{}, error) {
    if p.json[p.index:p.index+4] == "null" {
        p.index += 4
        return nil, nil
    }

    return nil, fmt.Errorf("unexpected null value: %s", p.json[p.index:])
}

func (p *Parser) skipWhitespace() {
    for unicode.IsSpace(rune(p.peek())) {
        p.advance()
    }
}

func (p *Parser) peek() byte {
    if p.index < len(p.json) {
        return p.json[p.index]
    }
    return 0
}

func (p *Parser) advance() {
    p.index++
}

func (p *Parser) expect(ch byte) error {
    if p.peek() == ch {
        p.advance()
        return nil
    }
    return fmt.Errorf("expected %c, got %c", ch, p.peek())
}

func isValidJSON(json string) bool {
    parser := &Parser{json: json}

    _, err := parser.parseValue()
    if err != nil {
        return false
    }

    parser.skipWhitespace()

    return parser.index == len(json)
}
