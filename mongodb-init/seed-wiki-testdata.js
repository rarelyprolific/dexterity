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
