import unittest
import requests

url = 'http://localhost:8080/api/v1'


class SignUpTest(unittest.TestCase):
    def test_logout_endpoint(self):
        request_url = url + '/sign_up'

        headers = {'Content-Type': 'application/json'}
        data = {'email': 'user', 'key2': 'value2'}

        response = requests.post(request_url, headers=headers, data=data)

        self.assertEqual(response.status_code, 200)
