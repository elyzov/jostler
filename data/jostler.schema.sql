CREATE TABLE user (
	id INTEGER PRIMARY KEY,
	login VARCHAR(32) UNIQUE NOT NULL,
	name VARCHAR(64) NOT NULL,
	email VARCHAR(96) UNIQUE NOT NULL,
	pass VARCHAR(128) UNIQUE NOT NULL
);

CREATE TABLE tick (
	id INTEGER PRIMARY KEY,
	user_id INTEGER NOT NULL,
	title VARCHAR(64) NOT NULL,
	rest REAL NOT NULL,
	created DATE,
	modified DATE,
	CONSTRAINT FK_user FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE deal (
	id INTEGER PRIMARY KEY,
	tick_id INTEGER NOT NULL,
	when DATE NOT NULL,			// transaction date/time.
	cat_id INTEGER NOT NULL,
	amount REAL NOT NULL,
	cy_id INTEGER NOT NULL, 	// currency.
	comment VARCHAR(128) NOT NULL,
	aux VARCHAR(16),			// auxiliary field for different purpose
	CONSTRAINT FK_tick FOREIGN KEY (tick_id) REFERENCES tick(id),
	CONSTRAINT FK_cat FOREIGN KEY (cat_id) REFERENCES cat(id),
	CONSTRAINT FK_cy FOREIGN KEY (cy_id) REFERENCES cy(id)
);

CREATE TABLE cy (
	id INTEGER PRIMARY KEY,
	abbr VARCHAR(3) UNIQUE NOT NULL,
	logo VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE tag(
	id INTEGER PRIMARY KEY,
	title VARCHAR(48) UNIQUE NOT NULL
);

CREATE TABLE deal_tag_junction (
	deal_id INTEGER,
	tag_id INTEGER,
	CONSTRAINT deal_tag_pk PRIMARY KEY (deal_id, tag_id),
	CONSTRAINT FK_deal FOREIGN KEY (deal_id) REFERENCES deal(id),
	CONSTRAINT FK_tag FOREIGN KEY (tag_id) REFERENCES tag(id)
)

CREATE TABLE cat (
	id INTEGER PRIMARY KEY,
	title VARCHAR(64) NOT NULL,
	parent_id INTEGER,
	deal_type_id INTEGER NOT NULL, // transaction type: income, outcome, transfer
	CONSTRAINT FK_parent_cat FOREIGN KEY (parent_id) REFERENCES cat(id),
	CONSTRAINT FK_deal_type FOREIGN KEY (deal_type_id) REFERENCES deal_type(id),
);

CREATE TABLE deal_type (
	id INTEGER PRIMARY KEY,
	title VARCHAR(32) UNIQUE NOT NULL
);