INSERT INTO artist_genre ( artist_id, genre_id ) VALUES
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM genre WHERE name = 'foreignrap' ) )
;