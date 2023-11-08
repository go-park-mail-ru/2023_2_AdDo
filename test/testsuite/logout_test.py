import unittest
import requests
import utils


class LogoutTest(unittest.TestCase):
    def test_logout_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        response = requests.post(utils.url + '/logout', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 204)
        self.assertEqual(response.cookies, {})
