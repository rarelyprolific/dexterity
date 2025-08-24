// Task statuses
// no status - Not yet fully defined or has questions
// READY - Ready to Start
// ACTIVE - In Progress
// PAUSED - Paused
// COMPLETE - 

// If there are unanswered questions on a task before it is started it cannot be READY.
// If there are questions on a tasks after it has been active it may be PAUSED.

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
