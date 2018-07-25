db.createUser(
    {
        user: "email_sender",
        pwd: "password",
        roles: [ { role: "readWrite", db: "email_sender" } ]
    }
)