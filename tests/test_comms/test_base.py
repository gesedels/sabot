"""
Tests for 'sabot.comms.base'.
"""

import click
from sabot.comms.base import group
from sabot.items.book import Book


@group.command
@click.pass_obj
def mock(book: Book):
    click.echo(book.path)


def test_group(run_command):
    # success
    book, code, text = run_command(mock)
    assert book
    assert code == 0
    assert text == ":memory:\n"
