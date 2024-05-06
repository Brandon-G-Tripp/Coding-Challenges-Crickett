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
