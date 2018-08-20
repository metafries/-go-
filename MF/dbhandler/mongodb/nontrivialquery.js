db.club_info.find({
    "id": {
        $gt: 0
    },
    "league": {
        "$in": ["Premier League", "Lega Serie A"]
    }
}).pretty();