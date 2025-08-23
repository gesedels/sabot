"""
Tests for 'sabot.items.note'.
"""

import time
import pytest
import sqlite3
from sabot.items.note import Note


@pytest.fixture(scope="function")
def note(mock: sqlite3.Connection) -> Note:
    return Note(mock, 1)


def test_init(note):
    # success
    assert note.dbse
    assert note.n_id == 1


def test_eq(note):
    # success
    assert note == note
    assert note != Note(note.dbse, 2)
    assert note != "not a Note"


def test_repr(note):
    # setup
    note.dbse = None

    # success
    assert repr(note) == "Note(None, 1)"


def test_str(note):
    # success
    assert str(note) == "alpha"


def test_init_(note):
    # success
    assert note.init == time.localtime(1000)


def test_name(note):
    # success
    assert str(note) == "alpha"


def test_delete(note):
    # success
    note.delete()
    code = "select exists(select 1 from Notes where n_id=1) as exst"
    drow = note.dbse.execute(code).fetchone()
    assert drow["exst"] == 0


def test_exists(note):
    # success
    assert note.exists()
