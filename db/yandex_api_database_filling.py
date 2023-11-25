from yandex_music import Client
import yandex_music

client = Client("y0_AgAAAAAbRHgFAAG8XgAAAADyWaE2jVlJZBQTS2e2hxPG-7ELnuOHHhY").init()

data_artist = open("fill_data/data_artist.sql", "w")
data_album = open("fill_data/data_album.sql", "w")
data_track = open("fill_data/data_track.sql", "w")
data_genre = open("fill_data/data_genre.sql", "w")

data_artist_album = open("fill_data/data_artist_album.sql", "w")
data_artist_track = open("fill_data/data_artist_track.sql", "w")
data_album_track = open("fill_data/data_album_track.sql", "w")

data_artist_genre = open("fill_data/data_artist_genre.sql", "w")
data_album_genre = open("fill_data/data_album_genre.sql", "w")
data_track_genre = open("fill_data/data_track_genre.sql", "w")

artist_comma_flag = False
album_comma_flag = False
track_comma_flag = False

album_artist_comma_flag = False
track_artist_comma_flag = False
track_album_comma_flag = False

artist_genre_comma_flag = False
album_genre_comma_flag = False
track_genre_comma_flag = False

data_artist.write(f'INSERT INTO artist ( name, avatar ) VALUES\n')
data_album.write(f'INSERT INTO album ( name, preview, release_date, year ) VALUES\n')
data_track.write(f'INSERT INTO track ( name, preview, duration, content, lyrics, duration ) VALUES\n')

data_artist_album.write(f'INSERT INTO artist_album ( artist_id, album_id ) VALUES\n')
data_artist_track.write(f'INSERT INTO artist_track ( artist_id, track_id ) VALUES\n')

data_album_track.write(f'INSERT INTO album_track ( album_id, track_id ) VALUES\n')

data_artist_genre.write(f'INSERT INTO artist_genre ( artist_id, genre_id ) VALUES\n')
data_album_genre.write(f'INSERT INTO album_genre ( album_id, genre_id ) VALUES\n')
data_track_genre.write(f'INSERT INTO track_genre ( track_id, genre_id ) VALUES\n')

MAX_ARTIST_BATCH = 3000


def form_request(start):
    result = []
    for j in range(start, start + MAX_ARTIST_BATCH):
        result.append(j)
    return result


def add_file_del(file):
    file.write(',\n')


def add_artist_genre_record(artist_name, genres):
    global artist_genre_comma_flag
    for genre in genres:
        if artist_genre_comma_flag is False:
            artist_genre_comma_flag = True
        else:
            add_file_del(data_artist_genre)
        data_artist_genre.write(
            f'( ( SELECT id FROM artist WHERE name = \'{artist_name}\' ), ( SELECT id FROM genre WHERE name = \'{genre}\' ) )')


def add_album_genre_record(album_name, genre):
    global album_genre_comma_flag
    if album_genre_comma_flag is False:
        album_genre_comma_flag = True
    else:
        add_file_del(data_album_genre)
    data_album_genre.write(
        f'( ( SELECT id FROM album WHERE name = \'{album_name}\' ), ( SELECT id FROM genre WHERE name = \'{genre}\' ) )')


def add_track_genre_record(track_title, genre):
    global track_genre_comma_flag
    if track_genre_comma_flag is False:
        track_genre_comma_flag = True
    else:
        add_file_del(data_track_genre)
    data_track_genre.write(
        f'( ( SELECT id FROM track WHERE name = \'{track_title}\' ), ( SELECT id FROM genre WHERE name = \'{genre}\' ) )')


def add_album_artist_record(album_name, artist_name):
    global album_artist_comma_flag
    if album_artist_comma_flag is False:
        album_artist_comma_flag = True
    else:
        add_file_del(data_artist_album)
    data_artist_album.write(
        f'( ( SELECT id FROM artist WHERE name = \'{artist_name}\' ), ( SELECT id FROM album WHERE name = \'{album_name}\' ) )')


def add_genre_record(genres):
    for genre in genres:
        data_genre.write(f'INSERT INTO genre ( name ) VALUES ( \'{genre}\' );\n')


def add_artist_record(artist_name):
    global artist_comma_flag
    if artist_comma_flag is False:
        artist_comma_flag = True
    else:
        add_file_del(data_artist)
    artist_avatar_url = "/images/artists/" + artist_name.replace(" ", "_") + ".webp"
    data_artist.write(f'( \'{artist_name}\', \'{artist_avatar_url}\' )')


def add_album_record(artist_name, title, genre, date, year):
    global album_comma_flag
    if album_comma_flag is False:
        album_comma_flag = True
    else:
        add_file_del(data_album)
    album_preview_url = ("/images/tracks/" + artist_name + "/" + title + ".webp").replace(" ", "_")
    data_album.write(f'( \'{title}\', \'{album_preview_url}\', \'{date}\', \'{year}\' )')


def add_track_record(artist_name, album_title, track_title, text, dur_ms):
    global track_comma_flag
    if track_comma_flag is False:
        track_comma_flag = True
    else:
        add_file_del(data_track)
    url = ("/images/tracks/" + artist_name + "/" + album_title + ".webp").replace(" ", "_")
    track_url = ("/audio/" + artist_name + "/" + album_title + "/" + track_title + ".mp3").replace(" ", "_")
    data_track.write(f'( \'{track_title}\', \'{url}\', \'{track_url}\', \'{text}\', {dur_ms / 1000} )')


