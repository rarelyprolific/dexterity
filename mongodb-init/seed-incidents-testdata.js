// Incident statuses
// no status - Unstarted
// ACTIVE - In Progress
// PAUSED - Paused
// ABANDONED - Abandoned
// SOLVED - Resolved 

db.incidents.deleteMany({})

db.incidents.insertMany([
    {
        summary: "Some game clients are failing during the auto-update process?",
        status: "ACTIVE",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        log: [
            {
                description: "The installer appears to start and partially attempts the update process but fails with a 'cannot find eoldat3c.part' file' error.",
                createdBy: "PlayerOne",
                createdOn: new Date(),
            },
            {
                description: "The previous version of this installer didn't write this datafile. It is possible it is trying to delete a file which doesn't exist. Checking!",
                createdBy: "PlayerTwo",
                createdOn: new Date(),
            }
        ]
    },

    {
        summary: "Multiple users cannot log in",
        status: "SOLVED",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        log: [
            {
                description: "The login issues seem to be occurring across all global regions",
                createdBy: "PlayerOne",
                createdOn: new Date(),
                questions: [
                    {
                        question: "Approximately when is the earliest known instance of this problem occurring?",
                        askedBy: "PlayerOne",
                        askedOn: new Date(),
                        answer: "We logged a report from a user at around 3am this morning.",
                        answeredBy: "PlayerOne",
                        answeredOn: new Date()
                    },
                    {
                        question: "Is it ALL users or are we seeing SOME successful logins?",
                        askedBy: "PlayerOne",
                        askedOn: new Date(),
                    }
                ]
            },
            {
                description: "There is a dramatic reduction in the amount of reports about this issue.",
                createdBy: "PlayerOne",
                createdOn: new Date(),
            }
        ],
        resolution: {
            description: "There was a temporary network outage affecting traffic from all regions which is now resolved.",
            resolvedBy: "PlayerOne",
            resolvedOn: new Date(),
        }
    },

    {
        summary: "Unable to reset login credentials for early-access users.",
        status: "SOLVED",
        createdBy: "PlayerTwo",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        log: [
            {
                description: "User minotaur75 cannot self-reset password. Also failing with 'Error 34cfa' when support staff attempt to reset on behalf of user",
                createdBy: "PlayerOne",
                createdOn: new Date(),
                questions: [
                    {
                        question: "There are two versions of user maintenance screens in use. Did you attempt to reset user password using 'UserMaintV1' or 'UserPortal'?",
                        askedBy: "PlayerOne",
                        askedOn: new Date(),
                        answer: "We tried using 'UserPortal' because it was the only method we were aware of.",
                        answeredBy: "PlayerOne",
                        answeredOn: new Date()
                    }
                ]
            },
            {
                description: "Attempt to reset the user password using 'UserMaintV1'. There is a known issue with early users and the new 'UserPortal'.",
                createdBy: "PlayerOne",
                createdOn: new Date(),
                questions: [
                    {
                        question: "Can you confirm if this works so we know if it is the same issue?",
                        askedBy: "PlayerOne",
                        askedOn: new Date(),
                        answer: "Yes. Reset minotaur75's password successfully using 'UserMaintV1'.",
                        answeredBy: "PlayerOne",
                        answeredOn: new Date()
                    }
                ]
            }
        ],
        resolution: {
            description: "This is a known issue but has a workaround. Added ['EARLY-USERS-MAINT-ISSUE'] page to wiki for support.",
            resolvedBy: "PlayerOne",
            resolvedOn: new Date(),
        }
    }
])
