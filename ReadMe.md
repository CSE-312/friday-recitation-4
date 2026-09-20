# Recitation 4: Database Debugging Activity

> **Objective:** Debug a small Go and PostgreSQL application by applying the HW1 concepts covered in recent recitations.

---

## Activity Overview

| Item | Details |
| --- | --- |
| Topic | Go server, Docker Compose, PostgreSQL, and environment variables |
| Goal | Find and fix the bugs preventing the app from running correctly |
| Bug count | 8 total bugs |
| Submission | Upload a `.zip` file to Autolab |

You may use recitation slides, your notes, previous examples, and Jesse's slides while completing this activity.

---

## Files to Inspect

The 8 bugs are located across the following files:

| File | Purpose |
| --- | --- |
| `.env` | Local environment variables |
| `database.go` | Database connection and setup logic |
| `docker-compose.yaml` | Container configuration |
| `Dockerfile` | App image setup |
| `main.go` | Server and request-handling logic |

---

## Getting Started

Clone the repo, open it in the IDE of your choice, and move into this activity directory.

Then initialize the Go module and install the required dependencies:

```bash
go mod init RecActivity4
go mod tidy
```

---

## Running the Server

Build and start the app with Docker Compose:

```bash
docker compose up --build --force-recreate
```

If you need to reset the database and start with a fresh one, run:

```bash
docker compose up --build --force-recreate --renew-anon-volumes
```

> **Database reset warning**
>
> The reset command clears the anonymous database volume. Only use it when you intentionally want to remove the existing database data.

---

## Testing the Server

Open a second terminal while the server is running.

### Add a Row

Use the `/Add` route to insert a value into the database:

```bash
curl http://localhost:8080/Add/<any-url-safe-text>
```

Example:

```bash
curl http://localhost:8080/Add/hello
```

### View All Rows

Use the `/` route to view all rows currently stored in the database:

```bash
curl http://localhost:8080/
```

---

## Protected Code

The `main.go` file contains three functions that do **not** include any bugs:

- `PrintStartUpMessage`
- `handler`
- `handle_response`

> Do not change the code inside these functions.

---

## Outcome

Once all eight bugs are corrected, your app should be able to:

- Both `Docker Containers` successfully create when using the docker compose command
- Use the `/Add` route to insert a value into the database.
- Use the `/` route to view all rows currently stored in the database.
- Store data in the database so that it persists across container creation and destruction.

---

## Submission

Submit your work when your code compiles and is bug-free, or when the recitation period ends.

Upload the completed project as a `.zip` file to Autolab.
