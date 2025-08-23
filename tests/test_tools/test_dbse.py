"""
Tests for 'sabot.tools.dbse'.
"""

from sabot.tools import dbse


def test_order_by():
    # success
    code = dbse.order_by("select * from Notes {OB}", "name", reverse=True)
    assert code == "select * from Notes order by name desc"
