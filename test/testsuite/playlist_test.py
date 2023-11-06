import unittest
import requests
import utils


class PlaylistTest(unittest.TestCase):
    def test_playlist_use_case_success(self):
        headers_creator, cookies_creator = utils.init_authorized_user_headers_and_cookies()
        headers_random_user, cookies_random_user = utils.init_authorized_user_headers_and_cookies()

        response = requests.post(utils.url + '/playlist', headers=headers_creator, cookies=cookies_creator)
        self.assertEqual(200, response.status_code)

        playlist_id = response.json()['Id']
        first_track_id = 1
        second_track_id = 2
        response = requests.post(utils.url + '/playlist/' + str(playlist_id) + '/add_track', headers=headers_random_user, cookies=cookies_random_user, json={'Id': first_track_id})
        self.assertEqual(403, response.status_code)

        response = requests.post(utils.url + '/playlist/' + str(playlist_id) + '/add_track', headers=headers_creator, cookies=cookies_creator, json={'Id': first_track_id})
        self.assertEqual(204, response.status_code)

        response = requests.post(utils.url + '/playlist/' + str(playlist_id) + '/add_track', headers=headers_creator, cookies=cookies_creator, json={'Id': second_track_id})
        self.assertEqual(204, response.status_code)

        response = requests.delete(utils.url + '/playlist/' + str(playlist_id) + '/remove_track', headers=headers_creator, cookies=cookies_creator, json={'Id': second_track_id})
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/playlist/' + str(playlist_id), headers=headers_random_user, cookies=cookies_random_user)
        self.assertEqual(200, response.status_code)
        self.assertNotEqual('', response.json()['AuthorName'])
        self.assertNotEqual('', response.json()['AuthorId'])
        self.assertNotEqual('', response.json()['Name'])
        self.assertNotEqual([], response.json()['Tracks'])

        response = requests.put(utils.url + '/playlist/' + str(playlist_id) + '/make_private', headers=headers_creator, cookies=cookies_creator)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/playlist/' + str(playlist_id), headers=headers_random_user, cookies=cookies_random_user)
        self.assertEqual(403, response.status_code)

        response = requests.put(utils.url + '/playlist/' + str(playlist_id) + '/make_public', headers=headers_creator, cookies=cookies_creator)
        self.assertEqual(204, response.status_code)

        response = requests.get(utils.url + '/playlist/' + str(playlist_id), headers=headers_random_user, cookies=cookies_random_user)
        self.assertEqual(200, response.status_code)

        preview = open('user_avatar_image.png', 'rb')
        files = {'Preview': preview}
        response = requests.post(utils.url + '/playlist/' + str(playlist_id) + '/update_preview', headers=headers_creator, cookies=cookies_creator, files=files)
        self.assertEqual(204, response.status_code)

        response = requests.delete(utils.url + '/playlist/' + str(playlist_id), headers=headers_creator, cookies=cookies_creator)
        self.assertEqual(204, response.status_code)
