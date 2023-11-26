import wget
from yandex_music import Client
import yandex_music
from pathlib import Path
client = Client("y0_AgAAAAAbRHgFAAG8XgAAAADyWaE2jVlJZBQTS2e2hxPG-7ELnuOHHhY").init()

MAX_ARTIST_BATCH = 6000

MINIO_PATH = '/data/minio'


def form_request(start):
    result = []
    for j in range(start, start + MAX_ARTIST_BATCH):
        result.append(j)
    return result


def download_artist_image(a):
    artist_avatar_path = "/images/artists/" + a.name.replace(" ", "_") + ".webp"

    Path(MINIO_PATH + "/images/artists/" + a.name.replace(" ", "_")).mkdir(parents=True, exist_ok=True)
    Path(MINIO_PATH + "/images/tracks/" + a.name.replace(" ", "_")).mkdir(parents=True, exist_ok=True)

    a.cover.download(filename=MINIO_PATH + artist_avatar_path, size='400x400')


def download_album_preview(artist_name, a):
    album_preview_path = ("/images/tracks/" + artist_name + "/" + a.title + ".webp").replace(" ", "_")

    path_for_audio = MINIO_PATH + "/audio/" + artist_name + '/' + a.title
    path_for_audio = path_for_audio.replace(" ", "_")
    Path(path_for_audio).mkdir(parents=True, exist_ok=True)

    a.download_cover(filename=MINIO_PATH + album_preview_path, size='400x400')


def download_track_content(artist_name, album_title, t):
    track_content_avatar_path = ("/audio/" + artist_name + "/" + album_title + "/" + t.title + ".mp3").replace(" ", "_")
    t.download(filename=MINIO_PATH + track_content_avatar_path)


for i in range(1, 12000, MAX_ARTIST_BATCH):
    ids = form_request(i)
    print(f'{i}\n')
    artists = client.artists(ids)
    for artist in artists:
        if artist.error is None and artist.name is not None and artist.ratings is not None and artist.cover is not None and artist.ratings.month < 100:
            artist.name = artist.name.replace("'", "''").replace('"', '\"')
            download_artist_image(artist)

            albums = artist.getAlbums()
            for album in albums:
                album = album.withTracks()
                album.title = album.title.replace("'", "''").replace('"', '\"')

                download_album_preview(artist.name, album)

                for volume in album.volumes:
                    for track in volume:
                        track.title = track.title.replace("'", "''").replace('"', '\"')
                        download_track_content(artist.name, album.title, track)
