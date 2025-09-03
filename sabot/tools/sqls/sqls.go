// Package sqls implements SQLite pragma and schema definitions.
package sqls

// Pragma is the default always-on database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Flags (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null unique,
		flag boolean not null
	);

	create table if not exists Lines (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		mode text    not null,
		body text    not null
	);

	create table if not exists Notes (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null unique
	);

	create table if not exists Pages (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null,
		body text    not null,
		hash text    not null,

		foreign key (note) references Notes(id) on delete cascade
	);

	create index if not exists FlagNames on Flags(name);
	create index if not exists NoteNames on Notes(name);
`
