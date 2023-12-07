create extension if not exists "uuid-ossp";

create table if not exists profile (
    id uuid default uuid_generate_v4() primary key,
    email      varchar(128) unique not null,
    -- очень маленькая вероятность встретить электронную почту длиннее 128 символов
    password   varchar(32)        not null,
    nickname   varchar(32)        not null,
    -- Посчитали, что не нужно давать пользователям задавать имя пользователя и пароль длиннее 32 символов для экономии места в базе и
    -- более быстрого получения хэша из пароля
    birth_date date               not null,
    avatar_url varchar(1024)
);

create table if not exists artist (
    id     serial primary key,
    name   varchar(128) not null unique,
    -- имена исполнителей редко превышают 32 символа
    avatar varchar(1024)
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
);

create table if not exists genre (
    id serial primary key,
    name varchar(128) not null unique
);

create table if not exists playlist (
    id         serial primary key,
    name       varchar(128) not null default 'New playlist',
    -- Посчитали, что не стоит давать пользователю создавать плейлисты и именами длиннее 128 символов
    creator_id uuid         not null,
    foreign key (creator_id) references profile (id) on delete cascade,
    preview    varchar(1024),
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    creating_date timestamptz not null default now(),
    is_private bool not null default false
);

create table if not exists album (
    id           serial primary key,
    name         varchar(128) not null,
    -- Отдельно хочется выделить, что varchar можно индексировать, в будущем это поможет нам
    -- осуществлять более быстрый поиск по артистам, альбомам и трэкам
    preview      varchar(1024),
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    release_date timestamptz,
    year int
);

create table if not exists track(
    id serial primary key,
    name varchar(128) not null,
    -- аналогично с альбомом и исполнителем
    preview varchar(1024),
    content varchar(1024),
    duration int not null,
    lyrics text,
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    play_count int not null default 0,
    valence double precision default 0,
    arousal double precision default 0
);

create table if not exists album_track (
    id       serial primary key,
    album_id int not null,
    track_id int not null,
    foreign key (album_id) references album (id) on delete cascade ,
    foreign key (track_id) references track (id) on delete cascade,
    constraint unique_album_track UNIQUE (album_id, track_id)
);

create table if not exists artist_track (
    id       serial primary key,
    artist_id int not null,
    track_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade ,
    foreign key (track_id) references track (id) on delete cascade,
    constraint unique_artist_track UNIQUE (artist_id, track_id)
);

create table if not exists artist_genre (
    id       serial primary key,
    artist_id int not null,
    genre_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade ,
    foreign key (genre_id) references genre (id) on delete cascade,
    constraint unique_artist_genre UNIQUE (artist_id, genre_id)
);

create table if not exists album_genre (
    id       serial primary key,
    album_id int not null,
    genre_id int not null,
    foreign key (album_id) references album (id) on delete cascade ,
    foreign key (genre_id) references genre (id) on delete cascade,
    constraint unique_album_genre UNIQUE (album_id, genre_id)
);

create table if not exists track_genre (
    id       serial primary key,
    track_id int,
    genre_id int,
    foreign key (track_id) references track (id) on delete cascade ,
    foreign key (genre_id) references genre (id) on delete cascade,
    constraint unique_track_genre UNIQUE (track_id, genre_id)
);

create table if not exists playlist_track (
    id          serial primary key,
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id) on delete cascade ,
    track_id    int not null,
    foreign key (track_id) references track (id) on delete cascade,
    constraint unique_playlist_track UNIQUE (playlist_id, track_id)
);

create table if not exists profile_track (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    track_id int not null,
    foreign key (track_id) references track (id) on delete cascade,
    constraint unique_profile_track UNIQUE (profile_id, track_id)
);

create table if not exists profile_artist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    artist_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade,
    constraint unique_profile_artist UNIQUE (profile_id, artist_id)
);

create table if not exists profile_album (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    album_id int not null,
    foreign key (album_id) references album (id) on delete cascade,
    constraint unique_profile_album UNIQUE (profile_id, album_id)
);

create table if not exists profile_playlist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id) on delete cascade,
    constraint unique_profile_playlist UNIQUE (profile_id, playlist_id)
);

create table if not exists artist_album (
    id         serial primary key,
    artist_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade ,
    album_id int not null,
    foreign key (album_id) references album (id) on delete cascade,
    constraint unique_artist_album UNIQUE (artist_id, album_id)
);

create table if not exists track_listen (
    id serial primary key,
    profile_id int not null,
    foreign key (profile_id) references profile(id) on delete cascade,
    track_id int not null,
    foreign key (track_id) references track(id) on delete cascade,
    duration int not null default 0,
    count int not null default 0
);

-- create table if not exists recommendation_cluster(
--      id serial primary key,
--      centre double precision[] not null default [0]
-- );
--
-- create table if not exists track_cluster(
--     id serial primary key,
--     track_id uuid not null,
--     foreign key (track_id) references track (id) on delete cascade ,
--     cluster_id int not null,
--     foreign key (cluster_id) references recommendation_cluster (id) on delete cascade,
--     constraint unique_track_cluster UNIQUE (track_id, cluster_id),
--     onehot_genre_pos int not null,
--     onehot_artist_pos int not null
-- );
