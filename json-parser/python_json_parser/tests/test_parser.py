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
