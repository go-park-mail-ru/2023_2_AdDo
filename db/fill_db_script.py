import json


def create_artist_command(artist_name, artist_avatar):
    command = (
        f"INSERT INTO artist (name, avatar) "
        f" VALUES ('{artist_name}', '{artist_avatar}');"
    )
    return command


def create_album_command(artist_name, album_name, album_preview, release_album):
    command = (
        f"INSERT INTO album (name, artist_id, preview, release_date)"
        f" VALUES ('{album_name}', (SELECT id FROM artist WHERE name = '{artist_name}') , '{album_preview}', '{release_album}');"
    )
    return command


def create_track_command(track_name, track_preview, track_content, album_name, artist_name):
    command = (
        f"INSERT INTO track (name, preview, content)"
        f" VALUES ('{track_name}', '{track_preview}', '{track_content}');"

        f"INSERT INTO album_track (album_id, track_id)"
        f" VALUES ((SELECT id FROM album WHERE name = '{album_name}'), (SELECT id FROM track WHERE name = '{track_name}'));"

        f"INSERT INTO artist_track (artist_id, track_id)"
        f" VALUES ((SELECT id FROM artist WHERE name = '{artist_name}'), (SELECT id FROM track WHERE name = '{track_name}'));"
    )
    return command


def create_single_command(track_name, track_preview, track_content, artist_name):
    command = (
        f"INSERT INTO track (name, preview, content)"
        f" VALUES ('{track_name}', '{track_preview}', '{track_content}');"

        f"INSERT INTO artist_track (artist_id, track_id)"
        f" VALUES ((SELECT id FROM artist WHERE name = '{artist_name}'), (SELECT id FROM track WHERE name = '{track_name}'));"
    )
    return command


with open("data_for_db_filling.json") as file:
    conf = json.load(file)


with open("fill_data/data_filling.sql", "w") as file:
    for artist in conf["artists"]:
        artist_name = artist["artist_name"]

        artist_avatar_url = "/images/avatars/artists/" + artist_name.replace(" ", "_") + ".jpg"
        file.write(create_artist_command(artist_name, artist_avatar_url) + '\n')

        for album in artist["albums"]:
            album_name = album["album_name"]
            album_release = album["album_release_date"]

            album_image = "/images/tracks/" + artist_name.replace(" ", "_") + "/" + album_name.replace(" ", "_") + ".jpg"
            file.write(create_album_command(artist_name, album_name, album_image, album_release) + '\n')

            tracks = album["tracks"]
            for track in tracks:
                track = track.replace("'", "''")
                track_url = "/audio/" + artist_name.replace(" ", "_") + "/" + album_name.replace(" ", "_") + "/" + track.replace(" ", "_") + ".mp3"
                file.write(create_track_command(track, album_image, track_url, album_name, artist_name) + '\n')

file.close()