import time
import pickle
import wget
from yandex_music import Client
import yandex_music
from pathlib import Path
from minio import Minio
from minio.error import S3Error
import requests
import psycopg2
from psycopg2.extras import NamedTupleCursor
from minio import Minio
import os
import tensorflow as tf
import librosa
import numpy as np


# Флоу обучения модели.
# Playlist parsing. We need 200 different playlists, which has less than 50 tracks
# 6 * 50 * 200 = 60000mb = 60gb
# we save playlists into map.
# {
#     playlistId: [
#         {
#             trackId
#             title
#             artist: name, avatar
#             album: cover, avatar
#             genre: title
#             duration
#         }
#     ]
# }
# after that we need to add this tracks in our app. We need download two covers, audio, get av for track
# upload it to minio, check urls, add track to database.
# we need save playlist id to our db tracks id
# after that we can get and fit our model with two different classes: like and skip
# we will take two playlists(may be batch tracks) and outer section in our track sets will be skips. other tracks will be likes
# we decided to take such vec with track features(dur/max_dur, arousal, valence, onehot genre, onehot artist)
# so, after fitting the model in classification service we will take our tracks vectors from candidate_micros(cluster repo)
# and in the end sorting our candidates (1 - like, 0 - skip)

# stop talking, let's code
# playlist_ids = {}
#
# client = Client("y0_AgAAAAAbRHgFAAG8XgAAAADyWaE2jVlJZBQTS2e2hxPG-7ELnuOHHhY").init()
#
# minio_client = Minio(
#     "api.s3.musicon.space",
#     access_key="BeKn8vgv9gmYGrfPkvoc",
#     secret_key="EzWtcUjJUeci8tKVg3SJjetPYK07Zw1N60DQ5IuD",
# )
# texts = ['for', 'для', 'плейлист', 'playlist', 'рок', 'попса', 'электро', 'rock', 'pop', 'electro', 'русский рок', 'new year', 'birthday', 'game', 'rap', 'workout', 'travelling', 'g', 'h', 'energy', 'sad', 'депрессия', 'девочки', 'мальчики', 'boys', 'girls', 'share', 'friends', 'classic', 'favourite', 'любимое', 'избранное']
# counter = 1
# counter_text = 0
# text = texts[counter_text]
# while len(playlist_ids) < 200:
#     if counter > 200:
#         counter = 0
#         counter_text += 1
#         text = texts[counter_text]
#     result = requests.get(f'https://music.yandex.ru/handlers/music-search.jsx?text={text}&type=playlists&page={counter}&ncrnd=0.706970986915534&clientNow=1703285149701&lang=ru&external-domain=music.yandex.ru')
#     try:
#         result.json()['playlists']['items']
#     except:
#         counter += 1
#         continue
#     for playlist in result.json()['playlists']['items']:
#         print(playlist['trackCount'], counter)
#         if int(playlist['trackCount']) < 50 and playlist['duration'] not in playlist_ids:
#             kind = playlist['kind']
#             owner = playlist['owner']['login']
#             response = requests.get(f'https://music.yandex.ru/handlers/playlist.jsx?owner={owner}&kinds={kind}&light=true&madeFor=&lang=ru&external-domain=music.yandex.ru&overembed=false&ncrnd=0.3663209517083549')
#             try:
#                 response.json()
#             except:
#                 counter += 1
#                 continue
#             playlist_ids[response.json()['playlist']['duration']] = response.json()['playlist']['tracks']
#             print(len(playlist_ids), 'playlist')
#     counter += 1
#
# with open('playlistDict.pickle', 'wb') as f:
#     pickle.dump(playlist_ids, f)

with open('playlistDict.pickle', 'rb') as f:
    playlistDict = pickle.load(f)

minio_client = Minio(
    "api.s3.musicon.space",
    access_key="BeKn8vgv9gmYGrfPkvoc",
    secret_key="EzWtcUjJUeci8tKVg3SJjetPYK07Zw1N60DQ5IuD",
)
client = Client("y0_AgAAAAAbRHgFAAG8XgAAAADyWaE2jVlJZBQTS2e2hxPG-7ELnuOHHhY").init()

conn = psycopg2.connect('postgresql://musicon:Music0nSecure@82.146.45.164:5433/musicon')
cursor = conn.cursor(cursor_factory=NamedTupleCursor)

genresDict = {}
artistsDict = {}
albumsDict = {}
trackDict = {}

cursor.execute(f'select * from genre')
genres = cursor.fetchall()
for genre in genres:
    genresDict[genre[0]] = genre[1]

cursor.execute(f'select * from artist')
artists = cursor.fetchall()
for artist in artists:
    artistsDict[artist[0]] = artist[1]

cursor.execute(f'select * from album')
albums = cursor.fetchall()
for album in albums:
    albumsDict[album[0]] = album[1]

cursor.execute(f'select * from track')
tracks = cursor.fetchall()
for track in tracks:
    trackDict[track[0]] = track[1]

server = 'https://api.s3.musicon.space/'
def checkUrl(url):
    r = requests.get(server + url)
    if r.status_code == 200:
        return True
    return False

sample_rate = 22050
duration = 30
model_v = tf.keras.models.load_model('v.keras')
model_a = tf.keras.models.load_model('a.keras')
max_val = 10877.5947265625
offset = 15
emptyLyrics = ''

