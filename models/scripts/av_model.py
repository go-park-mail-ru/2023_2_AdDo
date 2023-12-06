import librosa
import os
import numpy as np
import tensorflow as tf
import keras
from keras.models import Sequential
from keras.layers import Conv2D, MaxPooling2D, Flatten, Dense, Dropout
from keras.preprocessing.image import ImageDataGenerator
from keras.layers import BatchNormalization
from keras.callbacks import EarlyStopping, ModelCheckpoint, ReduceLROnPlateau
import csv
import glob

sr = 22050
duration = 30
batch_size = 32

audio_dir = 'DEAM_audio/MEMD_audio'
feature_dir = 'librosa_features'

input_dim = 0
def Create_Valence_Model():
    model = Sequential()

    # Добавляем слои для извлечения признаков
    model.add(Dense(64, activation='relu', input_dim=input_dim))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(128, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(256, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(512, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(1024, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(512, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(256, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(128, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(64, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(1, activation='tanh', name='valence'))

    return model

def Create_Arousal_Model():
    model = Sequential()

    model.add(Dense(64, activation='relu', input_dim=input_dim))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(128, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(256, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(512, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(1024, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(512, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(256, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(128, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(64, activation='relu'))
    model.add(BatchNormalization())
    model.add(Dropout(0.25))

    model.add(Dense(1, activation='tanh', name='arousal'))
    return model

# for filename in os.listdir(audio_dir):
#     # Load the audio file
#     y, dummy = librosa.load(os.path.join(audio_dir, filename), sr=sr, duration=duration, mono=True)
#     # Extract the features
#     mfccs = librosa.feature.mfcc(y=y, sr=sr, n_mfcc=13)
#     chroma = librosa.feature.chroma_stft(y=y, sr=sr)
#     mel = librosa.feature.melspectrogram(y=y, sr=sr)
#     contrast = librosa.feature.spectral_contrast(y=y, sr=sr)
#     tonnetz = librosa.feature.tonnetz(y=y, sr=sr)
#     # Save the features to a file
#     np.savez(os.path.join(feature_dir, os.path.splitext(filename)[0] + '.npz'),
#              mfccs=mfccs, chroma=chroma, mel=mel, contrast=contrast, tonnetz=tonnetz)

# получение входных данных для модели
data = {}
flag = False
max_val = 0
y_arousal_train = []
y_valence_train = []
max_valence = 0
max_arousal = 0
with open(
        'DEAM_Annotations/annotations/annotations_averaged_per_song/song_level/static_annotations_averaged_songs_1_2000.csv',
        'r') as file:
    reader = csv.reader(file)
    for row in reader:
        try:
            int(row[0])
        except:
            continue
        # if int(row[0]) == 500:
        #     break
        y_arousal_train.append(float(row[3]))
        max_arousal = max(max_arousal, float(row[3]))
        y_valence_train.append(float(row[1]))
        max_valence = max(max_valence, float(row[1]))

with open(
        'DEAM_Annotations/annotations/annotations_averaged_per_song/song_level/static_annotations_averaged_songs_2000_2058.csv',
        'r') as file:
    reader = csv.reader(file)
    for row in reader:
        try:
            int(row[0])
        except:
            continue
        y_arousal_train.append(float(row[3]))
        y_valence_train.append(float(row[1]))

def custom_sort(file_path):
    file_name = file_path.split('/')[-1]  # Получаем имя файла из пути
    file_name_without_extension = file_name.split('.')[0]  # Удаляем расширение файла, если есть
    try:
        return int(file_name_without_extension)
    except ValueError:
        return 0

file_paths = glob.glob(feature_dir + '/*')
sorted_file_paths = sorted(file_paths, key=custom_sort)

for filename in sorted_file_paths:
    audio_features = np.load(filename)

    concatenated_vector = np.concatenate((audio_features['mfccs'], audio_features['chroma'], audio_features['mel'],
                                          audio_features['contrast'], audio_features['tonnetz']), axis=0)
    tensor = tf.convert_to_tensor(concatenated_vector.flatten())
    input_dim = len(tensor)
    max_val = max(tf.reduce_max(tensor), max_val)

    print(filename)

    if flag is False:
        data[0] = tensor
        data[0] = tf.reshape(data[0], (1, len(data[0])))
        flag = True
    else:
        tensor = tf.reshape(tensor, (1, len(tensor)))
        data[0] = tf.concat([data[0], tensor], axis=0)

y_a_train = []
for element in y_arousal_train:
    y_a_train.append(element / max_arousal)

y_v_train = []
for element in y_valence_train:
    y_v_train.append(element / max_valence)

data = tf.divide(data[0], max_val)

model_v = Create_Valence_Model()
model_v.summary()
model_v.compile(loss='mean_absolute_error', optimizer='adam', metrics=['mean_absolute_error'])
model_v.fit(x=data, y=tf.convert_to_tensor(y_v_train), epochs=18, batch_size=batch_size)
model_v.save('v.keras')

model_a = Create_Valence_Model()
model_a.summary()
model_a.compile(loss='mean_absolute_error', optimizer='adam', metrics=['mean_absolute_error'])
model_a.fit(x=data, y=tf.convert_to_tensor(y_a_train), epochs=18, batch_size=batch_size)
model_a.save('a.keras')
