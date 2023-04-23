![USAGE.PNG](https://github.com/Lukmanern/go-supabase/blob/master/media/usage.png)

# Go ToDoList with Supabase

A simple command-line todo list application written in Go, using a PostgreSQL database.

# Features :

0. Exit App
1. Create TODO
2. Read TODO (in 3 ways)
3. Update TODO (Edit)
4. Update TODO Status
5. SoftDelete TODO
6. Restore TODO (From SoftDelete)
7. Destroy (Delete Permanent)
8. Hard Reset (DROP TODO TABLE -> RECREATE TABLE)
9. Verify User Action (For Destroy and Hard Reset)

# Set Up Note

1. ENV File
   Copy .env.example -> .env

for unix, command :

```sh
cp .env.example .env
```

or create manual :

```sh
dbname=postgres
user=postgres
port=5432
host=MAYBE_SECRET
password=ALWAYS_SECRET
```

2. Variables

Loot at `func getENV()`, and don't forget to change the `path` value.

```sh
path := "C:/Users/Lenovo/go/src/DB_CLI/.env"
```

# Directory Structure

```sh
   ├── .gitignore
   ├── README.md
   ├── banner
   │   └── banner.go -> show banner
   ├── database
   │   ├── .env
   │   ├── .env.example
   │   ├── create-table.sql -> connection and create sql
   │   └── db.go
   ├── dev-note.md -> note for developers
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── error.go -> throw err and exit app
   ├── helper
   │   ├── input.go -> get user input
   │   └── options.go -> show menu/ options
   ├── main.go
   └── usage.png
```

![update.PNG](https://github.com/Lukmanern/go-supabase/blob/master/media/update.png)

# Supabase

Supabase is an open-source platform for building scalable web applications. It provides a set of tools and services for building modern, real-time applications, including a PostgreSQL database, a serverless functions platform, and a real-time API.

Supabase is designed to be easy to use and offers a number of features that make it a powerful tool for building web applications. It has a built-in API that allows you to interact with your database using simple HTTP requests, and it supports real-time updates using websockets. It also has a built-in functions platform that allows you to run serverless functions in response to events such as new data being added to your database.

In addition to these core features, Supabase also offers a number of other tools and services that can help you build better web applications. It has a powerful set of API functions that allow you to do things like search, sort, and filter data, and it has a number of integrations with popular third-party tools such as Stripe, Twilio, and Slack.

Overall, Supabase is a `powerful` and `flexible` platform for building modern web applications, and it is particularly well-suited for building real-time applications that need to handle large amounts of data and provide real-time updates to users.
