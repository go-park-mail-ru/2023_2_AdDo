import unittest
import requests

url = 'http://localhost:8080/api/v1'


class SignUpTest(unittest.TestCase):
    def test_signup_endpoint(self):
        pre_response = requests.get('http://localhost:8080/api/v1/music')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        data = {
            'Email': 'user@mail.ru',
            'Password':
            'userPassword',
            'Username': 'username',
            'BirthDate': '1-12-2003',
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }
        response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)

        print(response.text)

        self.assertEqual(response.status_code, 200)
        self.assertNotEquals(response.cookies['JSESSIONID'], "")
        print(response.cookies['JSESSIONID'])