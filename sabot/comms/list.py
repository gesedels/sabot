"""
List Click command definition.
"""

import click
from sabot.comms.base import group
from sabot.items.book import Book

SORTS = click.Choice(["name", "init"])


@group.command("list", add_help_option=False)
@click.help_option("-h", "--help")
@click.argument("text", default="")
@click.option("-r", "--reverse", help="reverse sort", is_flag=True)
@click.option("-s", "--sort", default="name", help="sort by", type=SORTS)
@click.pass_obj
def list_(book: Book, text: str, sort: str, reverse: bool):
    """
    List all notes, or notes matching a substring.
    """

    for note in book.match(text, sort=sort, reverse=reverse):
        click.echo(note.name)
