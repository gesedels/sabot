"""
Base Click group definition.
"""

import click
from sabot import tools
from sabot.items.book import Book


@click.group("sabot")
@click.pass_context
def group(ctxt: click.Context):
    """
    Initialise the base Click group and context.
    """

    if not ctxt.obj:
        path = tools.clui.dbse_path("SABOT_DB")
        ctxt.obj = Book(path)
