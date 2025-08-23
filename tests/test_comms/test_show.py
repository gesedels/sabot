"""
Tests for 'sabot.comms.show'.
"""

from sabot.comms.show import show


def test_show(run_command):
    # success
    _, code, text = run_command(show, "alpha")
    assert code == 0
    assert text == "Alpha new.\n"
