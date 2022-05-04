create table if not exists tiny_urls (
	id serial constraint tiny_urls_pk primary key,
	original_url varchar not null,
	meme_percentage FLOAT not null,
	created_at timestamp default current_timestamp
);

alter table tiny_urls owner to postgres;

create unique index if not exists tiny_urls_id_uindex on tiny_urls (id);