INSERT INTO artist_album ( artist_id, album_id ) VALUES
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Gospel' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Last One Standing' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Killer' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Music To Be Murdered By - Side B' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'The Adventures Of Moon Man & Slim Shady' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Music To Be Murdered By' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Bang' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Killshot' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Kamikaze' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Nowhere Fast' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Chloraseptic' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Revival' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Campaign Speech' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Detroit Vs. Everybody' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Twerk Dat Pop That (feat. Eminem & Royce da 5''9")' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Twerk Dat Pop That (Clean) [feat. Eminem & Royce da 5''9"]' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'The Marshall Mathers LP2' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'My Life' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'I Need A Doctor' ) ),
( ( SELECT id FROM artist WHERE name = 'Eminem' ), ( SELECT id FROM album WHERE name = 'Recovery' ) )
;