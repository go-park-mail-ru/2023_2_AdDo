import unittest
import requests
import utils


class MeTest(unittest.TestCase):
    def test_me_forbidden(self):
        response = requests.get(utils.url + '/me')
        self.assertEqual(response.status_code, 403)

    def test_me_unauthorized(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.get(utils.url + '/me', cookies=cookies, headers=headers)
        self.assertEqual(response.status_code, 401)

    def test_me_success(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()
        data = utils.gen_random_valid_register_data()
        response = requests.post(utils.url + "/sign_up", headers=headers, json=data, cookies=cookies)
        self.assertEqual(response.status_code, 204)

        cookies['JSESSIONID'] = response.cookies['JSESSIONID']

        response = requests.get(utils.url + '/me', headers=headers, cookies=cookies)

        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.json()['Id'], 0)
        self.assertEqual(response.json()['Username'], data['Username'])
        self.assertEqual(response.json()['Email'], data['Email'])
        self.assertEqual(response.json()['BirthDate'], data['BirthDate'])
        self.assertEqual(response.json()['Avatar'], '')
