db.users.insertOne({
    name: "Alice",
    email: "alice@gmail.com",
    contact: {
        phone: "0987654321",
        address: "456 Wonderland Ave"
    },
    "job": {
        title: "CEO",
        company: "Microsoft"
    },
    hobbies: ["painting", "cycling", "traveling","mancing"],
    posts: [
        {
            title: "My Travel Adventures",
            content: "I love exploring new places and cultures."
        },
        {
            title: "Cycling Tips",
            content: "Here are some tips for safe and enjoyable cycling."
        }
    ]
})