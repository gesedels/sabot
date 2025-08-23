"""
Global unit testing fixtures.
"""

import click
import pytest
import sqlite3
from click.testing import CliRunner
from sabot.items.book import Book
from sabot.tools import sqls
from typing_extensions import Callable

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
    Return a new in-memory SQLite database.
    """

    dbse = sqlite3.connect(":memory:")
    dbse.row_factory = sqlite3.Row
    return dbse


@pytest.fixture(scope="function")
def mock(dbse: sqlite3.Connection) -> sqlite3.Connection:
    """
    Return a new in-memory SQLite database with pragma, schema and mock inserts.
    """

    dbse.executescript(sqls.PRAGMA + sqls.SCHEMA + INSERTS)
    dbse.commit()
    return dbse


@pytest.fixture(scope="function")
def run_command() -> Callable:
    """
    Return a function that returns the Book, exit code and output from a Command.
    """

    def run_command(comm: click.Command, *elems: str) -> tuple[Book, int, str]:
        book = Book(":memory:")
        book.dbse.executescript(INSERTS)
        rslt = CliRunner().invoke(comm, elems, catch_exceptions=False, obj=book)
        return book, rslt.exit_code, rslt.output

    return run_command
