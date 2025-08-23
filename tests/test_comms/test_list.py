"""
Tests for 'sabot.comms.list'.
"""

from sabot.comms.list import list_


def test_list_(run_command):
    # success - no argument, with options
    _, code, text = run_command(list_, "-r", "-s", "name")
    assert code == 0
    assert text == "bravo\nalpha\n"

    # success - with argument, no options
    _, code, text = run_command(list_, "alph")
    assert code == 0
    assert text == "alpha\n"
