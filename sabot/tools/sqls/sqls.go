// Package sqls implements SQLite query constants.
package sqls

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
`

// Schema is the default first-run database schema.
const Schema = `
	create table Notes (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null unique
	);

	create table Pages (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null references Notes(id) on delete cascade,
		body text    not null
	);
`
