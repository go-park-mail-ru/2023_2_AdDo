import unittest
import requests
import utils


class TrackTest(unittest.TestCase):
    def test_track_listen_success(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.post(utils.url + '/listen', headers=headers, cookies=cookies, json={'Id': 1})
        self.assertEqual(204, response.status_code)

    def test_track_listen_forbidden(self):
        response = requests.post(utils.url + '/listen', json={'Id': 1})
        self.assertEqual(403, response.status_code)

    def test_track_like_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        track_id = 1
        response = requests.get(utils.url + '/track/' + str(track_id) + '/is_like', headers=headers, cookies=cookies)
        print(response.json)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

        response = requests.post(utils.url + '/track/' + str(track_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/track/' + str(track_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(True, response.json()['IsLiked'])

        response = requests.delete(utils.url + '/track/' + str(track_id) + '/unlike', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/track/' + str(track_id) + '/is_like', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(False, response.json()['IsLiked'])

    def test_track_collection_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        track_id = 1
        response = requests.post(utils.url + '/track/' + str(track_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        track_id = 2
        response = requests.post(utils.url + '/track/' + str(track_id) + '/like', headers=headers, cookies=cookies)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/collection/tracks', headers=headers, cookies=cookies)
        self.assertEqual(200, response.status_code)
        self.assertEqual(2, len(response.json()['Tracks']))
