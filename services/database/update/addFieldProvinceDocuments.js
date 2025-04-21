use("nuky_db");

db.getCollection("provinces").updateMany(
  {},
  {
    $set: {
      destroymentRound: -1,
      updatedDate: null,
      deletedDate: null
    }
  }
);
