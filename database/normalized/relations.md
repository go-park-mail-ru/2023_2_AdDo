Отношение для профилей пользователей

Relation profile:

{id} -> email, password, nickname, birth_date, avatar_url


Отношение для хранения информации об исполнителях

Relation artist:

{id} -> name, avatar


Отношение для альбомов, где превью - картинка альбома

Relation album:

{id} -> name, preview, release_date, artist_id

{artist_id} -> artist (id)


Отношение для трэков, где превью - картинка трэка

Relation track:

{id} -> name, preview, content, play_count

Отношение для плейлистов, где creator_id - айдишник создателя

Relation playlist:

{id} -> name, preview, creating_date, creator_id

{creator_id} -> profile(id)


Отношение многие ко многим для альбомов к трэкам 

Relation album_track

{id} -> album_id, track_id

{album_id} -> album (id)

{track_id} -> track (id)


Отношение многие ко многим для исполнителей к трэкам

Relation artist_track

{id} -> artist_id, track_id

{artist_id} -> artist (id)

{track_id} -> track (id)


Отношение многие ко многим для плейлистов к трэкам

Relation playlist_track

{id} -> playlist_id, track_id

{playlist_id} -> playlist (id)

{track_id} -> track (id)


Отношение многие ко многим для пользователей к понравившимся релизам

Relation profile_track

{id} -> profile_id, track_id

{profile_id} -> profile (id)

{track_id} -> track (id)

Отношение многие ко многим для пользователей к понравившимся исполнителям

Relation profile_artist

{id} -> profile_id, artist_id

{profile_id} -> profile (id)

{artist_id} -> artist (id)


Отношение многие ко многим для пользователей к понравившимся альбомам

Relation profile_album

{id} -> profile_id, album_id

{album_id} -> album (id)

{profile_id} -> profile (id)


Отношение многие ко многим для пользователей к понравившимся плейлистам ДРУГИХ пользователей

Relation profile_playlist

{id} -> profile_id, playlist_id

{profile_id} -> profile (id)

{playlist_id} -> playlist (id)


Идентификатор (id) является первичным ключом и функционально определяет все остальные поля во всех таблицах. В некоторых таблицах определены внешние ключи, которые связаны с первичными ключами в других таблицах, это означает, что внешний ключ функционально зависит от первичного ключа. Следовательно, функциональные зависимости описывают связь между атрибутами отношений и позволяют определять значения полей исходя из первичного ключа.