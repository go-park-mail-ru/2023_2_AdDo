import json

with open('en_ru_genres_dict.json') as f:
    en_ru_dict = json.load(f)

with open('genres.sql', 'w') as f:
    for en, ru in en_ru_dict.items():
        url = '/images/genre_icons/' + en + '.png'
        f.write("UPDATE genre SET ru_name = '{}', icon_url = '{}' WHERE name = '{}';\n".format(ru, url, en))
