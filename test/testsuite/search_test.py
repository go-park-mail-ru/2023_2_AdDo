import unittest
import requests
import utils


class SearchTest(unittest.TestCase):
    def test_search_success(self):
        response = requests.get(utils.url + '/search?query=eminem')
        self.assertEqual(response.status_code, 200)
