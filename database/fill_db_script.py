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


with open("data_for_db_filling.json") as file:
    conf = json.load(file)


server = conf["server_addr"]
artist_name = conf["artist_name"]
album_name = conf["album_name"]
album_release = conf["album_release_date"]
track_name = conf["track_name"]

artist_avatar_url = server + "/images/avatars/artists/" + artist_name + ".jpg"
album_url = server + "/audio/" + artist_name + "/albums/" + album_name
track_url = server + "/audio/" + artist_name + "/albums/" + album_name + "/" + track_name

print(artist_avatar_url, "\n", album_url, "\n", track_url)


with open("data_init.sql", "w") as file:
    file.write(create_artist_command(artist_name, artist_avatar_url) + '\n')
    file.write(create_album_command(artist_name, album_name, album_url, album_release) + '\n')
    file.write(create_track_command(track_name, track_url, "music url here", album_name, artist_name) + '\n')

file.close()
