import unittest
import requests

url = 'http://localhost:8080/api/v1'


class LoginTest(unittest.TestCase):
    def test_login_endpoint_success(self):
        # get csrf token
        pre_response = requests.get(url + '/auth')

        # use it in post header and cookies (double submit)
        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        # our register data
        register_data = {
            'Email': 'new@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=register_data, cookies=cookies)
        print(response.text)
        # get ok code and sessionId
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        login_data = {
            'Email': 'new@mail.ru',
            'Password': 'userPassword',
        }
        response = requests.post(url + '/login', headers=headers, json=login_data, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

    def test_login_endpoint_without_csrf(self):
        # we have no csrf and do request
        # it's not important what data we have
        data = {}
        response = requests.post(url + "/login", data=data)
        self.assertEqual(response.status_code, 403)

    def test_login_without_account(self):
        # get csrf token
        pre_response = requests.get(url + '/auth')

        # use it in post header and cookies (double submit)
        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        login_data = {
            'Email': 'we.have.no@mail.ru',
            'Password': 'anyPassword',
        }
        response = requests.post(url + '/login', headers=headers, json=login_data, cookies=cookies)
        self.assertEqual(response.status_code, 403)


login_test = LoginTest()
login_test.test_login_endpoint_success()
login_test.test_login_endpoint_without_csrf()
login_test.test_login_without_account()
