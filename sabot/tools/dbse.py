"""
Database query helper functions.
"""


def order_by(code: str, name: str, *, reverse: bool = False) -> str:
    """
    Return a SQLite query with {OB} replaced with an 'order by' clause.
    """

    mode = "desc" if reverse else "asc"
    return code.format(OB=f"order by {name} {mode}")
