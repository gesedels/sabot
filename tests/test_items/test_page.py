"""
Tests for 'sabot.items.page'.
"""

import time
import pytest
import sqlite3
from sabot.items.page import Page


@pytest.fixture(scope="function")
def page(mock: sqlite3.Connection) -> Page:
    return Page(mock, 1)


def test_init(page):
    # success
    assert page.dbse
    assert page.p_id == 1


def test_eq(page):
    # success
    assert page == page
    assert page != Page(page.dbse, 2)
    assert page != "not a Page"


def test_repr(page):
    # setup
    page.dbse = None

    # success
    assert repr(page) == "Page(None, 1)"


def test_str(page):
    # success
    assert str(page) == "Alpha old.\n"


def test_body(page):
    # success
    assert page.body == "Alpha old.\n"


def test_init_(page):
    # success
    assert page.init == time.localtime(1000)


def test_note(page):
    # success
    assert page.note == 1


def test_delete(page):
    # success
    page.delete()
    code = "select exists(select 1 from Pages where p_id=1) as exst"
    drow = page.dbse.execute(code).fetchone()
    assert drow["exst"] == 0


def test_exists(page):
    # success
    assert page.exists()
