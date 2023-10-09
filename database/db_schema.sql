create extension if not exists "uuid-ossp";

create table if not exists profile (
    id serial primary key,
    email varchar(32) unique not null,
    password varchar(32) not null,
    nickname varchar(32) not null,
    birth_date date not null,
    avatar_url varchar(1024)
    -- premium
);

create table if not exists artist (
    id serial primary key,
    name varchar(32) not null unique,
    avatar varchar(1024)
);

create table if not exists album (
    id serial primary key,
    name varchar(32) not null unique,
    artist_id int not null ,

    FOREIGN KEY (artist_id) references artist(id),
    preview varchar(1024),
    release_date date
);

create table if not exists track(
    id serial primary key,
    name varchar(32) not null,
    preview varchar(1024),
    content varchar(1024)
    -- song_text text
);

create table if not exists album_track(
    id serial primary key,
    album_id int  not null,
    track_id int  not null,
    FOREIGN KEY (album_id) references album(id),
    FOREIGN KEY (track_id) references track(id)

);

create table if not exists artist_track(
    id serial primary key,
    artist_id int not null,
    track_id int not null,
    FOREIGN KEY (artist_id) references artist(id),
    FOREIGN KEY (track_id) references track(id)
);

create table if not exists playlist (
                                        id serial primary key,
                                        name varchar(32) not null unique,
                                        creator_id int not null,
    FOREIGN KEY (creator_id) references profile(id),
    preview varchar(1024)
    );

create table if not exists playlist_track(
                                             id serial primary key,
                                             playlist_id int not null,
                                             FOREIGN KEY (playlist_id) references playlist(id),
                                             track_id int not null,
     FOREIGN KEY (track_id) references track(id)
    );

create table if not exists podcast(
                                      id serial primary key,
                                      name varchar(32) not null,
                                      artist_id int not null,
    FOREIGN KEY (artist_id) references artist(id),
    preview varchar(1024),
    description varchar(256),
    release_date date
    );

create table if not exists profile_playlist(
                                               id serial primary key,
                                               profile_id int not null,
                                               FOREIGN KEY (profile_id)references profile(id),
                                               playlist_id int not null,
    FOREIGN KEY (playlist_id) references playlist(id)
    );

create table if not exists profile_album(
                                            id serial primary key,
                                            profile_id int not null,
                                             FOREIGN KEY (profile_id) references profile(id),
                                             album_id int not null,
    FOREIGN KEY (album_id) references album(id)
    );

create table if not exists profile_podcast(
                                              id serial primary key,
                                              profile_id int not null,
                                              FOREIGN KEY (profile_id) references profile(id),
                                              podcast_id int not null,
    FOREIGN KEY (podcast_id) references podcast(id)
    );

create table if not exists session (
                                       id serial primary key,
                                       session_id uuid default uuid_generate_v4() not null unique,
    expiration timestamp with time zone not null default current_timestamp,
                             profile_id int not null references profile(id)
    );

-- table with like for songs:
--      id like_author
--      id like_song

-- table with like for albums:
--      id like_author
--      id like_albums

-- table with like for playlist:
--      id like_author
--      id like_playlist

-- table with like for podcasts:
--      id like_author
--      id like_podcast

create function session_delete_expired_rows() returns trigger
    language plpgsql
    AS $$
begin
delete from session where session.expiration < now();
return NEW;
end;
    $$;

create trigger session_delete_expired_rows_trigger
    after insert on session
    EXECUTE PROCEDURE session_delete_expired_rows();
