import requests

# url = 'http://localhost:8888/api/v1'
url = "https://musicon.space/api/v1"

albumsNum = 700

with open('set_is_single.sql', "w") as f:
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
            f.write('update album set is_single = true where id = {};\n'.format(i))

        print('\n')
