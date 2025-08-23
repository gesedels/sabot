"""
Tests for 'sabot.comms.list'.
"""

from sabot.comms.list import list_


def test_list_(run_command):
    # success - no argument
    _, code, text = run_command(list_)
    assert code == 0
    assert text == "alpha\nbravo\n"

    # success - with argument
    _, code, text = run_command(list_, "alph")
    assert code == 0
    assert text == "alpha\n"
