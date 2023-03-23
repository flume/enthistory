# Changelog

### [v0.4.2](https://github.com/flume/enthistory/compare/v0.4.1...v0.4.2) (2023-03-23)

* Dependency version updates
* Wrap generic type in hooks with `any()` to prevent linter errors. 
  See https://github.com/ent/ent/pull/3360 for more info

### [v0.4.1](https://github.com/flume/enthistory/compare/v0.4.0...v0.4.1) (2023-03-19)

* Change the auditing function to be off by default, able to turn on with `enthistory.WithAudit()`.
* Change `.Audit()` to return [][]string instead of bytes so less constraint for users

### [v0.4.0](https://github.com/flume/enthistory/compare/v0.3.2...v0.4.0) (2023-03-18)

* Drop `.EarliestHistory()`, `.LatestHistory()`, & `.HistoryAt()` and introduce normalized functions in their place on the HistoryQuery
  `.Latest()`, `.Earliest()`, & `.AsOf()`
* Introduce `.Next()` & `.Prev()` functions for getting the next/prev history before/after the current history.
* Introduce `.Diff()` function for diffing two histories of the same type and ref to identify the changes
* Introduce `.Audit()` function for creating an opinion audit log of the history tables

### [v0.3.2](https://github.com/flume/enthistory/compare/v0.3.1...v0.3.2) (2023-03-17)

* Fix bug introduced in v0.3.1 where func maps on templates were incorrect
* General code clean up

### [v0.3.1](https://github.com/flume/enthistory/compare/v0.3.0...v0.3.1) (2023-03-17)

* Handle Immutable Fields correctly in code generation & set them on appropriate history fields

### [v0.3.0](https://github.com/flume/enthistory/compare/v0.2.0...v0.3.0) (2023-03-16)

* Introduce `enthistory.Annotations` for better configurability on schemas, with ability to exclude
schemas from history tracking altogether. Also use annotations to mark history schemas instead of relying
on naming conventions.
* Introduce `Restore()` method on history models. Allows a user to restore a history row back to the actual
row in the original table. 
* Introduce common history query functions:
  * `LatestHistory()` - Returns the most recent history row for a tracked model
  * `EarliestHistory()` - Returns the first history row for a tracked model
  * `HistoryAt()` - given a time, will return the state of this model at that time via 
  the history_time field

### [v0.2.0](https://github.com/flume/enthistory/compare/v0.1.4...v0.2.0) (2023-03-16)

* Create `UpdatedBy()` option for configuring enthistory, when not supplied, no updated_by
field is tracked in history. Can specify type of the value (either int or string).

### [v0.1.4](https://github.com/flume/enthistory/compare/v0.1.3...v0.1.4) (2023-03-15)

* Variadic options when creating a new extension

### [v0.1.3](https://github.com/flume/enthistory/compare/v0.1.2...v0.1.3) (2023-03-13)

* Create Templates and Workflow (#1)
* Caveat Documentation (#2)
* Better character theming & expand on edges (#3)

### [v0.1.2](https://github.com/flume/enthistory/compare/v0.1.1...v0.1.2) (2023-03-13)

* Correct placement of historyActivated

### [v0.1.1](https://github.com/flume/enthistory/compare/v0.1.0...v0.1.1) (2023-03-13)

* Check if history hooks have been activated already

### [v0.1.0](https://github.com/flume/enthistory/compare/2aad2099edc62162830d9fc780c46e9e243f32cf...v0.1.0) (2023-03-13)

* Release enthistory under Flume Health
