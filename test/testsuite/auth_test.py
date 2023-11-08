import unittest
import requests
import utils


class AuthTest(unittest.TestCase):
    def test_auth_unauthorized(self):
        response = requests.get(utils.url + '/auth')
        self.assertEqual(response.status_code, 401)

    def test_auth_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        response = requests.get(utils.url + '/auth', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 204)
