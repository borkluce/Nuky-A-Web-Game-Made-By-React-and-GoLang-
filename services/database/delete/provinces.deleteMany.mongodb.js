use("nuky_db")

db.provinces.updateMany({}, { $unset: { id: "" } })
