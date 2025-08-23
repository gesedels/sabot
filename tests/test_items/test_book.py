"""
Tests for 'sabot.items.book'.
"""

import pytest
import sqlite3
from sabot.items.book import Book


@pytest.fixture(scope="function")
def book(mock: sqlite3.Connection) -> Book:
    book = Book(":memory:")
    book.dbse = mock
    return book


def test_init(book):
    # success
    assert book.dbse
    assert book.path == ":memory:"


def test_eq(book, tmp_path):
    # success
    assert book == book
    assert book != Book(tmp_path / "test.db")
    assert book != "not a Book"


def test_repr(book):
    # success
    assert repr(book) == "Book(':memory:')"


def test_create(book):
    # success
    note = book.create("name")
    assert note.name == "name"


def test_get(book):
    # success - existing Note
    note = book.get("alpha")
    assert note.name == "alpha"

    # success - non-existent Note
    note = book.get("nope")
    assert note is None


def test_match(book):
    # success
    notes = list(book.match("alph"))
    assert len(notes) == 1
    assert notes[0].name == "alpha"
