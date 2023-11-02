

def test_album_tracks_success(self):
    album_id = 1
    response = requests.get(utils.url + '/album/' + str(album_id))

    self.assertEqual(response.status_code, 200)
    self.assertEqual(response.json()['Id'], album_id)
    self.assertNotEqual(response.json()['Name'], '')
    self.assertNotEqual(response.json()['Preview'], '')
    self.assertNotEqual(response.json()['ArtistId'], 0)
    self.assertNotEqual(response.json()['ArtistName'], '')
    self.assertNotEqual(response.json()['Tracks'], None)