# Sabot

**Sabot** is a command-line key-value storage system, written in Go 1.24 by Stephen Malone.

- See [`changes.md`][ch] for the complete changelog.
- See [`license.md`][li] for the open-source license.

[ch]: https://github.com/gesedels/sabot/blob/main/changes.md
[li]: https://github.com/gesedels/sabot/blob/main/license.md

## To Do

- [ ] `comms` packages:
  - [ ] `comms/comm` for the `Command` interface and collections.
  - [ ] `comms/help` for the `HelpCommand` help printing command.
  - [ ] `comms/list` for the `ListCommand` note listing/matching command.
- [ ] `items` packages:
  - [ ] `items/conf` for the `Conf` configuration pair type.
  - [ ] `items/line` for the `Line` event logging type.
  - [ ] `items/note` for the `Note` note metadata type.
  - [ ] `items/page` for the `Page` note content version type.
- [ ] `tools` packages:
  - [ ] `tools/clui` for creating FlagSets and getting the database path.
  - [ ] `tools/dbse` for database handling functions.
  - [ ] `tools/errs` for standardised error messages.
  - [ ] `tools/neat` for sanitisation/validation functions.
  - [ ] `tools/sqls` for pragma and schema definitions.

## Concepts

- [ ] Command interface has `Name() string`, `Parse(*FlagSet) error` and `Run(io.Writer) error`.
- [ ] items packages all have Create/Get/etc functions.
- [ ] tools/clui has .MainFlags and .SubFlags to return predefined FlagSets.
- [ ] tools/errs package has standardised user-facing error messages.
- [ ] tools/neat package validates and sanitises (eg: `Name(s) (string, error)`).
