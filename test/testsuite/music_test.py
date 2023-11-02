import unittest
import requests
import utils


class MusicTest(unittest.TestCase):
    def test_new_success(self):
        response = requests.get(utils.url + '/new')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertEqual(response.json()[0]['Tracks'], [])

    def test_feed_success(self):
        response = requests.get(utils.url + '/feed')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertEqual(response.json()[0]['Tracks'], [])

    def test_popular_success(self):
        response = requests.get(utils.url + '/popular')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertEqual(response.json()[0]['Tracks'], [])

    def test_most_liked_success(self):
        response = requests.get(utils.url + '/most_liked')

        self.assertEqual(response.status_code, 200)
        self.assertGreater(len(response.json()), 0)

        self.assertNotEqual(response.json()[0]['Id'], 0)
        self.assertNotEqual(response.json()[0]['Name'], '')
        self.assertNotEqual(response.json()[0]['Preview'], '')
        self.assertNotEqual(response.json()[0]['ArtistId'], 0)
        self.assertNotEqual(response.json()[0]['ArtistName'], '')
        self.assertEqual(response.json()[0]['Tracks'], [])
