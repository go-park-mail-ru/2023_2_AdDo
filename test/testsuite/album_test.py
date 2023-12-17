import unittest
import requests
import utils


class AlbumTest(unittest.TestCase):
    def test_getting_album(self):
        album_id = '💕💕💕💕💕💕💕💕💕'
        response = requests.get(utils.url + '/album/' + str(album_id))
        self.assertEqual(response.status_code, 400)

        album_id = '🌸🌸🌸🌿️🌿️🌿️✨✨✨'
        response = requests.get(utils.url + '/album/' + str(album_id))
        self.assertEqual(response.status_code, 400)

        album_id = '0'
        response = requests.get(utils.url + '/album/' + str(album_id))
        self.assertEqual(response.status_code, 500)

        album_id = 1
        response = requests.get(utils.url + '/album/' + str(album_id))
        self.check_success_album_response(response)

    def test_getting_album_with_required_track(self):
        track_id = '(｡♥‿♥｡)  (◍•ᴗ•◍)❤'
        response = requests.get(utils.url + '/track/' + str(track_id))
        self.assertEqual(response.status_code, 400)

        track_id = '-100'
        response = requests.get(utils.url + '/track/' + str(track_id))
        self.assertEqual(response.status_code, 500)

        track_id = 10
        response = requests.get(utils.url + '/track/' + str(track_id))
        self.check_success_album_response(response)

    def check_success_album_response(self, response):
        self.assertEqual(response.status_code, 200)
        self.assertNotEmpty(response.json())

        self.assertGreater(response.json()['Id'], 0)
        self.assertNotEmpty(response.json()['Name'])
        self.assertNotEmpty(response.json()['Preview'])
        self.assertGreater(response.json()['ArtistId'], 0)
        self.assertNotEmpty(response.json()['ArtistName'])
        self.assertNotEmpty(response.json()['Tracks'])

    def test_album_like(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        album_id = -10
        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(404, response.status_code)

        album_id = 1
        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

        response = requests.post(utils.url + '/album/' + str(album_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.post(utils.url + '/album/' + str(album_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(500, response.status_code)

        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(True, response.json()['IsLiked'])

        response = requests.delete(utils.url + '/album/' + str(album_id) + '/unlike', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

    def assertEmpty(self, obj):
        self.assertFalse(obj)

    def assertNotEmpty(self, obj):
        self.assertTrue(obj)
