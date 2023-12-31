import auth_test
import track_test
import album_test
import artist_test
import playlist_test
import login_test
import user_avatar_test
import logout_test
import me_test
import update_user_info
import music_test
import sign_up_test
import onboarding_test

auth_test = auth_test.AuthTest()
auth_test.test_auth_unauthorized()
auth_test.test_auth_success()


sign_up_test = sign_up_test.SignUpTest()
#sign_up_test.test_signup_forbidden()
sign_up_test.test_signup_bad_request_no_data()
sign_up_test.test_signup_bad_request_invalid_data()
sign_up_test.test_signup_no_content()
sign_up_test.test_signup_conflict()


login_test = login_test.LoginTest()
#login_test.test_login_forbidden()
login_test.test_login_bad_request()
login_test.test_login_success()


user_avatar_test = user_avatar_test.UserAvatarTest()
#user_avatar_test.test_user_avatar_forbidden()
user_avatar_test.test_user_avatar_unauthorized()
user_avatar_test.test_user_avatar_success()


update_user_info = update_user_info.UpdateInfo()
#update_user_info.test_update_info_forbidden()
update_user_info.test_update_info_unauthorized()
update_user_info.test_update_info_success()


me_test = me_test.MeTest()
#me_test.test_me_forbidden()
me_test.test_me_unauthorized()
me_test.test_me_success()


logout_test = logout_test.LogoutTest()
logout_test.test_logout_success()


track_test = track_test.TrackTest()
#track_test.test_track_listen_forbidden()
track_test.test_track_listen_success()
track_test.test_track_like_success()
track_test.test_track_collection_success()


album_test = album_test.AlbumTest()
album_test.test_getting_album()
album_test.test_getting_album_with_required_track()
album_test.test_album_like()


artist_test = artist_test.ArtistTest()
artist_test.test_artist_get_success()
artist_test.test_artist_like_success()


music_test = music_test.MusicTest()
music_test.test_feed_success()
music_test.test_new_success()
music_test.test_most_liked_success()
music_test.test_popular_success()


playlist_test = playlist_test.PlaylistTest()
playlist_test.test_playlist_use_case_success()


onboarding_test = onboarding_test.OnboardingTest()
onboarding_test.test_onboarding_success()