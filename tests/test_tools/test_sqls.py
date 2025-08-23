"""
Tests for 'sabot.tools.sqls'.
"""

import sqlite3
from sabot.tools import sqls


def test_pragma(dbse: sqlite3.Connection):
    # success
    dbse.executescript(sqls.PRAGMA)
    drow = dbse.execute("pragma foreign_keys").fetchone()
    assert drow["foreign_keys"] == 1


def test_schema(dbse: sqlite3.Connection):
    # success
    dbse.executescript(sqls.SCHEMA)
    drow = dbse.execute("select count(*) from SQLITE_SCHEMA").fetchone()
    assert drow["count(*)"] > 0
