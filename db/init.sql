create table if not exists tiny_urls (
	id serial constraint tiny_urls_pk primary key,
	original_url varchar not null,
	meme_percentage FLOAT not null,
	created_at timestamp default current_timestamp
);

create table if not exists memes (
	id serial constraint memes_pk primary key,
	url varchar not null,
	added_at timestamp default current_timestamp
);

alter table tiny_urls owner to postgres;

alter table memes owner to postgres;

create unique index if not exists tiny_urls_id_uindex on tiny_urls (id);

create unique index if not exists memes_id_uindex on memes (id);

-- add more memes here and then perhaps keep this updated in some other way

insert into
	memes (url)
values
	('https://imgur.com/gallery/wJBTzF1');

CREATE EXTENSION tsm_system_rows;