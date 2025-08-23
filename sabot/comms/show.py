"""
Show Click command definition.
"""

import click
from sabot.comms.base import group
from sabot.items.book import Book


@group.command("show", add_help_option=False)
@click.argument("name", required=True)
@click.help_option("-h", "--help")
@click.pass_obj
def show(book: Book, name: str):
    """
    Print the contents of a note, if it exists.
    """

    if note := book.get(name):
        if page := note.latest():
            click.echo(page.body, nl=False)
