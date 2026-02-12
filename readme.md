# Notes App (Go + React) — Learning Project

**Notes App** is a simple full-stack project built as a personal learning challenge.

My main goal is to improve my backend skills using **Go (Golang)** by building a real-world style REST API with authentication and CRUD features.  
To keep the project realistic and motivating, the frontend is built separately using a modern **React** UI.

This project is intentionally designed like a real team workflow:
- **Frontend** focuses on UI/UX, routing, state, and API integration readiness.
- **Backend** focuses on API design, authentication, security, and database logic.

---
### Tech stack
![React](https://img.shields.io/badge/React-20232A?logo=react&logoColor=61DAFB)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?logo=typescript&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-00ADD8?logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)

---
## Goals

### Primary Goal (Backend)
- Practice building a complete backend API in Go
- Learn how to design API structure properly (routes, request/response, errors)
- Implement authentication flows (register/login/forgot password/reset password)
- Build a Notes CRUD system with authorization per user
- Improve skills in:
  - middleware
  - validation
  - security basics
  - database schema design
  - clean project structure

### Secondary Goal (Frontend)
- Learn React by integrating it with a real backend later
- Build a clean UI that is ready to connect to an API
- Practice common frontend patterns:
  - protected routes
  - form validation
  - loading/error/empty states
  - component reuse
  - service layer abstraction

---

## Application Features

### Authentication
- Register new account
- Login
- Logout
- Forgot password (send reset link)
- Reset password (using token)

### Notes Management
- Create note
- View all notes
- Search notes
- Edit note
- Delete note (with confirmation)

---

## Application Specifications

### Frontend (React)
Frontend is designed to be API-ready, but initially uses mock services.

#### Pages
- Login
- Register
- Forgot Password
- Reset Password
- Notes Dashboard (List notes + Search)
- Create Note
- Edit Note
- Settings (Profile + Logout)

#### UI/UX Requirements
- Responsive design (mobile-first)
- Clean and minimal interface
- Toast notifications for user actions
- Validation for all forms
- Loading, error, and empty states
- Delete confirmation modal

#### Architecture (Frontend)
Frontend uses a clean folder structure:
- `pages/` — route-level pages
- `components/` — reusable UI components
- `layouts/` — AuthLayout, AppLayout
- `services/` — API service layer (mock first, real API later)
- `types/` — TypeScript types
- `hooks/` — custom hooks (auth, notes, etc.)
- `utils/` — helpers

---

### Backend (Go)
Backend is the main learning target of this project.

#### Planned Backend Responsibilities
- User authentication
- Password hashing
- Password reset flow with token + expiry
- Notes CRUD operations (per-user ownership)
- Secure route protection with middleware
- Input validation
- Consistent error responses
- Database integration

#### Planned Endpoints (Not Final Yet)
This project intentionally avoids locking the backend into a fixed contract too early.
However, the backend will eventually expose endpoints similar to:

- Auth:
  - register
  - login
  - forgot password
  - reset password
  - logout (optional)

- Notes:
  - list notes
  - get note by id
  - create note
  - update note
  - delete note

---

## Project Philosophy

This repository is built with a "learning-by-doing" approach.

- The frontend is treated like a separate teammate.
- The backend API is designed and implemented manually.
- Mock services are used first to avoid forcing backend decisions too early.
- Once backend is ready, mock services will be replaced with real API calls.

---

## Tech Stack

### Frontend
- React
- TypeScript
- React Router
- Context API (Auth state)
- Mock services (localStorage/in-memory)

### Backend (Planned)
- Go (Golang)
- REST API
- Database (TBD)
- JWT or session-based auth (TBD)

---

## How the App Works (High-Level)

1. User registers or logs in
2. User is redirected to the Notes Dashboard
3. User can:
   - create notes
   - edit notes
   - delete notes
   - search notes
4. Notes are always owned by the logged-in user
5. If user forgets password:
   - user requests reset link
   - user resets password using token

---

## Future Improvements (Optional)
- Pagination for notes list
- Tags or categories
- Rich text editor
- Note pinning
- Dark mode
- Refresh tokens
- Email sending service integration

---

## Status
This project is actively developed as a learning challenge.

Frontend is built first using mock data, then later integrated with a Go backend API.

---
