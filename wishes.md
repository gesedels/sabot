# Wishes

## 2025-09-01
### Safety & Security
- [x] Change `book.Close()` to return an error.
- [x] Use `try(book.Close())` in main script.
- [x] Change default permissions to `0660` (user- and group-readable only).
- [x] Add timeout options to database configuration.

### Data Inconsistencies
- [ ] Change multiple `time.Now()` calls into a single stored call.
- [ ] Add `neat.Pairs` to return a consistent database dict with time values.

### Unit Testing
- [x] Change `test.DB` to `test.MockDB`.
- [ ] Use tools to ensure code is fully covered.

### Codebase Corrections
- [x] Change package name in `sabot/comms` to `comms`.
- [ ] Ensure doc comments are used everywhere and consistently written.
- [ ] Consider changing codebase to single-character variable names.

### User-Facing Issues
- [ ] Write full readme for project.
- [ ] Create `clui.DatabasePath` to get database path based on environs.
- [ ] Create `tools/errs` package and decide on consistent user-friendly format.
- [ ] Ensure command help matches demo, e.g.: `Print NOTE if it exists` with `show NOTE`.
- [ ] Add -d debug command and use it to reveal logging throughout the system.

### New Command System
- [ ] Remove `Command` type and replace with a single `Run` function:
  - Switch logic based on parsed flags, not subcommands.
  - E.g.: `-d` for debug, `-e` for edit, use a FlagSet.
