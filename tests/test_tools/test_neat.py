"""
Tests for 'sabot.tools.neat'.
"""

import time as time_
from sabot.tools import neat


def test_body():
    # success
    body = neat.body("\tBody.\n")
    assert body == "Body.\n"


def test_name():
    # success
    name = neat.name("\tNAME_123!\n")
    assert name == "name-123"


def test_time():
    # success
    time = neat.time(1234)
    assert time == time_.localtime(1234)
