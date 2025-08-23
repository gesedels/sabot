"""
Command-line user interface functions.
"""

import os
from click import ClickException
from pathlib import Path


def dbse_path(name: str) -> Path:
    """
    Return a database path from a custom environment variable, $XDG_CONFIG_HOME or $HOME.
    """

    if path := os.environ.get(name, None):
        return Path(path)

    if path := os.environ.get("XDG_CONFIG_HOME", None):
        return Path(path) / "sabot.db"

    if path := os.environ.get("HOME", None):
        return Path(path) / ".sabot"

    raise ClickException("cannot detect database path")
