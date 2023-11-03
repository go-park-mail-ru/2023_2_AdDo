import unittest
import requests
import utils


class PlaylistTest(unittest.TestCase):
    def test_playlist_get_success(self):
        playlist_id = 1
        response = requests.get(utils.url + '/album/' + str(play))

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['Id'], album_id)
        self.assertNotEqual(response.json()['Name'], '')
        self.assertNotEqual(response.json()['Preview'], '')
        self.assertNotEqual(response.json()['ArtistId'], 0)
        self.assertNotEqual(response.json()['ArtistName'], '')
        self.assertNotEqual(response.json()['Tracks'], None)

    def test_playlist_use_case_success(self):
        headers_creator, cookies_creator = utils.init_authorized_user_headers_and_cookies()
        headers_random_user, cookies_random_user = utils.init_authorized_user_headers_and_cookies()


        response = requests.post(utils.url + '/playlist/' + str(album_id) + '/is_like', headers=headers, cookies=cookies)
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
