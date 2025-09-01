# Wishes

## 2025-09-01
### Safety & Security
- [ ] Change `book.Close()` to return an error.
- [ ] Use `defer try(book.Close())` in main script.
- [ ] Change default permissions to `0660` (user- and group-readable only).
- [ ] Add timeout options to database configuration.

### Data Inconsistencies
- [ ] Change multiple `time.Now()` calls into a single stored call.

### Unit Testing
- [ ] Change `test.DB` to `test.MockDB`.
- [ ] Use tools to ensure code is fully covered.

### Codebase Corrections
- [ ] Change package name in `sabot/comms` to `comms`.
- [ ] Ensure doc comments are used everywhere and consistently written.
- [ ] Consider changing codebase to single-character variable names.

### User-Facing Issues
- [ ] Write full readme for project.
- [ ] Create `clui.DatabasePath` to get database path based on environs.
- [ ] Create `tools/errs` package and decide on consistent user-friendly format.
- [ ] Ensure command help matches demo, e.g.: `Print NOTE if it exists` with `show NOTE`.
