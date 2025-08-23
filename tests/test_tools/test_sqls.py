"""
Tests for 'sabot.tools.sqls'.
"""

import sqlite3
from sabot.tools import sqls


def test_pragma(dbse: sqlite3.Connection):
    # success
    dbse.executescript(sqls.PRAGMA)
    code = "pragma foreign_keys"
    drow = dbse.execute(code).fetchone()
    assert drow["foreign_keys"] == 1


def test_schema(dbse: sqlite3.Connection):
    # success
    dbse.executescript(sqls.SCHEMA)
    code = "select count(*) from SQLITE_SCHEMA"
    drow = dbse.execute(code).fetchone()
    assert drow["count(*)"] > 0
