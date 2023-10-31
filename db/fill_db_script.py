import json
import sys

from pydub import AudioSegment
import requests
import io


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


def create_track_command(track_name, track_preview, track_content, track_duration, album_name, artist_name):
    command = (
        f"INSERT INTO track (name, preview, content, duration)"
        f" VALUES ('{track_name}', '{track_preview}', '{track_content}', '{track_duration}');"

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


server = "https://api.s3.musicon.space"


def check_url(url):
    response = requests.get(server + url)
    if response.status_code != 200:
        print("Wrong url")
        sys.exit(1)


def get_track_duration(track_url):
    response = requests.get(server + track_url)
    if response.status_code != 200:
        print("Wrong url")
        sys.exit(1)

    audio_data = io.BytesIO(response.content)
    audio_segment = AudioSegment.from_file(audio_data)
    return len(audio_segment) // 1000


with open("data_for_db_filling.json") as file:
    conf = json.load(file)


with open("fill_data/data_filling.sql", "w") as file:
    for artist in conf["artists"]:
        artist_name = artist["artist_name"]
        print("Artist:", artist_name)

        # во всех url пробел заменяется на нижнее подчеркивание
        artist_avatar_url = "/images/avatars/artists/" + artist_name.replace(" ", "_") + ".jpg"
        print("Avatar:", server + artist_avatar_url, "\n")
        check_url(artist_avatar_url)

        # в запросах надо удваивать одинарные кавычки
        file.write(create_artist_command(artist_name.replace("'", "''"),
                                         artist_avatar_url.replace("'", "''")
                                         ) + '\n')

        for album in artist["albums"]:
            album_name = album["album_name"]
            print("Album:", album_name)

            album_release = album["album_release_date"]
            album_image = ("/images/tracks/" + artist_name + "/" + album_name + ".jpg").replace(" ", "_")
            print("Album image:", server + album_image, "\n")
            check_url(album_image)

            file.write(create_album_command(artist_name.replace("'", "''"),
                                            album_name.replace("'", "''"),
                                            album_image.replace("'", "''"),
                                            album_release
                                            ) + '\n')

            for track_name in album["tracks"]:
                track_url = ("/audio/" + artist_name + "/" + album_name + "/" + track_name + ".mp3").replace(" ", "_")

                print("Track:", track_name)
                print("Track url:", server + track_url)

                track_duration = get_track_duration(track_url)
                print("Duration:", track_duration, '\n')

                file.write(create_track_command(track_name.replace("'", "''"),
                                                album_image.replace("'", "''"),
                                                track_url.replace("'", "''"),
                                                track_duration,
                                                album_name.replace("'", "''"),
                                                artist_name.replace("'", "''")
                                                ) + '\n')

file.close()
