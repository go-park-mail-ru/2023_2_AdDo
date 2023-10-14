import json
import unittest
import requests

url = 'http://localhost:8080/api/v1'


class MusicTest(unittest.TestCase):
    def test_artist_info_success(self):
        artist_id = 1
        response = requests.get(url + '/artist/' + str(artist_id))

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['Id'], artist_id)
        self.assertNotEqual(response.json()['Name'], '')
        self.assertNotEqual(response.json()['Avatar'], '')
        self.assertNotEqual(response.json()['Albums'], None)
        self.assertNotEqual(response.json()['Tracks'], None)

    def test_album_tracks_success(self):
        album_id = 1
        response = requests.get(url + '/album/' + str(album_id))

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['Id'], album_id)
        self.assertNotEqual(response.json()['Name'], '')
        self.assertNotEqual(response.json()['Preview'], '')
        self.assertNotEqual(response.json()['ArtistId'], 0)
        self.assertNotEqual(response.json()['ArtistName'], '')
        self.assertNotEqual(response.json()['Tracks'], None)

    def test_new_success(self):
        response = requests.get(url + '/new')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertNotEqual(response.json()[0]['Tracks'], None)

    def test_feed_success(self):
        response = requests.get(url + '/feed')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertNotEqual(response.json()[0]['Tracks'], None)
