import unittest
import requests
import names

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

        username = names.get_full_name().replace(' ', '').lower()
        email = username + '@mail.ru'
        register_data = {
            'Email': email,
            'Password': 'userPassword',
            'Username': username,
            'BirthDate': '2003-01-12',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=register_data, cookies=cookies)
        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")

        cookies['JSESSIONID'] = response.cookies['JSESSIONID']

        response = requests.get(url + '/me', headers=headers, cookies=cookies)

        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.json()['Id'], 0)
        self.assertEqual(response.json()['Username'], register_data['Username'])
        self.assertEqual(response.json()['Email'], register_data['Email'])
        self.assertEqual(response.json()['BirthDate'], register_data['BirthDate'])
        self.assertEqual(response.json()['Avatar'], '')


me_test = MeTest()
me_test.test_me_success()
