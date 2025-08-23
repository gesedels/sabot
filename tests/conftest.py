"""
Global unit testing fixtures.
"""

import pytest
import sqlite3


@pytest.fixture(scope="function")
def dbse() -> sqlite3.Connection:
    """
    Return a new in-memory SQLite database connection.
    """

    dbse = sqlite3.connect(":memory:")
    dbse.row_factory = sqlite3.Row
    return dbse
