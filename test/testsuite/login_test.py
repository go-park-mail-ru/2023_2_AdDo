import unittest
import requests
import utils


class LoginTest(unittest.TestCase):
    def test_login_success(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()
        data = utils.gen_random_valid_register_data()

        response = requests.post(utils.url + "/sign_up", headers=headers, json=data, cookies=cookies)
        self.assertEqual(response.status_code, 204)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        user_cred = {
            'Email': data['Email'],
            'Password': data['Password']
        }

        response = requests.post(utils.url + '/login', headers=headers, json=user_cred, cookies=cookies)
        self.assertEqual(response.status_code, 204)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

    def test_login_forbidden(self):
        response = requests.post(utils.url + "/login")
        self.assertEqual(response.status_code, 403)

        headers, cookies = utils.get_csrf_headers_and_cookies()
        data = utils.gen_random_valid_register_data()
        user_cred = {
            'Email': data['Email'],
            'Password': data['Password']
        }

        response = requests.post(utils.url + '/login', headers=headers, json=user_cred, cookies=cookies)
        self.assertEqual(response.status_code, 403)

    def test_login_bad_request(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()
        response = requests.post(utils.url + '/login', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 400)
