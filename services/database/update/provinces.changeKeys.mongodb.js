use("nuky_db")

db.provinces.find({}).forEach(function (doc) {
    db.provinces.updateOne(
        { _id: doc._id },
        {
            $unset: { attack_count: "", support_count: "" },
            $set: {
                attackCount: doc.attack_count,
                supportCount: doc.support_count,
            },
        }
    )
})
