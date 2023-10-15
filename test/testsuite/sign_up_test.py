import unittest
import requests

url = 'http://localhost:8080/api/v1'


class SignUpTest(unittest.TestCase):
    def test_signup_endpoint_success(self):
        # get csrf token
        pre_response = requests.get(url + '/album/1')

        # use it in post header and cookies (double submit)
        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        # our register data
        data = {
            'Email': 'new.user@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)
        # get ok code and sessionId
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

    def test_signup_endpoint_without_csrf(self):
        # we have no csrf and do request
        # it's not important what data we have
        data = {}
        response = requests.post(url + "/sign_up", data=data)
        self.assertEqual(response.status_code, 403)

    def test_signup_endpoint_with_invalid_data(self):
        # we have csrf, but we have invalid data
        # use another get method for getting csrf
        pre_response = requests.get(url + '/auth')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }
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

        response = requests.post(url + "/sign_up", headers=headers, json=invalid_email, cookies=cookies)
        self.assertEqual(response.status_code, 400)
        response = requests.post(url + "/sign_up", headers=headers, json=invalid_username, cookies=cookies)
        self.assertEqual(response.status_code, 400)
        response = requests.post(url + "/sign_up", headers=headers, json=invalid_password, cookies=cookies)
        self.assertEqual(response.status_code, 400)

    def test_signup_endpoint_create_same_user(self):
        # get csrf token
        pre_response = requests.get(url + '/album/1')

        # use it in post header and cookies (double submit)
        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        # our register data
        data = {
            'Email': 'new.unique_user@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)
        # get ok code and sessionId
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")
        # we have this user! conflict!

        response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)
        self.assertEqual(response.status_code, 409)


sign_up_test = SignUpTest()
sign_up_test.test_signup_endpoint_success()
sign_up_test.test_signup_endpoint_without_csrf()
sign_up_test.test_signup_endpoint_with_invalid_data()
sign_up_test.test_signup_endpoint_create_same_user()
