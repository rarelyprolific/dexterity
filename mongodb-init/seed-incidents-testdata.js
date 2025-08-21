// Create a mythical software system which we'll define incidents and tasks for in dexterity.
// This will be the basis of test data to prove the design of the domain model.

db.incidents.deleteMany({})

db.incidents.insertMany([
    { summary: "Something bad happened.", createdBy: "PlayerOne", createdOn: new Date() },
    { summary: "User cannot log in.", createdBy: "PlayerTwo", createdOn: new Date() }
])
