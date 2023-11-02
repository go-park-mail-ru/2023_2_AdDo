def test_artist_info_success(self):
    artist_id = 1
    response = requests.get(utils.url + '/artist/' + str(artist_id))

    self.assertEqual(response.status_code, 200)
    self.assertEqual(response.json()['Id'], artist_id)
    self.assertNotEqual(response.json()['Name'], '')
    self.assertNotEqual(response.json()['Avatar'], '')
    self.assertNotEqual(response.json()['Albums'], None)
    self.assertNotEqual(response.json()['Tracks'], None)