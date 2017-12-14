    db.app_stats.aggregate([
        {
            $match: {
                "end_at": {
                    "$gte":  ISODate("2017-08-17T08:00:03.626Z")
                }
            }
        },

        {
            $project: {
                "app_id": "$app_id",
                "count": "$count",
                "interval": {
                                "$floor": {
                                    "$divide": [
                                        {"$subtract": ["$end_at", ISODate("2017-08-17T08:00:03.626Z")]},
                                        300000
                                    ]
                                }
                            }
            }
        },

        {
            $group: {
                "_id": {
                    "app_id": "$app_id",
                    "interval": "$interval"
                },
                "total": {$sum: "$count"}
            }
        },

        {
            $sort: {
                "_id.interval": -1
            }
        }
    ])



        db.app_stats.aggregate([
        {
            $match: {
                "app_id": "9d2e8364"
            }
        },

        {
            $project: {
                "app_id": "$app_id",
                "count": "$count",
            }
        },

        {
            $group: {
                "_id": {
                    "app_id": "$app_id",
                },
                "total": {$sum: "$count"}
            }
        }
    ])