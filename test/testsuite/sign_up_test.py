import unittest
import requests
import utils


class SignUpTest(unittest.TestCase):
    def test_signup_no_content(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()
        data = utils.gen_random_valid_register_data()

        response = requests.post(utils.url + "/sign_up", headers=headers, json=data, cookies=cookies)
        self.assertEqual(response.status_code, 204)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

    def test_signup_forbidden(self):
        response = requests.post(utils.url + "/sign_up")
        self.assertEqual(response.status_code, 403)

    def test_signup_bad_request_no_data(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        response = requests.post(utils.url + "/sign_up", headers=headers, cookies=cookies,)
        self.assertEqual(response.status_code, 400)

    def test_signup_bad_request_invalid_data(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()

        invalid_email = {
            'Email': 'Регулярка😘😘😘😘😘😘😘@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003'
        }
        invalid_username = {
            'Email': 'new.user@mail.ru',
            'Password': 'userPassword',
            'Username': ' 😘😘😘😘😘😘😘я сломаю  твой бэк',
            'BirthDate': '12-01-2003'
        }
        invalid_password = {
            'Email': 'new.user@mail.ru',
            'Password': 'veryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryveryverylongpassword',
            'Username': 'username',
            'BirthDate': '12-01-2003'
        }

        response = requests.post(utils.url + "/sign_up", headers=headers, json=invalid_email, cookies=cookies)
        self.assertEqual(response.status_code, 400)
        response = requests.post(utils.url + "/sign_up", headers=headers, json=invalid_username, cookies=cookies)
        self.assertEqual(response.status_code, 400)
        response = requests.post(utils.url + "/sign_up", headers=headers, json=invalid_password, cookies=cookies)
        self.assertEqual(response.status_code, 400)

    def test_signup_conflict(self):
        headers, cookies = utils.get_csrf_headers_and_cookies()
        data = utils.gen_random_valid_register_data()

        response = requests.post(utils.url + "/sign_up", headers=headers, json=data, cookies=cookies)

        self.assertEqual(response.status_code, 204)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        response = requests.post(utils.url + "/sign_up", headers=headers, json=data, cookies=cookies)
        self.assertEqual(response.status_code, 409)
