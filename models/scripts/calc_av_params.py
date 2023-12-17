import wget
import psycopg2
from psycopg2.extras import NamedTupleCursor
import os
import tensorflow as tf
import librosa
import numpy as np

server = "https://api.s3.musicon.space"

sample_rate = 22050
duration = 30
track_path = 'track.mp3'
model_v = tf.keras.models.load_model('v.keras')
model_a = tf.keras.models.load_model('a.keras')
max_val = 10877.5947265625
offset = 15


def calc_av(track):
    try:
        wget.download(server + track[3], track_path)
        y, dummy = librosa.load(track_path, sr=sample_rate, duration=duration, mono=True, offset=10)

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
        print('error')
        return 0, 0

file = open('arous_valence_part2.sql', 'w')

conn = psycopg2.connect('postgresql://musicon:Music0nSecure@82.146.45.164:5433/musicon')
cursor = conn.cursor(cursor_factory=NamedTupleCursor)
cursor.execute(f'select * from track')
tracks = cursor.fetchall()
flag = False
for track in tracks:
    if track[0] != 1828 and flag is not True:
        continue
    flag = True
    try:
        os.remove(track_path)
    except:
        print('no such file')
    v, a = calc_av(track)
    file.write(f'update track set valence = {v}, arousal = {a} where id = {track[0]};\n')


file.close()
