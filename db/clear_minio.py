import psycopg2
from psycopg2.extras import NamedTupleCursor
from minio import Minio

minio_client = Minio(
    "api.s3.musicon.space",
    access_key="BeKn8vgv9gmYGrfPkvoc",
    secret_key="EzWtcUjJUeci8tKVg3SJjetPYK07Zw1N60DQ5IuD",
)
minio_map = {}
track_map = {}

conn = psycopg2.connect('postgresql://musicon:Music0nSecure@82.146.45.164:5433/musicon')
cursor = conn.cursor(cursor_factory=NamedTupleCursor)
cursor.execute(f'select * from track')
tracks = cursor.fetchall()

for track in tracks:
    track_map[str(track[3])[7:]] = 1

objects = minio_client.list_objects("audio", recursive=True)

for obj in objects:
    if obj.object_name not in track_map:
        print(obj.object_name)
        # minio_client.remove_object("audio", obj.object_name)

