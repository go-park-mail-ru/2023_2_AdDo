import unittest
import requests
import utils


class UserAvatarTest(unittest.TestCase):
    def test_user_avatar_forbidden(self):
        response = requests.get(utils.url + '/upload_avatar')
        self.assertEqual(response.status_code, 403)

        response = requests.get(utils.url + '/remove_avatar')
        self.assertEqual(response.status_code, 403)

    def test_user_avatar_unauthorized(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.get(utils.url + '/upload_avatar', cookies=cookies, headers=headers)
        self.assertEqual(response.status_code, 401)

        response = requests.get(utils.url + '/remove_avatar', cookies=cookies, headers=headers)
        self.assertEqual(response.status_code, 401)

    def test_user_avatar_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()
        avatar = open('user_avatar_image.png', 'rb')
        files = {'Avatar': avatar}

        response = requests.get(utils.url + '/upload_avatar', headers=headers, cookies=cookies, files=files)
        self.assertEqual(response.status_code, 204)
        self.assertNotEqual(response.json()['Url'], '')

        response = requests.get(utils.url + '/remove_avatar', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 204)
        self.assertEqual(response.json()['Url'], '')
