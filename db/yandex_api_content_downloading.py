import time

import wget
from yandex_music import Client
import yandex_music
from pathlib import Path
from minio import Minio
from minio.error import S3Error

client = Client("y0_AgAAAAAbRHgFAAG8XgAAAADyWaE2jVlJZBQTS2e2hxPG-7ELnuOHHhY").init()

minio_client = Minio(
    "api.s3.musicon.space",
    access_key="BeKn8vgv9gmYGrfPkvoc",
    secret_key="EzWtcUjJUeci8tKVg3SJjetPYK07Zw1N60DQ5IuD",
)

MAX_ARTIST_BATCH = 10000

MINIO_PATH = '/data/minio'

total_errors = 0


def form_request(start):
    result = []
    for j in range(start, start + MAX_ARTIST_BATCH):
        result.append(j)
    return result


def download_artist_image(a):
    artist_avatar_path = "/images/artists/" + a.name.replace(" ", "_") + ".webp"

    minio_artist_path = ('artists/' + a.name.replace(" ", "_") + ".webp").replace(' ', '_').replace('\'', '_').replace(
        '\"', '_')
    Path(MINIO_PATH + "/images/artists/" + a.name.replace(' ', '_').replace('\'', '_').replace('\"', '_')).mkdir(
        parents=True, exist_ok=True)
    Path(MINIO_PATH + "/images/tracks/" + a.name.replace(' ', '_').replace('\'', '_').replace('\"', '_')).mkdir(
        parents=True, exist_ok=True)

    filename = (MINIO_PATH + artist_avatar_path).replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        a.cover.download(filename=filename, size='400x400')

        minio_client.fput_object(
            "images", minio_artist_path, filename,
        )
    except:
        print('error artist file')


def download_album_preview(artist_name, a):
    album_preview_path = ("/images/tracks/" + artist_name + "/" + a.title + ".webp").replace(' ', '_').replace('\'',
                                                                                                               '_').replace(
        '\"', '_')
    minio_album_path = ('tracks/' + artist_name + "/" + a.title + ".webp").replace(' ', '_').replace('\'', '_').replace(
        '\"', '_')
    path_for_audio = (MINIO_PATH + "/audio/" + artist_name + '/' + a.title).replace(' ', '_').replace('\'',
                                                                                                      '_').replace('\"',
                                                                                                                   '_')
    Path(path_for_audio).mkdir(parents=True, exist_ok=True)

    filename = (MINIO_PATH + album_preview_path).replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        a.download_cover(filename=filename, size='400x400')
        minio_client.fput_object(
            "images", minio_album_path, filename,
        )
    except:
        'error album cover'


def download_track_content(artist_name, album_title, t):
    track_content_avatar_path = ("/audio/" + artist_name + "/" + album_title + "/" + t.title + ".mp3").replace(' ',
                                                                                                               '_').replace(
        '\'', '_').replace('\"', '_')
    minio_track_path = (artist_name + "/" + album_title + "/" + t.title + ".mp3").replace(' ', '_').replace('\'',
                                                                                                            '_').replace(
        '\"', '_')
    filename = (MINIO_PATH + track_content_avatar_path).replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        t.download(filename=filename)
        minio_client.fput_object(
            "audio", minio_track_path, filename,
        )
    except:
        print('error music file')


for i in range(229000, 1000000, MAX_ARTIST_BATCH):
    ids = form_request(i)
    print(f'{i}\n')
    try:
        artists = client.artists(ids)
        for artist in artists:
            try:
                if artist.error is None and artist.name is not None and artist.ratings is not None and artist.cover is not None and artist.ratings.month < 200:
                    artist.name = artist.name.replace("'", "''").replace('"', '\"')
                    download_artist_image(artist)

                    albums = artist.getAlbums()
                    for album in albums:
                        try:
                            album = album.withTracks()
                            album.title = album.title.replace("'", "''").replace('"', '\"')

                            download_album_preview(artist.name, album)

                            for volume in album.volumes:
                                for track in volume:
                                    track.title = track.title.replace("'", "''").replace('"', '\"')
                                    download_track_content(artist.name, album.title, track)
                        except:
                            print('error albums')
            except:
                print('error wrong artist')
                print(i)
    except:
        print('error batch')
        time.sleep(2)