import os
from src.parser import parse_json


def test_parse_valid_json():
    with open(os.path.join(os.path.dirname(__file__), 'step1', 'valid.json'), 'r') as file:
        valid_json = file.read()
        assert parse_json(valid_json) == True


def test_parse_invalid_json():
    with open(os.path.join(os.path.dirname(__file__), 'step1', 'invalid.json'), 'r') as file:
        invalid_json = file.read()
        assert parse_json(invalid_json) == False


def test_parse_valid_simple_object():
    with open(os.path.join(os.path.dirname(__file__), 'step2', 'valid.json'), 'r') as file:
        valid_json = file.read()
        assert parse_json(valid_json) == True


def test_parse_valid_object_with_multiple_pairs():
    with open(os.path.join(os.path.dirname(__file__), 'step2', 'valid2.json'), 'r') as file:
        valid_json = file.read()
        assert parse_json(valid_json) == True


def test_parse_invalid_object_missing_key():
    with open(os.path.join(os.path.dirname(__file__), 'step2', 'invalid.json'), 'r') as file:
        invalid_json = file.read()
        assert parse_json(invalid_json) == False


def test_parse_invalid_object_with_invalid_key():
    with open(os.path.join(os.path.dirname(__file__), 'step2', 'invalid2.json'), 'r') as file:
        invalid_json = file.read()
        assert parse_json(invalid_json) == False

def test_parse_valid_object_with_different_value_types():
    valid_json = '''{
        "key1": true,
        "key2": false,
        "key3": null,
        "key4": "value",
        "key5": 101
    }'''

    assert parse_json(valid_json) == True


def test_parse_valid_object_with_object_value():
    valid_json = '''{
        "key": "value",
        "nested": {
            "key": "value"
        }
    }'''

    assert parse_json(valid_json) == True


def test_parse_valid_object_with_array_value():
    valid_json = '''{
        "key": "value",
        "array": [1, 2, 3]
    }'''

    assert parse_json(valid_json) == True


def test_parse_valid_object_with_nested_object_and_array():
    valid_json = '''{
        "key": "value",
        "nested": {
            "key": "value",
            "array": [1, 2, 3]
        }
    }'''

    assert parse_json(valid_json) == True


def test_parse_valid_object_with_empty_object_and_array():
    valid_json = '''{
        "key": "value",
        "empty_object": {},
        "empty_array": []
    }'''

    assert parse_json(valid_json) == True


def test_parse_invalid_object_with_missing_comma():
    invalid_json = '''{
        "key": "value"
        "another_key": "value"
    }'''

    assert parse_json(invalid_json) == False


def test_parse_invalid_array_with_missing_comma():
    invalid_json = '''{
        "key": "value",
        "array": [1 2 3]
    }'''

    assert parse_json(invalid_json) == False
