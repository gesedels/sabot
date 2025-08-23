"""
Tests for 'sabot.tools.clui'.
"""

import os
import pytest
from click import ClickException
from pathlib import Path
from sabot.tools import clui


def test_dbse_path():
    # setup
    os.environ = {"NAME": "test.db"}

    # success - custom variable
    path = clui.dbse_path("NAME")
    assert path == Path("test.db")

    # setup
    os.environ = {"XDG_CONFIG_HOME": "/xdgc"}

    # success - XDG_CONFIG_HOME
    path = clui.dbse_path("")
    assert path == Path("/xdgc/sabot.db")

    # setup
    os.environ = {"HOME": "/home"}

    # success - HOME
    path = clui.dbse_path("")
    assert path == Path("/home/.sabot")

    # setup
    os.environ = {}

    # exception - cannot detect database path
    with pytest.raises(ClickException) as excp:
        clui.dbse_path("")
    excp.match("cannot detect database path")
