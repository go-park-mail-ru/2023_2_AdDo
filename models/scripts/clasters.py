import numpy as np
from sklearn.decomposition import PCA
from sklearn.preprocessing import LabelBinarizer
import psycopg2
from psycopg2.extras import NamedTupleCursor
import json
import matplotlib.pyplot as plt
from sklearn.cluster import KMeans
import pickle

conn = psycopg2.connect('postgresql://musicon:Music0nSecure@82.146.45.164:5433/musicon')
cursor = conn.cursor(cursor_factory=NamedTupleCursor)

cursor.execute(f'select * from genre')
genres_db = cursor.fetchall()
genres = []
for genre in genres_db:
    genres.append(genre[1])

label_binarizer_genre = LabelBinarizer()
genres_encoded = label_binarizer_genre.fit_transform(genres)

cursor.execute(f'select * from artist')
artists_db = cursor.fetchall()
artists = []
for artist in artists_db:
    artists.append(artist[1])

label_binarizer_artist = LabelBinarizer()
artists_encoded = label_binarizer_artist.fit_transform(artists)

cursor.execute(
    f'select track.duration, track.valence, track.arousal, genre.name genre_name, artist.name artist_name, track.id from track join track_genre tg on tg.track_id = track.id join artist_track a on track.id = a.track_id join genre on tg.genre_id = genre.id join artist on a.artist_id = artist.id')
tracks_db = cursor.fetchall()
max_value = 9729
X = np.zeros((1, 2966))
trackIdToVec = {}
for track in tracks_db:
    temp = [[float(track[0]) / max_value, float(track[1]), float(track[2])]]

    g = artists_encoded[genres.index(track[3])]
    a = artists_encoded[artists.index(track[4])]

    temp[0].extend(g)
    temp[0].extend(a)
    temp[0].append(int(track[5]))

    trackIdToVec[int(track[5])] = temp

    X = np.concatenate((X, temp), axis=0)
np.save('cluster_data.npy', X)
#
# with open('trackIdsToVec.pickle', 'rb') as f:
#     playlistToIds = pickle.load(f)
#
# resultPlaylist = {}
# for playlistId in playlistToIds:
#     resultPlaylist[playlistId] = []
#     for trackId in playlistToIds[playlistId]:
#         try:
#             resultPlaylist[playlistId].append(trackIdToVec[trackId])
#         except:
#             print('no this id in clusters')
#
X = np.load('cluster_data.npy', allow_pickle=True)
data_without_id = X[:, :-1]

kmeans = KMeans(n_clusters=40, random_state=0, n_init=30)
clusters = kmeans.fit_predict(data_without_id)

centroids = kmeans.cluster_centers_
labels = kmeans.labels_

data = {
    'data': X.tolist(),
    'centroids': centroids.tolist(),
    'labels': labels.tolist()
}

with open('../../db/cluster_data/clustering_data.json', 'w') as f:
    json.dump(data, f)
