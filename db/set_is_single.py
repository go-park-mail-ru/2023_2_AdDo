import requests

# url = 'http://localhost:8888/api/v1'
url = "https://musicon.space/api/v1"
albumsNum = 3200

# UPDATE album SET is_single = true WHERE id IN (1, 2, 3);

with open('set_is_single.sql', "w") as f:
    f.write("UPDATE album SET is_single = true WHERE id IN (")

    first_id_is_written = False

    for i in range(1, albumsNum):
        print('url:', url + '/album/' + str(i))
        r = requests.get(url + '/album/' + str(i))

        if r.status_code != 200:
            print('Not found')
            continue

        tracks = r.json()['Tracks']
        print('Album name:', r.json()['Name'], ',  Tracks num:', len(tracks))

        if len(tracks) == 1:
            print('Single')

            if first_id_is_written:
                f.write(', ')
            f.write(str(i))

            if not first_id_is_written:
                first_id_is_written = True

        print('\n')

    f.write(');')
