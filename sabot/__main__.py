"""
Main function for 'sabot'.
"""

from sabot.comms import group


def main(elems: list[str] | None = None):
    """
    Run the main Sabot program.
    """

    group.main(elems)


if __name__ == "__main__":
    main()
