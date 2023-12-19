#!/bin/bash

echo "Generating mocks..."
mockgen -source=internal/pkg/user/domain.go -destination=test/mocks/user/user_mock.go -package=user_mock \
&& mockgen -source=internal/pkg/session/domain.go -destination=test/mocks/session/session_mock.go -package=session_mock \
&& mockgen -source=internal/pkg/track/domain.go -destination=test/mocks/track/track_mock.go -package=track_mock \
&& mockgen -source=internal/pkg/artist/domain.go -destination=test/mocks/artist/artist_mock.go -package=artist_mock \
&& mockgen -source=internal/pkg/album/domain.go -destination=test/mocks/album/album_mock.go -package=album_mock \
&& mockgen -source=internal/pkg/image/domain.go -destination=test/mocks/avatar/avatar_mock.go -package=avatar_mock \
&& mockgen -source=internal/pkg/playlist/domain.go -destination=test/mocks/playlist/playlist_mock.go -package=playlist_mock \
&& mockgen -source=internal/pkg/activity/domain.go -destination=test/mocks/activity/activity_mock.go -package=activity_mock

if [ $? -ne "0" ];then
    echo "Generating mocks is failed"
    exit 1
fi
echo "Mocks successful generated"

# for domain in internal/pkg/*/domain.go; do
#     package_name=$(basename $(dirname $domain))
#     echo "mockgen -source=${domain} -destination=test/mocks/${package_name}/${package_name}_mock.go -package=${package_name}_mock;"
#     mockgen -source=${domain} -destination=test/mocks/${package_name}/${package_name}_mock.go -package=${package_name}_mock;
# done
# mockgen -source=internal/pkg/image/domain.go -destination=test/mocks/avatar/avatar_mock.go -package=avatar_mock;
