import unittest
import requests

url = 'http://localhost:8888/api/v1'


class ListenTest(unittest.TestCase):
    def test_listen_success(self):
        pre_response = requests.get(url + '/auth')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        response = requests.post(url + '/listen', headers=headers, cookies=cookies, json={'Id': 1})
        self.assertEqual(200, response.status_code)

    def test_listen_without_csrf(self):
        response = requests.post(url + '/listen', json={'Id': 1})
        self.assertEqual(403, response.status_code)

test_listen = ListenTest()
test_listen.test_listen_without_csrf()
test_listen.test_listen_success()