def calc_av(track_path):
    try:
        y, dummy = librosa.load(track_path, sr=sample_rate, duration=duration, mono=True, offset=offset)

        mfccs = librosa.feature.mfcc(y=y, sr=sample_rate, n_mfcc=13)
        chroma = librosa.feature.chroma_stft(y=y, sr=sample_rate)
        mel = librosa.feature.melspectrogram(y=y, sr=sample_rate)
        contrast = librosa.feature.spectral_contrast(y=y, sr=sample_rate)
        tonnetz = librosa.feature.tonnetz(y=y, sr=sample_rate)

        concatenated_vector = np.concatenate((mfccs, chroma, mel, contrast, tonnetz), axis=0)
        data = tf.convert_to_tensor(concatenated_vector.flatten())
        data = tf.reshape(data, (1, len(data)))

        data = tf.divide(data, max_val)
        valence = model_v.predict(data)
        arousal = model_a.predict(data)
        return valence[0][0], arousal[0][0]
    except:
        print('error with valence')
        return 0, 0

def uploadArtistAvatar(a):
    wget.download('https://' + a['cover']['uri'].replace('%%', '400x400'), out='artistTemp.webp')
    aV = ('artists/' + a['name'] + '.webp').replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        minio_client.fput_object(
            "images", aV, 'artistTemp.webp',
        )
    except:
        print('error artist uploading file')
    try:
        os.remove('artistTemp.webp')
    except:
        print('no such artist file')

    return '/images/' + aV

def uploadAlbumCover(a):
    wget.download('https://' + a['coverUri'].replace('%%', '400x400'), out='albumTemp.webp')
    aC = ('albums/' + a['title'] + '.webp').replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        minio_client.fput_object(
            "images", aC, 'albumTemp.webp',
        )
    except:
        print('error album uploading file')
    try:
        os.remove('albumTemp.webp')
    except:
        print('no such album file')

    return '/images/' + aC

def uploadTrackContent(a):
    client.tracks([a['id']])[0].download(filename='track.mp3')
    tC = ('tracks/' + a['albums'][0]['title'] + a['title'] + '.mp3').replace(' ', '_').replace('\'', '_').replace('\"', '_')
    try:
        minio_client.fput_object(
            "audio", tC, 'track.mp3',
        )
    except:
        print('error track uploading file')
    a, v = calc_av('track.mp3')
    try:
        os.remove('track.mp3')
    except:
        print('no such album file')

    return '/audio/' + tC, a, v
# I think we will ignore genres in playlists
playlistFromYaToOurIds = {}

def find_key_by_value(my_map, value):
    for key, val in my_map.items():
        if val == value:
            return key
    return None

for playlistId in playlistDict:
    try:
        playlistTracks = playlistDict[playlistId]
        playlistFromYaToOurIds[playlistId] = []
        for track in playlistTracks:
            newGenre = track['albums'][0]['genre']
            newGenreId = 0
            if newGenre not in genresDict.values():
                cursor.execute('insert into genre (name) values (%s) on conflict do nothing returning id;', (newGenre, ))
                newGenreId = cursor.fetchone()[0]
                genresDict[newGenreId] = newGenre
            newGenreId = find_key_by_value(genresDict, newGenre)

            newArtist = track['artists'][0]
            newArtistId = 0
            if newArtist['name'] not in artistsDict.values():
                artistAvatar = uploadArtistAvatar(newArtist)
                cursor.execute(f'insert into artist (name, avatar) values (%s, %s) returning id;', (newArtist['name'], artistAvatar, ))
                newArtistId = cursor.fetchone()[0]
                artistsDict[newArtistId] = newArtist['name']
            newArtistId = find_key_by_value(artistsDict, newArtist['name'])

            newAlbum = track['albums'][0]
            newAlbumId = 0
            if newAlbum['title'] not in albumsDict.values():
                albumCover = uploadAlbumCover(newAlbum)
                cursor.execute(f'insert into album (name, preview, release_date, year) values (%s, %s, %s, %s) returning id ;', (newAlbum['title'], albumCover, newAlbum['releaseDate'], newAlbum['year']))
                newAlbumId = cursor.fetchone()[0]
                albumsDict[newAlbumId] = newAlbum['title']
            newAlbumId = find_key_by_value(albumsDict, newAlbum['title'])

            if track['title'] not in trackDict.values():
                trackContent, a, v = uploadTrackContent(track)
                cursor.execute(f'insert into track (name, preview, content, duration, valence, arousal, lyrics) values (%s, %s, %s, %s, %s, %s, %s) returning id;', (track['title'], albumCover, trackContent, track['durationMs'] / 1000, float(v), float(a), emptyLyrics))
                newTrackId = cursor.fetchone()[0]
                cursor.execute(f'insert into artist_track (artist_id, track_id) values (%s, %s) on conflict do nothing ;', (newArtistId, newTrackId))
                cursor.execute(f'insert into track_genre (genre_id, track_id) values (%s, %s) on conflict do nothing;', (newGenreId, newTrackId))
                cursor.execute(f'insert into album_track (album_id, track_id) values (%s, %s) on conflict do nothing;', (newAlbumId, newTrackId))
                cursor.execute(f'insert into artist_album (artist_id, album_id) values (%s, %s) on conflict do nothing;', (newArtistId, newAlbumId))
                trackDict[newTrackId] = track['title']
                playlistFromYaToOurIds[playlistId].append(newTrackId)
                conn.commit()
    except:
        conn.rollback()
        print('error while all executing cycle')

cursor.close()
conn.close()
with open('playlistYaToOurIds.pickle', 'wb') as f:
    pickle.dump(playlistFromYaToOurIds, f)