def add_track_album_record(track_title, album_title):
    global track_album_comma_flag
    if track_album_comma_flag is False:
        track_album_comma_flag = True
    else:
        add_file_del(data_album_track)
    data_album_track.write(
        f'( ( SELECT id FROM album WHERE name = \'{album_title}\' ), ( SELECT id FROM track WHERE name = \'{track_title}\' ) )')


def add_track_artist_record(track_title, artist_name):
    global track_artist_comma_flag
    if track_artist_comma_flag is False:
        track_artist_comma_flag = True
    else:
        add_file_del(data_artist_track)
    data_artist_track.write(
        f'( ( SELECT id FROM artist WHERE name = \'{artist_name}\' ), ( SELECT id FROM track WHERE name = \'{track_title}\' ) )')


for i in range(1, 6000, MAX_ARTIST_BATCH):
    ids = form_request(i)
    print(f'{i}\n')
    artists = client.artists(ids)
    for artist in artists:
        if artist.error is None and artist.name is not None and artist.ratings is not None and artist.ratings.month < 100:
            artist.name = artist.name.replace("'", "''").replace('"', '\"')

            add_artist_record(artist.name)

            add_genre_record(artist.genres)
            add_artist_genre_record(artist.name, artist.genres)

            albums = artist.getAlbums()
            for album in albums:
                album = album.withTracks()
                album.title = album.title.replace("'", "''").replace('"', '\"')

                add_album_record(artist.name, album.title, album.genre, album.release_date, album.year)
                add_album_genre_record(album.title, album.genre)
                add_album_artist_record(album.title, artist.name)

                for volume in album.volumes:
                    for track in volume:
                        track.title = track.title.replace("'", "''").replace('"', '\"')
                        lyrics = ''
                        lyrics_text = ''
                        try:
                            lyrics = track.get_lyrics('LRC')
                            lyrics_text = lyrics.fetchLyrics()
                        except yandex_music.exceptions.NotFoundError:
                            print()
                        except yandex_music.exceptions.TimedOutError:
                            print()

                        lyrics_text = lyrics_text.replace("'", "''").replace('"', '\"')
                        add_track_record(artist.name, album.title, track.title, lyrics_text,
                                         track.duration_ms)
                        add_track_album_record(track.title, album.title)
                        add_track_artist_record(track.title, artist.name)
                        add_track_genre_record(track.title, album.genre)

data_artist.write('\n;')
data_album.write('\n;')
data_track.write('\n;')

data_artist_album.write('\n;')
data_artist_track.write('\n;')
data_album_track.write('\n;')

data_artist_genre.write('\n;')
data_album_genre.write('\n;')
data_track_genre.write('\n;')

data_artist.close()
data_album.close()
data_track.close()
data_genre.close()

data_artist_album.close()
data_artist_track.close()
data_album_track.close()

data_artist_genre.close()
data_album_genre.close()
data_track_genre.close()

# for i in range(23000000, 23000010):
#     album = client.albums_with_tracks(i)
#     tracks = []
#     for i, volume in enumerate(album.volumes):
#         if len(album.volumes) > 1:
#             tracks.append(f'💿 Диск {i + 1}')
#         tracks += volume
#
#     text = 'АЛЬБОМ\n'
#     text += f'{album.title}\n'
#     text += f"Исполнитель: {', '.join([artist.name for artist in album.artists])}\n"
#     text += f'{album.year} · {album.genre}\n'
#
#     cover = album.cover_uri
#     if cover:
#         text += f'Обложка: {cover.replace("%%", "400x400")}\n\n'
#
#     text += 'Список треков:'
#
#     print(f'{text}\n')
#
#     for track in tracks:
#         if isinstance(track, str):
#             print(track)
#         else:
#             artists = ''
#             if track.artists:
#                 artists = ' - ' + ', '.join(artist.name for artist in track.artists)
#             print(track.title + artists)
#             print()
#
# CHART_ID = 'world'

# chart = client.chart(CHART_ID).chart
#
# text = [f'🏆 {chart.title}', chart.description, '', 'Треки:']
#
# i = 1
# for track_short in chart.tracks:
#     track, chart = track_short.track, track_short.chart
#     artists = ''
#     if track.artists:
#         artists = ' - ' + ', '.join(artist.name for artist in track.artists)
#
#     track_text = f'{track.title}{artists}'
#     info = track.get_download_info(get_direct_links=True)
#     print(info)
#
#     wget.download(info[0]['direct_link'], f'track {i}')
#     print()
#     i += 1
#
#     if chart.progress == 'down':
#         track_text = '🔻 ' + track_text
#     elif chart.progress == 'up':
#         track_text = '🔺 ' + track_text
#     elif chart.progress == 'new':
#         track_text = '🆕 ' + track_text
#     elif chart.position == 1:
#         track_text = '👑 ' + track_text
#
#     track_text = f'{chart.position} {track_text}'
#     text.append(track_text)
#
# print('\n'.join(text))
