    erDiagram
     
    profile ||..o{ playlist : create

    profile ||--o{ profile_album : liked
    album ||--o{ profile_album : contains

    profile ||--o{ profile_artist : liked
    artist ||--o{ profile_artist : contains

    profile ||--o{ profile_playlist : liked
    playlist ||--o{ profile_playlist : contains

    profile ||--o{ profile_track : liked
    track ||--o{ profile_track : contains
    
    track ||--o{ album_track : includes
    album ||--o{ album_track : contains

    track ||--o{ artist_track : includes
    artist ||--o{ artist_track : create

    track ||--o{ playlist_track : includes
    playlist ||--o{ playlist_track : contains

    profile {
        uuid id PK
        string email
        string nickname
        string password
        date birth_date
        string avatar_url
    }

    artist {
        serial id PK
        string name 
        string avatar
    }

    playlist {
        serial id PK
        string name 
        uuid creator_id FK
        string preview
        timestamp creating_date
    }

    album {
        serial id
        string name 
        int artist_id FK
        string preview
        date release_date
    }

    track {
        serial id
        string name 
        string preview
        string content
        int play_count
    }

    album_track {
        int track_id FK
        int album_id Fk
    }

    artist_track {
        int track_id FK
        int artist_id FK
    }

    playlist_track {
        int track_id FK
        int playlist_id
    }

    profile_track {
        uuid profile_id FK
        int track_id FK
    } 
     
    profile_artist {
        uuid profile_id FK
        int artist_id FK
    } 
    profile_album {
        uuid profile_id FK
        int album_id FK
    } 

    profile_playlist {
        uuid profile_id FK
        int playlist_id FK
    }
