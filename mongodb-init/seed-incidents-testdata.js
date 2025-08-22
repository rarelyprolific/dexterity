// Create a mythical software system which we'll define incidents and tasks for in dexterity.
// This will be the basis of test data to prove the design of the domain model.

// Incident statuses
// no status - Unstarted
// ACTIVE - In Progress
// PAUSED - Paused
// ABANDONED - Abandoned
// SOLVED - Resolved 

db.incidents.deleteMany({})

db.incidents.insertMany([
    { summary: "Something bad happened.", createdBy: "PlayerOne", createdOn: new Date() },
    { summary: "User cannot log in.", createdBy: "PlayerTwo", createdOn: new Date() },

    {
        summary: "Multiple users cannot log in",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        status: "ACTIVE"
    }
])

// Task statuses
// no status - Not yet fully defined or has questions
// READY - Ready to Start
// ACTIVE - In Progress
// PAUSED - Paused
// COMPLETE - 

db.tasks.deleteMany({})

db.tasks.insertMany([
    {
        summary: "Improve crash reporting.",
        createdBy: "PlayerTwo",
        createdOn: new Date(),
        status: "READY"
    }
])
