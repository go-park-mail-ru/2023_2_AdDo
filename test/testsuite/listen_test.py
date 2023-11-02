import unittest
import requests
import utils


class ListenTest(unittest.TestCase):
    def test_listen_success(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.post(utils.url + '/listen', headers=headers, cookies=cookies, json={'Id': 1})
        self.assertEqual(204, response.status_code)

    def test_listen_forbidden(self):
        response = requests.post(utils.url + '/listen', json={'Id': 1})
        self.assertEqual(403, response.status_code)
