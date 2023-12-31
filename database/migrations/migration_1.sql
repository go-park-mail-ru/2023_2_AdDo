create extension if not exists "uuid-ossp";

create table if not exists profile (
    id uuid default uuid_generate_v4() primary key,
    email      varchar(32) unique not null,
    password   varchar(32)        not null,
    nickname   varchar(32)        not null,
    birth_date date               not null,
    avatar_url varchar(1024)
);

create table if not exists artist (
    id     serial primary key,
    name   varchar(32) not null unique,
    avatar varchar(1024)
);

create table if not exists playlist (
    id         serial primary key,
    name       varchar(32) not null unique,
    creator_id uuid         not null,
    foreign key (creator_id) references profile (id),
    preview    varchar(1024),
    creating_date date not null default now()
);

create table if not exists album (
    id           serial primary key,
    name         varchar(32) not null,
    artist_id    int         not null,
    foreign key (artist_id) references artist (id),
    preview      varchar(1024),
    release_date date not null
);

create table if not exists track(
    id serial primary key,
    name varchar(50) not null,
    preview varchar(1024),
    content varchar(1024),
    play_count int not null default 0
);

create table if not exists album_track (
    id       serial primary key,
    album_id int not null,
    track_id int not null,
    foreign key (album_id) references album (id),
    foreign key (track_id) references track (id)
);

create table if not exists artist_track (
    id       serial primary key,
    artist_id int not null,
    track_id int not null,
    foreign key (artist_id) references artist (id),
    foreign key (track_id) references track (id)
);

create table if not exists playlist_track (
    id          serial primary key,
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id),
    track_id    int not null,
    foreign key (track_id) references track (id)
);

create table if not exists profile_track (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id),
    track_id int not null,
    foreign key (track_id) references track (id)
);

create table if not exists profile_artist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id),
    artist_id int not null,
    foreign key (artist_id) references artist (id)
);

create table if not exists profile_album (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id),
    album_id int not null,
    foreign key (album_id) references album (id)
);

create table if not exists profile_playlist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id),
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id)
);

--create table if not exists podcast (
--    id           serial primary key,
--    name         varchar(32) not null,
--   artist_id    int         not null,
--  foreign key (artist_id) references artist (id),
-- preview      varchar(1024),
--description  varchar(256),
-- release_date date
--);