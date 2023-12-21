# PoulTracker

## Description

PoulTracker is a minimalist web application designed for individuals who raise chickens. It provides an intuitive and easy-to-use platform to track the daily egg production of chickens. This application is perfect for those seeking a straightforward method to record and analyze their chickens' egg-laying patterns.

## Key Feature

- Daily Recording: Enables users to log the total number of eggs laid each day.

## Technology Stack

- **Backend**: Go
- **Frontend**: JavaScript & Tailwind CSS
- **Database**: SQLite

## Roadmap

- [ ] Addition of data visualizations
- [ ] Purchasing tracking
- [ ] Setting the number of hens
- [ ] Multi-users
- [ ] HTMX integration
- [ ] Calendar vue
- [ ] Water tracking
- [ ] Medication tracking
- [ ] PWA
- [ ] Reminders
- [ ] Notifications

## Configuration

| Environment Variable      | Default          | Comment                                          |
| ------------------------- | ---------------- | ------------------------------------------------ |
| POULTRACKER_HOST          | "localhost"      | Server Host                                      |
| POULTRACKER_PORT          | 8080             | Server Port                                      |
| POULTRACKER_DATABASE_PATH | "poultracker.db" | Relative or absolut path for the SQLite database |
| POULTRACKER_AUTH_USERNAME | "admin"          | AUTH username                                    |
| POULTRACKER_AUTH_PASSWORD | "password"       | AUTH password                                    |
