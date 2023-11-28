import unittest
import requests
import utils


class AlbumTest(unittest.TestCase):
    def test_album_get_success(self):
        album_id = 1
        response = requests.get(utils.url + '/album/' + str(album_id))

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['Id'], album_id)
        self.assertNotEqual(response.json()['Name'], '')
        self.assertNotEqual(response.json()['Preview'], '')
        self.assertNotEqual(response.json()['ArtistId'], 0)
        self.assertNotEqual(response.json()['ArtistName'], '')
        self.assertNotEqual(response.json()['Tracks'], None)

    def test_getting_albums_with_required_track(self):
        track_id = 10
        response = requests.get(utils.url + '/track/' + str(track_id))

        self.assertEqual(response.status_code, 200)
        self.assertTrue(response.json())
        self.assertTrue(response.json()[0]['Id'])
        self.assertTrue(response.json()[0]['Name'])
        self.assertTrue(response.json()[0]['Preview'])
        self.assertTrue(response.json()[0]['ArtistId'])
        self.assertTrue(response.json()[0]['ArtistName'])
        self.assertFalse(response.json()[0]['Tracks'])

    def test_album_like_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        album_id = 1
        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

        response = requests.post(utils.url + '/album/' + str(album_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(True, response.json()['IsLiked'])

        response = requests.delete(utils.url + '/album/' + str(album_id) + '/unlike', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/album/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])
