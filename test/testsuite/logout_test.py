import unittest
import requests

url = 'http://localhost:8080/api/v1'


class LogoutTest(unittest.TestCase):
    def test_logout_success(self):
        pre_response = requests.get(url + '/album/1')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        register_data = {
            'Email': 'ivanivanov@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=register_data, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        cookies['JSESSIONID'] = response.cookies['JSESSIONID']

        response = requests.post(url + '/logout', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.cookies, {})

logout_test = LogoutTest()
logout_test.test_logout_success()
