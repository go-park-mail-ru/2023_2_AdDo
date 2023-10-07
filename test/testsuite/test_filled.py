import unittest
import requests

url = 'http://localhost:8080/api/v1'

class SignUpTest(unittest.TestCase):
    def test_logout_endpoint(self):
        request_url = url + '/sign_up'
        response = requests.post(request_url)

        # Проверка кода состояния
        self.assertEqual(response.status_code, 200)

        # Проверка ожидаемого ответа
        self.assertEqual(response.text, 'Logout successful')
