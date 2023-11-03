import unittest
import requests
import utils


class UpdateInfo(unittest.TestCase):
    def test_update_info_forbidden(self):
        response = requests.put(utils.url + '/update_info')
        self.assertEqual(response.status_code, 403)

    def test_update_info_unauthorized(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.put(utils.url + '/update_info', cookies=cookies, headers=headers)
        self.assertEqual(response.status_code, 401)

    def test_update_info_success(self):
        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        new_data = utils.gen_random_valid_update_data()
        response = requests.put(utils.url + '/update_info', headers=headers, cookies=cookies, json=new_data)
        self.assertEqual(response.status_code, 204)

        response = requests.get(utils.url + '/me', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.json()['Id'], 0)
        self.assertEqual(response.json()['Username'], new_data['Username'])
        self.assertEqual(response.json()['Email'], new_data['Email'])
        self.assertEqual(response.json()['BirthDate'], new_data['BirthDate'])
