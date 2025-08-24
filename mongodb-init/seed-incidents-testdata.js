// Create a mythical software system which we'll define incidents and tasks for in dexterity.
// This will be the basis of test data to prove the design of the domain model.

// Incident statuses
// no status - Unstarted
// ACTIVE - In Progress
// PAUSED - Paused
// ABANDONED - Abandoned
// SOLVED - Resolved 

// It is valid for questions be added to an on-going incident to learn more about what is going on.

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
                        answededOn: new Date()
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
                        answededOn: new Date()
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
                        answededOn: new Date()
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

// Task statuses
// no status - Not yet fully defined or has questions
// READY - Ready to Start
// ACTIVE - In Progress
// PAUSED - Paused
// COMPLETE - 

// If there are unanswered questions on a task before it is started it cannot be READY.
// If there are questions on a tasks after it has been active it may be PAUSED.

// Define three tasks.

db.tasks.deleteMany({})

db.tasks.insertMany([
    {
        summary: "Create a splash screen whilst 'Echoes of Logic' is loading",
        status: "READY",
        createdBy: "PlayerTwo",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        description: "The game client currently loads and connects before any UI is shown to the user. We need to add a simple splash screen to communicate to the end user that the game is loading.",
        questions: [
            {
                question: "Do we need the graphics team to supply a design for the splash screen?",
                askedBy: "PlayerOne",
                askedOn: new Date(),
                answer: "Not for the initial version. We only need something simple to reduce the amount of support incidents. Just use the main EOL logo from the main menu on a black background.",
                answeredBy: "PlayerOne",
                answeredOn: new Date()
            },
            {
                question: "So, no loading progress bar or loading percentage either for the time being?",
                askedBy: "PlayerOne",
                askedOn: new Date(),
                answer: "No, not needed for version one. But good idea! We'll add it in at a later date via a new task.",
                answeredBy: "PlayerOne",
                answeredOn: new Date()
            }
        ],
        justification: "There is a long (30-60 second) pause whilst the game client is loading and connecting to login servers. It can look like it is not loading or working.",
        consequences: "Support are already getting a high volume of incidents from users who believe the game is broken. We have to keep telling them to just wait a little longer for the game to load and it'll work."
    },

    {
        summary: "We need to figure out high-level dynamic world events for next season",
        createdBy: "PlayerTwo",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        description: "Not all finer details are needed yet but we do need to plan which zones will be affected and which NPCs will be corrupted as an initial design document.",
        questions: [
            {
                question: "If we are planning to recycle event code/procedures from previous events: Do we just need design team on this initially instead of development? But can bring them in later.",
                askedBy: "PlayerOne",
                askedOn: new Date(),
            },
        ],
        justification: "Every new season of the game requires completely new dynamic events. It is the main selling point of the game and is expected by users.",
        consequences: "We can't start planning out development work on the events until we have the high-level plan. If we leave it too long it will impact the release date of the next season!"
    },

    {
        summary: "Improve game client crash reporting to enhance incidents being reported to support",
        status: "COMPLETE",
        createdBy: "PlayerTwo",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        description: "We already have the error codes and debug information in the client but we don't expose it. Add a simple crash reporter dialog which shows errors to user when main client process fails. Include a button which will auto-generate an email to support with details embedded.",
        justification: "We are getting anecdotal reports of crashes but no real details we can act upon. We need more information to know the scale of any potential problem.",
        consequences: "Support are logging regular incidents and escalating them to development. None of us can resolve but we are spending a disproportionate amount of time trying to investigate with nothing tangible to act upon. It is starting to significantly impact our ability to work on other tasks."
    }
])



// WIKI - documenting the system

db.wiki.deleteMany({})

db.wiki.insertMany([
    {
        title: "The 'Echos of Logic' development wiki",
        shortlink: "MAIN-MENU",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        text: "This tells you everything!\n\n If it doesn't, add it!\n\n Some text in **markdown** format. Some menu options!",
        tags: [
            "General"
        ]
    },

    {
        title: "What is 'Echos of Logic?'",
        shortlink: "GAME-OVERVIEW",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        text: "# Overview\n\nIt's a major MMORPG with the player base fighting co-operatively against a sentient AI. Also known as EOL or 'End Of Life'",
        tags: [
            "Design"
        ]
    },

    {
        title: "'UserMaintV1' should still be used for early-access users",
        shortlink: "EARLY-USERS-MAINT-ISSUE",
        createdBy: "PlayerOne",
        createdOn: new Date(),
        lastUpdatedBy: "PlayerOne",
        lastUpdatedOn: new Date(),
        text: "The 'UserPortal' is the main facility for maintaining most users but is new and we still have to iron out some quirks to make it work for very early users. If you get an error with 'UserPortal' but the user was part of the initial beta, try using 'UserMaintV1' for the time being. If this doesn't work, it's likely a new problem so raise it as a new incident.",
        tags: [
            "Support"
        ]
    }
])
