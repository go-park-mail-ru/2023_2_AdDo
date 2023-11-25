create index if not exists track_name_rus_gin_idx on track
    using gin(to_tsvector('russian', track.name));

create index if not exists album_name_rus_gin_idx on album
    using gin(to_tsvector('russian', album.name));

create index if not exists artist_name_rus_gin_idx on artist
    using gin(to_tsvector('russian', artist.name));

create index if not exists playlist_name_rus_gin_idx on playlist
    using gin(to_tsvector('russian', playlist.name));


create index if not exists track_name_eng_gin_idx on track
    using gin(to_tsvector('english', track.name));

create index if not exists album_name_eng_gin_idx on album
    using gin(to_tsvector('english', album.name));

create index if not exists artist_name_eng_gin_idx on artist
    using gin(to_tsvector('english', artist.name));

create index if not exists playlist_name_eng_gin_idx on playlist
    using gin(to_tsvector('english', playlist.name));
