create extension if not exists "uuid-ossp";

create table if not exists profile (
    id uuid default uuid_generate_v4() primary key,
    email      varchar(32) unique not null,
    -- очень маленькая вероятность встретить электронную почту длиннее 32 символов
    password   varchar(32)        not null,
    nickname   varchar(32)        not null,
    -- Посчитали, что не нужно давать пользователям задавать имя пользователя и пароль длиннее 32 символов для экономии места в базе и
    -- более быстрого получения хэша из пароля
    birth_date date               not null,
    avatar_url varchar(1024)
);

create table if not exists artist (
    id     serial primary key,
    name   varchar(32) not null unique,
    -- имена исполнителей редко превышают 32 символа
    avatar varchar(1024)
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
);

create table if not exists playlist (
    id         serial primary key,
    name       varchar(32) not null unique,
    -- Посчитали, что не стоит давать пользователю создавать плейлисты и именами длиннее 32 символов
    creator_id uuid         not null,
    foreign key (creator_id) references profile (id) on delete cascade,
    preview    varchar(1024),
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    creating_date timestamptz not null default now()
);

create table if not exists album (
    id           serial primary key,
    name         varchar(32) not null,
    -- Отдельно хочется выделить, что varchar можно индексировать, в будущем это поможет нам
    -- осуществлять более быстрый поиск по артистам, альбомам и трэкам
    artist_id    int         not null,
    foreign key (artist_id) references artist (id) on delete cascade,
    preview      varchar(1024),
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    release_date date not null
);

create table if not exists track(
    id serial primary key,
    name varchar(50) not null,
    -- аналогично с альбомом и исполнителем
    preview varchar(1024),
    content varchar(1024),
    -- ссылка на объект в s3 хранилище довольно длинная, порядка пяти сотен символов
    play_count int not null default 0
);

create table if not exists album_track (
    id       serial primary key,
    album_id int not null,
    track_id int not null,
    foreign key (album_id) references album (id) on delete cascade ,
    foreign key (track_id) references track (id) on delete cascade
);

create table if not exists artist_track (
    id       serial primary key,
    artist_id int not null,
    track_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade ,
    foreign key (track_id) references track (id) on delete cascade
);

create table if not exists playlist_track (
    id          serial primary key,
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id) on delete cascade ,
    track_id    int not null,
    foreign key (track_id) references track (id) on delete cascade
);

create table if not exists profile_track (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    track_id int not null,
    foreign key (track_id) references track (id) on delete cascade
);

create table if not exists profile_artist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    artist_id int not null,
    foreign key (artist_id) references artist (id) on delete cascade
);

create table if not exists profile_album (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    album_id int not null,
    foreign key (album_id) references album (id) on delete cascade
);

create table if not exists profile_playlist (
    id         serial primary key,
    profile_id uuid not null,
    foreign key (profile_id) references profile (id) on delete cascade ,
    playlist_id int not null,
    foreign key (playlist_id) references playlist (id) on delete cascade
);
