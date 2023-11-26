INSERT INTO artist_genre ( artist_id, genre_id ) VALUES
( ( SELECT id FROM artist WHERE name = 'Eminem' limit 1), ( SELECT id FROM genre WHERE name = 'foreignrap' ) ) on conflict do nothing 
;