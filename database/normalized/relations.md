Relation profile:

{id} -> email, password, nickname, birth_date, avatar_url

{email} -> password, nickname, birth_date, avatar_url


Relation artist:

{id} -> name, avatar


Relation album:

{id} -> name, preview, release_date, artist_id

{artist_id} -> artist (id)


Relation track:

{id} -> name, preview, content, play_count


Relation playlist:

{id} -> name, preview, creating_date, creator_id

{creator_id} -> profile(id)


Relation album_track

{id} -> album_id, track_id

{album_id} -> album (id)

{track_id} -> track (id)


Relation artist_track

{id} -> artist_id, track_id

{artist_id} -> artist (id)

{track_id} -> track (id)


Relation playlist_track

{id} -> playlist_id, track_id

{playlist_id} -> playlist (id)

{track_id} -> track (id)


Relation profile_track

{id} -> profile_id, track_id

{profile_id} -> profile (id)

{track_id} -> track (id)


Relation profile_artist

{id} -> profile_id, artist_id

{profile_id} -> profile (id)

{artist_id} -> artist (id)


Relation profile_album

{id} -> profile_id, album_id

{album_id} -> album (id)

{profile_id} -> profile (id)


Relation profile_playlist

{id} -> profile_id, playlist_id

{profile_id} -> profile (id)

{playlist_id} -> playlist (id)
