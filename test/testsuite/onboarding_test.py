import unittest
import requests
import utils


class OnboardingTest(unittest.TestCase):
    def test_onboarding_success(self):
        response = requests.get(utils.url + '/artists')
        self.assertEqual(response.status_code, 200)
        self.assertNotEmpty(response.json())

        artists = response.json()['Artists']
        self.assertNotEmpty(artists[0]['Name'])
        self.assertNotEmpty(artists[0]['Avatar'])

        response = requests.get(utils.url + '/genres')
        self.assertEqual(response.status_code, 200)
        self.assertNotEmpty(response.json())

        genres = response.json()['Genres']
        self.assertNotEmpty(genres[0]['Preview'])
        self.assertNotEmpty(genres[0]['Name'])

        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        response = requests.post(utils.url + '/artists', cookies=cookies, headers=headers, json={'Artists': [{'Id': 1}]})
        self.assertEqual(response.status_code, 204)

        response = requests.post(utils.url + '/genres', cookies=cookies, headers=headers, json={'Genres': [{'Id': 2}]})
        self.assertEqual(response.status_code, 204)

    def assertNotEmpty(self, obj):
        self.assertTrue(obj)
