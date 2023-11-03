import unittest
import requests
import utils


class ArtistTest(unittest.TestCase):
    def test_artist_get_success(self):
        artist_id = 1
        response = requests.get(utils.url + '/artist/' + str(artist_id))

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['Id'], artist_id)
        self.assertNotEqual(response.json()['Name'], '')
        self.assertNotEqual(response.json()['Avatar'], '')
        self.assertNotEqual(response.json()['Albums'], None)
        self.assertNotEqual(response.json()['Tracks'], None)

    def test_artist_like_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        artist_id = 1
        response = requests.get(utils.url + '/artist/' + str(artist_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

        response = requests.post(utils.url + '/artist/' + str(artist_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/artist/' + str(artist_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(True, response.json()['IsLiked'])

        response = requests.delete(utils.url + '/artist/' + str(artist_id) + '/unlike', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/artist/' + str(artist_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])
