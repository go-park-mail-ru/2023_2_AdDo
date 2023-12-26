import librosa
import os
import numpy as np
import tensorflow as tf
import keras
from keras.models import Sequential
from keras.layers import Conv2D, MaxPooling2D, Flatten, Dense, Dropout, LSTM, BatchNormalization
from keras.preprocessing.image import ImageDataGenerator
from keras.layers import BatchNormalization
from keras.callbacks import EarlyStopping, ModelCheckpoint, ReduceLROnPlateau
import csv
import glob
import pickle


# playlistToIds = {}
# with open('trackIdsToVec.pickle', 'rb') as f:
#     playlistToIds = pickle.load(f)
#
# len_vec = 0
#
# for key, val in playlistToIds.items():
#     len_vec = len(val[0][0])
#     break
# end_token = np.zeros(len_vec)
# tensor = tf.convert_to_tensor(end_token)
# tensor = tf.reshape(tensor, [1, len_vec])
# likes = 0
# skips = 0
# for key, val in playlistToIds.items():
#     for v in playlistToIds[key]:
#         temp = np.array(v)
#         temp = np.reshape(temp, (1, len_vec))
#         tensor = tf.concat([tensor, temp], axis=0)
#         likes += 1
#         if likes == 5:
#             likes = 0
#             for key2, val2 in playlistToIds.items():
#                 if key2 == key:
#                     continue
#                 for v2 in playlistToIds[key2]:
#                     temp = np.array(v2)
#                     temp = np.reshape(temp, (1, len_vec))
#                     tensor = tf.concat([tensor, temp], axis=0)
#                     skips += 1
#                     if skips == 5:
#                         skips = 0
#                         end_token = np.reshape(end_token, (1, len_vec))
#                         tensor = tf.concat([tensor, end_token], axis=0)
#                         break_outer_loop = True
#                         break
#                 if 'break_outer_loop' in locals():
#                     break_outer_loop = False
#                     break
#
# with open('tensor.picle', 'wb') as file:
#     pickle.dump(tensor, file)

BATCH_SIZE = 64
EPOCHS = 5

with open('tensor.picle', 'rb') as f:
    train_x = pickle.load(f)

train_x = train_x[:, :-1]
tensor_size = 5263
print(train_x)
print(tf.reduce_max(train_x), 'max_x')
train_y = np.empty((5263, 1), dtype=float)  # Создаем пустой массив размерностью (5322, 2)

# Устанавливаем первое значение
train_y[0] = [0]

# Заполняем остальные значения согласно заданному шаблону чередующихся значений
for i in range(1, 5263):
    if i % 11 <= 5:
        train_y[i] = [1]
    else:
        train_y[i] = [0]

# Устанавливаем последнее значение
train_y[-1] = [0]
print(train_y)
print(tf.reduce_max(train_y), 'max_y')

input_dim = 1
model = Sequential()
model.add(LSTM(128, input_dim=input_dim, return_sequences=True))
model.add(Dropout(0.2))
model.add(BatchNormalization())

model.add(LSTM(128, input_dim=input_dim, return_sequences=True))
model.add(Dropout(0.1))
model.add(BatchNormalization())

model.add(LSTM(128, input_dim=input_dim, return_sequences=True))
model.add(Dropout(0.2))
model.add(BatchNormalization())

model.add(Dense(32, activation="relu"))
model.add(Dropout(0.2))

model.add(Dense(1, activation="softmax"))

model.compile(loss='binary_crossentropy', optimizer='adam', metrics=['accuracy'])
model.summary()

model.fit(x=train_x, y=train_y, batch_size=BATCH_SIZE, epochs=EPOCHS)
model.save('classifier.keras')
