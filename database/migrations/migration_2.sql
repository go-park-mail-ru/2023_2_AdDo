INSERT INTO artist (name, avatar)  VALUES ('Doja Cat', '/images/avatars/artists/Doja_Cat.jpg');
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Scarlet', (SELECT id FROM artist WHERE name = 'Doja Cat') , '/images/tracks/Doja_Cat/Scarlet.jpg', '2023-09-22');
INSERT INTO track (name, preview, content) VALUES ('97', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/97.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = '97'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = '97'));
INSERT INTO track (name, preview, content) VALUES ('Agora Hills', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Agora_Hills.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Agora Hills'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Agora Hills'));
INSERT INTO track (name, preview, content) VALUES ('Attention', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Attention.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Attention'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Attention'));
INSERT INTO track (name, preview, content) VALUES ('Balut', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Balut.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Balut'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Balut'));
INSERT INTO track (name, preview, content) VALUES ('Can''t Wait.mp3', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Can''t_Wait.mp3.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Can''t Wait.mp3'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Can''t Wait.mp3'));
INSERT INTO track (name, preview, content) VALUES ('Demons', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Demons.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Demons'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Demons'));
INSERT INTO track (name, preview, content) VALUES ('Fuck The Girls', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Fuck_The_Girls.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Fuck The Girls'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Fuck The Girls'));
INSERT INTO track (name, preview, content) VALUES ('Go Off', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Go_Off.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Go Off'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Go Off'));
INSERT INTO track (name, preview, content) VALUES ('Gun', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Gun.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Gun'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Gun'));
INSERT INTO track (name, preview, content) VALUES ('Love Life', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Love_Life.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Love Life'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Love Life'));
INSERT INTO track (name, preview, content) VALUES ('Often', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Often.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Often'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Often'));
INSERT INTO track (name, preview, content) VALUES ('Ouchies', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Ouchies.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Ouchies'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Ouchies'));
INSERT INTO track (name, preview, content) VALUES ('Paint The Town Red', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Paint_The_Town_Red.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Paint The Town Red'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Paint The Town Red'));
INSERT INTO track (name, preview, content) VALUES ('Shutcho', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Shutcho.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Shutcho'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Shutcho'));
INSERT INTO track (name, preview, content) VALUES ('Skull And Bones', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Skull_And_Bones.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Skull And Bones'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Skull And Bones'));
INSERT INTO track (name, preview, content) VALUES ('Wym Freestyle', '/images/tracks/Doja_Cat/Scarlet.jpg', '/audio/Doja_Cat/Scarlet/Wym_Freestyle.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Scarlet'), (SELECT id FROM track WHERE name = 'Wym Freestyle'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Doja Cat'), (SELECT id FROM track WHERE name = 'Wym Freestyle'));
INSERT INTO artist (name, avatar)  VALUES ('Drake', '/images/avatars/artists/Drake.jpg');
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Dark Lane Demo Tapes', (SELECT id FROM artist WHERE name = 'Drake') , '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '2020-05-01');
INSERT INTO track (name, preview, content) VALUES ('Chicago Freestyle', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Chicago_Freestyle.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Chicago Freestyle'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Chicago Freestyle'));
INSERT INTO track (name, preview, content) VALUES ('D4L (feat. Young Thug, Future)', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/D4L_(feat._Young_Thug,_Future).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'D4L (feat. Young Thug, Future)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'D4L (feat. Young Thug, Future)'));
INSERT INTO track (name, preview, content) VALUES ('Deep Pockets', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Deep_Pockets.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Deep Pockets'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Deep Pockets'));
INSERT INTO track (name, preview, content) VALUES ('Demons (feat. Fivio Foreign)', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Demons_(feat._Fivio_Foreign).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Demons (feat. Fivio Foreign)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Demons (feat. Fivio Foreign)'));
INSERT INTO track (name, preview, content) VALUES ('Desires (feat. Future)', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Desires_(feat._Future).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Desires (feat. Future)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Desires (feat. Future)'));
INSERT INTO track (name, preview, content) VALUES ('From Florida With Love', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/From_Florida_With_Love.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'From Florida With Love'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'From Florida With Love'));
INSERT INTO track (name, preview, content) VALUES ('Landed', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Landed.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Landed'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Landed'));
INSERT INTO track (name, preview, content) VALUES ('Losses', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Losses.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Losses'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Losses'));
INSERT INTO track (name, preview, content) VALUES ('Not You Too (feat. Chris Brown)', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Not_You_Too_(feat._Chris_Brown).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Not You Too (feat. Chris Brown)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Not You Too (feat. Chris Brown)'));
INSERT INTO track (name, preview, content) VALUES ('Pain 1993 (feat. Playboi Carti)', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Pain_1993_(feat._Playboi_Carti).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Pain 1993 (feat. Playboi Carti)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Pain 1993 (feat. Playboi Carti)'));
INSERT INTO track (name, preview, content) VALUES ('Time Flies', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Time_Flies.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Time Flies'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Time Flies'));
INSERT INTO track (name, preview, content) VALUES ('Toosie Slide', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/Toosie_Slide.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'Toosie Slide'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Toosie Slide'));
INSERT INTO track (name, preview, content) VALUES ('War', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/War.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'War'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'War'));
INSERT INTO track (name, preview, content) VALUES ('When To Say When', '/images/tracks/Drake/Dark_Lane_Demo_Tapes.jpg', '/audio/Drake/Dark_Lane_Demo_Tapes/When_To_Say_When.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Dark Lane Demo Tapes'), (SELECT id FROM track WHERE name = 'When To Say When'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'When To Say When'));
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Slime You Out (feat. SZA)', (SELECT id FROM artist WHERE name = 'Drake') , '/images/tracks/Drake/Slime_You_Out_(feat._SZA).jpg', '2023-09-15');
INSERT INTO track (name, preview, content) VALUES ('Slime You Out (feat. SZA)', '/images/tracks/Drake/Slime_You_Out_(feat._SZA).jpg', '/audio/Drake/Slime_You_Out_(feat._SZA)/Slime_You_Out_(feat._SZA).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Slime You Out (feat. SZA)'), (SELECT id FROM track WHERE name = 'Slime You Out (feat. SZA)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'Slime You Out (feat. SZA)'));
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('On The Radar Freestyle (feat. Central Cee)', (SELECT id FROM artist WHERE name = 'Drake') , '/images/tracks/Drake/On_The_Radar_Freestyle_(feat._Central_Cee).jpg', '2023-07-21');
INSERT INTO track (name, preview, content) VALUES ('On The Radar Freestyle (feat. Central Cee)', '/images/tracks/Drake/On_The_Radar_Freestyle_(feat._Central_Cee).jpg', '/audio/Drake/On_The_Radar_Freestyle_(feat._Central_Cee)/On_The_Radar_Freestyle_(feat._Central_Cee).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'On The Radar Freestyle (feat. Central Cee)'), (SELECT id FROM track WHERE name = 'On The Radar Freestyle (feat. Central Cee)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Drake'), (SELECT id FROM track WHERE name = 'On The Radar Freestyle (feat. Central Cee)'));
INSERT INTO artist (name, avatar)  VALUES ('Travis Scott', '/images/avatars/artists/Travis_Scott.jpg');
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Astroworld', (SELECT id FROM artist WHERE name = 'Travis Scott') , '/images/tracks/Travis_Scott/Astroworld.jpg', '2018-08-03');
INSERT INTO track (name, preview, content) VALUES ('Stargazing', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Stargazing.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Stargazing'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Stargazing'));
INSERT INTO track (name, preview, content) VALUES ('NC-17 (feat. 21 Savage)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/NC-17_(feat._21_Savage).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'NC-17 (feat. 21 Savage)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'NC-17 (feat. 21 Savage)'));
INSERT INTO track (name, preview, content) VALUES ('Houstonfornication', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Houstonfornication.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Houstonfornication'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Houstonfornication'));
INSERT INTO track (name, preview, content) VALUES ('Coffee Bean', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Coffee_Bean.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Coffee Bean'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Coffee Bean'));
INSERT INTO track (name, preview, content) VALUES ('Astrothunder', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Astrothunder.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Astrothunder'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Astrothunder'));
INSERT INTO track (name, preview, content) VALUES ('Stop Trying To Be God', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Stop_Trying_To_Be_God.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Stop Trying To Be God'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Stop Trying To Be God'));
INSERT INTO track (name, preview, content) VALUES ('Wake Up', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Wake_Up.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Wake Up'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Wake Up'));
INSERT INTO track (name, preview, content) VALUES ('Can''t Say', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Can''t_Say.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Can''t Say'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Can''t Say'));
INSERT INTO track (name, preview, content) VALUES ('Carousel', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Carousel.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Carousel'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Carousel'));
INSERT INTO track (name, preview, content) VALUES ('5% Tint', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/5%_Tint.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = '5% Tint'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = '5% Tint'));
INSERT INTO track (name, preview, content) VALUES ('Skeletons (feat. The Weekend)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Skeletons_(feat._The_Weekend).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Skeletons (feat. The Weekend)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Skeletons (feat. The Weekend)'));
INSERT INTO track (name, preview, content) VALUES ('No Bystanders (feat. Juice WRLD)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/No_Bystanders_(feat._Juice_WRLD).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'No Bystanders (feat. Juice WRLD)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'No Bystanders (feat. Juice WRLD)'));
INSERT INTO track (name, preview, content) VALUES ('Yosemite', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Yosemite.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Yosemite'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Yosemite'));
INSERT INTO track (name, preview, content) VALUES ('R.I.P. Screw (feat. Swae Lee)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/R.I.P._Screw_(feat._Swae_Lee).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'R.I.P. Screw (feat. Swae Lee)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'R.I.P. Screw (feat. Swae Lee)'));
INSERT INTO track (name, preview, content) VALUES ('Butterfly Effect', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Butterfly_Effect.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Butterfly Effect'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Butterfly Effect'));
INSERT INTO track (name, preview, content) VALUES ('Sicko Mode (feat. Drake)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Sicko_Mode_(feat._Drake).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Sicko Mode (feat. Drake)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Sicko Mode (feat. Drake)'));
INSERT INTO track (name, preview, content) VALUES ('Who What! (feat. Quavo)', '/images/tracks/Travis_Scott/Astroworld.jpg', '/audio/Travis_Scott/Astroworld/Who_What!_(feat._Quavo).mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Astroworld'), (SELECT id FROM track WHERE name = 'Who What! (feat. Quavo)'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Who What! (feat. Quavo)'));
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Escape Plan', (SELECT id FROM artist WHERE name = 'Travis Scott') , '/images/tracks/Travis_Scott/Escape_Plan.jpg', '2021-11-05');
INSERT INTO track (name, preview, content) VALUES ('Escape Plan', '/images/tracks/Travis_Scott/Escape_Plan.jpg', '/audio/Travis_Scott/Escape_Plan/Escape_Plan.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Escape Plan'), (SELECT id FROM track WHERE name = 'Escape Plan'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Escape Plan'));
INSERT INTO album (name, artist_id, preview, release_date) VALUES ('Highest In The Room', (SELECT id FROM artist WHERE name = 'Travis Scott') , '/images/tracks/Travis_Scott/Highest_In_The_Room.jpg', '2019-10-04');
INSERT INTO track (name, preview, content) VALUES ('Highest In The Room', '/images/tracks/Travis_Scott/Highest_In_The_Room.jpg', '/audio/Travis_Scott/Highest_In_The_Room/Highest_In_The_Room.mp3');INSERT INTO album_track (album_id, track_id) VALUES ((SELECT id FROM album WHERE name = 'Highest In The Room'), (SELECT id FROM track WHERE name = 'Highest In The Room'));INSERT INTO artist_track (artist_id, track_id) VALUES ((SELECT id FROM artist WHERE name = 'Travis Scott'), (SELECT id FROM track WHERE name = 'Highest In The Room'));
