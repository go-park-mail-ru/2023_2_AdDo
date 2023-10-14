import unittest
import requests

url = 'http://localhost:8080/api/v1'


class ListenTest(unittest.TestCase):
    def test_listen_success(self):
        pre_response = requests.get(url + '/album/1')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        data = {'Id': 1}
        response = requests.post(url + '/listen', headers=headers, json=data, cookies=cookies)
        self.assertEqual(200, response.status_code)

    def test_listen_without_csrf(self):
        response = requests.post(url + '/listen', data={'Id': 1})
        self.assertEqual(403, response.status_code)
