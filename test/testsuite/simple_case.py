import unittest
import requests

url = 'http://localhost:8080/api/v1'


class SignUpTest(unittest.TestCase):
    def test_signup_endpoint(self):
        pre_response = requests.get(url + '/music')

        headers = {
            'X-Csrf-Token': pre_response.headers['X-Csrf-Token']
        }
        cookies = {
            'X-Csrf-Token': pre_response.cookies['X-Csrf-Token']
        }

        data = {
            'Email': 'new.user@mail.ru',
            'Password': 'userPassword',
            'Username': 'username',
            'BirthDate': '12-01-2003',
        }
        response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)

        self.assertEqual(response.status_code, 200)
        self.assertNotEqual(response.cookies['JSESSIONID'], "")


sign_up_test = SignUpTest()
sign_up_test.test_signup_endpoint()

# Makefile и идея интеграционного тестирования для полной базы данных и пустой
# gomock + test connection
# как очистить базу между тестами, норм ли прям из питона
# не надо вообще чистить базу
# как отдавать csrf токен? сделать отдельную ручку для этого? или в любом гете
# Нужно ли поднимать сервак на https (http.ListenAndServeTLS) или мы обеспечиваем это только nginx'ом
# поднимаем на хттп, остальное через нгинкс
# Есть ли некое подобие макросов, если нужно делать что-то в дебаг сборке, а что-то в релиз
# что такое хром драйвер, имелись в виду полные тесты вместе с фронтом
