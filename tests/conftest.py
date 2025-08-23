"""
Global unit testing fixtures.
"""

import pytest
import sqlite3
from sabot.tools import sqls

INSERTS = """
    insert into Notes (init, name) values (1000, 'alpha');
    insert into Notes (init, name) values (2000, 'bravo');

    insert into Pages (init, note, body) values (1000, 1, 'Alpha old.' || char(10));
    insert into Pages (init, note, body) values (1100, 1, 'Alpha new.' || char(10));
    insert into Pages (init, note, body) values (2000, 2, 'Bravo.' || char(10));
"""


@pytest.fixture(scope="function")
def dbse() -> sqlite3.Connection:
    """
    Return a new in-memory SQLite database connection.
    """

    dbse = sqlite3.connect(":memory:")
    dbse.row_factory = sqlite3.Row
    return dbse


@pytest.fixture(scope="function")
def mock(dbse: sqlite3.Connection) -> sqlite3.Connection:
    """
    Return a new in-memory SQLite database connection containing pragma, schema
    and mock inserts.
    """

    dbse.executescript(sqls.PRAGMA + sqls.SCHEMA + INSERTS)
    dbse.commit()
    return dbse
