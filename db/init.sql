create table if not exists tiny_urls (
	id serial constraint tiny_urls_pk primary key,
	long_url varchar not null,
	created_at timestamp default current_timestamp
);

create table if not exists memes (
	id serial constraint memes_pk primary key,
	url varchar not null,
	added_at timestamp default current_timestamp
);

create table if not exists tiny_urls_memes (
	id serial constraint tiny_urls_memes_pk primary key,
	tiny_urls_id int not null,
	memes_id int not null,
	constraint fk_tiny_urls foreign key (tiny_urls_id) references tiny_urls(id),
	constraint fk_memes foreign key (memes_id) references memes(id)
);

alter table tiny_urls owner to postgres;

alter table memes owner to postgres;

alter table tiny_urls_memes owner to postgres;

create unique index if not exists tiny_urls_id_uindex on tiny_urls (id);

create unique index if not exists memes_id_uindex on memes (id);

create unique index if not exists tiny_urls_memes_id_uindex on tiny_urls_memes (id);