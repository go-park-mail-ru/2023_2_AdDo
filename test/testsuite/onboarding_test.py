import unittest
import requests
import utils


class OnboardingTest(unittest.TestCase):
    def test_onboarding_success(self):
        response = requests.get(utils.url + '/artists')
        self.assertEqual(response.status_code, 200)

        self.assertNotEqual(response.json(), [])
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Avatar'], '')

        response = requests.get(utils.url + '/genres')
        self.assertEqual(response.status_code, 200)

        self.assertNotEqual(response.json(), [])
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['Name'], '')

        headers, cookies = utils.init_authorized_user_headers_and_cookies()

        response = requests.post(utils.url + '/artists', cookies=cookies, headers=headers, json={'Artists': [{'Id': 1}]})
        self.assertEqual(response.status_code, 204)

        response = requests.post(utils.url + '/genres', cookies=cookies, headers=headers, json={'Genres': [{'Id': 2}]})
        self.assertEqual(response.status_code, 204)
