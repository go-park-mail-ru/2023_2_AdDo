docker build -t registry.musicon.space/user -f build/package/micros/user/Dockerfile .
docker build -t registry.musicon.space/session -f build/package/micros/session/Dockerfile .
docker build -t registry.musicon.space/images -f build/package/micros/images/Dockerfile .
docker build -t registry.musicon.space/track -f build/package/micros/track/Dockerfile .
docker build -t registry.musicon.space/album -f build/package/micros/album/Dockerfile .
docker build -t registry.musicon.space/playlist -f build/package/micros/playlist/Dockerfile .
docker build -t registry.musicon.space/artist -f build/package/micros/artist/Dockerfile .
docker build -t registry.musicon.space/entrypoint -f build/package/micros/entrypoint/Dockerfile .
