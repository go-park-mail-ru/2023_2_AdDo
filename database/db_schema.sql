create extension if not exists "uuid-ossp";
create table if not exists profile (
                                       id serial primary key,
                                       name text not null unique,
                                       password varchar(32) not null
    --avatar image
    );

create table if not exists session (
                                       id serial primary key,
                                       session_id uuid default uuid_generate_v4() not null unique,
    expiration timestamp with time zone not null default current_timestamp,
                             profile_id int not null references profile(id)
    );

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

create table if not exists artist (
                                      id serial primary key,
                                      name text not null unique
);

create table if not exists album (
                                     id serial primary key,
                                     name text not null unique,
    -- image path or url
                                     artist_id int not null references artist(id)
    );

create table if not exists playlist (
                                        id serial primary key,
                                        name text not null unique,
    -- image path or url
                                        profile_id int not null references profile(id)
    );

create table if not exists track(
                                    id serial primary key,
                                    name text not null,
                                    artist_id int references artist(id),
    playlist_id int references playlist(id),
    album_id int references album(id)
    -- content path or url
    -- song text
    );

create table if not exists podcast(
                                      id serial primary key,
                                      name text not null,
                                      artist_id int references artist(id),
    playlist_id int references playlist(id),
    album_id int references album(id)
    -- content path or url
    -- description
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
