import unittest
import requests

url = 'http://localhost:8080/api/v1'


class MeTest(unittest.TestCase):
    def test_me_success(self):
        pre_response = requests.get(url + '/album/1')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        register_data = {
            'Email': 'alex@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=register_data, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        cookies['JSESSIONID'] = response.cookies['JSESSIONID']

        response = requests.get(url + '/me', headers=headers, cookies=cookies)
        self.assertEqual(response.status_code, 200)

        self.assertNotEqual(response.json()['Id'], 0)
        self.assertEqual(response.json()['Username'], 'username')
        self.assertEqual(response.json()['Email'], 'alex@mail.ru')
        self.assertEqual(response.json()['Password'], 'userPassword')
        self.assertEqual(response.json()['BirthDate'], '12-01-2003')
        self.assertEqual(response.json()['Avatar'], '')
