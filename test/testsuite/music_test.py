import unittest
import requests
import utils


class MusicTest(unittest.TestCase):
    def test_new_success(self):
        response = requests.get(utils.url + '/new')
        self.check_albums_response(response)

    def test_feed_success(self):
        response = requests.get(utils.url + '/feed')
        self.check_albums_response(response)

    def test_popular_success(self):
        response = requests.get(utils.url + '/popular')
        self.check_albums_response(response)

    def test_most_liked_success(self):
        response = requests.get(utils.url + '/most_liked')
        self.check_albums_response(response)

    def check_albums_response(self, response):
        self.assertEqual(response.status_code, 200)
        self.assertNotEmpty(response.json())
        self.assertNotEmpty(response.json()['Albums'])

        first_album = response.json()['Albums'][0]

        self.assertGreater(first_album['Id'], 0)
        self.assertNotEmpty(first_album['Name'])
        self.assertNotEmpty(first_album['Preview'])
        self.assertNotEmpty(first_album['ArtistId'])
        self.assertNotEmpty(first_album['ArtistName'])
        self.assertEmpty(first_album['Tracks'])

    def assertEmpty(self, obj):
        self.assertFalse(obj)

    def assertNotEmpty(self, obj):
        self.assertTrue(obj)
