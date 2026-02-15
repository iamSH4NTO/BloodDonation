# Blood Donor Management System - Implementation Plan

## Project Overview

A comprehensive web-based platform to manage blood donors, facilitate searches, and maintain donation records. This plan is based on the original proposal but streamlined (removing Redis and Notifications) and enhanced for modern best practices.

## 1. Tech Stack (Enhanced)

| Component      | Technology                      | Rationale                                                                       |
| -------------- | ------------------------------- | ------------------------------------------------------------------------------- |
| **Backend**    | **Golang** (Gin Framework)      | High performance, strong typing, excellent for REST APIs.                       |
| **Database**   | **MySQL** (Default)             | Primary DB, but interchangeable via GORM (PostgreSQL/SQLite supported).         |
| **ORM**        | **GORM**                        | **Critical**: Abstraction layer to ensure zero code changes when switching DBs. |
| **Frontend**   | **Vue.js 3** + **Vite**         | Modern, fast reactive frontend framework.                                       |
| **Styling**    | **TailwindCSS**                 | Utility-first CSS for rapid and consistent UI design.                           |
| **State Mgmt** | **Pinia**                       | The standard state management library for Vue 3.                                |
| **Auth**       | **JWT** (Access + Refresh)      | Stateless, secure authentication.                                               |
| **Deployment** | **Docker** + **Docker Compose** | Containerized environment for consistency.                                      |

> **Note:** Redis has been removed. Caching will be handled via database optimizations (indexes) and in-memory application caching where strictly necessary (e.g., local Go maps for static config).

## 2. Role-Based Access Control (RBAC)

- **Admin**: Full system access, User Management, Logs, System Analytics.
- **Manager**: Regional control (District/Thana level), Donor verification.
- **Donor**: Profile management, privacy settings, donation history.

## 3. Core Features & Enhancements

### A. Donor Management & Search

- **Public Search**: Search by Blood Group, District, Thana.
  - _Enhancement_: Implement **Pagination** and **Filtering** (e.g., detailed availability status).
- **Geo-Location**: Calculate distance using Haversine formula (implemented in Postgres or Go) to sort donors by proximity.
- **Privacy**: Donors can toggle their visibility or partial contact info.
- **Cooldown Logic**: Automated calculation of "Next Eligible Date" (3-4 months) based on last donation.

### B. Security

- **Authentication**: JWT-based (Access Token: 15min, Refresh Token: 7days).
- **Password/Sensitive Data**: bcrypt hashing for passwords.
- **Rate Limiting**: Implemented via **Golang middleware** (using `x/time/rate`) instead of Redis.
- **Audit Logging**: Track who viewed whose contact number (Anti-Scraping measure).

### C. Logging & Analytics

- **Search Logs**: Record what was searched and when.
- **View Logs**: Record who viewed a donor's contact details to prevent abuse.

## 4. Database Schema (Enhanced GORM Models)

### `User`

- `ID` (UUID/UInt)
- `Email` (Unique, Indexed)
- `PasswordHash`
- `Role` (Enum: Admin, Manager, Donor)
- `IsActive` (Boolean)
- `CreatedAt`, `UpdatedAt`

### `DonorProfile` (One-to-One with User)

- `UserID` (FK)
- `Name`
- `BloodGroup` (Indexed)
- `Phone`
- `District` (Indexed)
- `City` (New)
- `AreaVillage` (New)
- `PostalCode` (Optional)
- `Latitude`, `Longitude` (Optional - For Geo-calculation)
- `GoogleMapLink` (Optional)
- `LastDonationDate` (Nullable)
- `IsAvailable` (Boolean - Global toggle)
- `PrivacySettings` (JSONB - Granular control)

### `DonationHistory`

- `ID`
- `DonorID` (FK)
- `DonationDate`
- `Location` / `Hospital`
- `Notes`

### `PhoneGroupViewLog`

- `ID`
- `ViewerID` (FK)
- `TargetDonorID` (FK)
- `Timestamp`
- `IPAddress`

## 5. API Structure (RESTful)

### Auth

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/refresh`

### Donors

- `GET /api/v1/donors` (Search with query params: `group`, `lat`, `long`, `dist`)
- `GET /api/v1/donors/:id` (Public details)
- `GET /api/v1/donors/:id/contact` (Protected; triggers View Log)
- `PUT /api/v1/profile` (Update own profile)

### Admin/Manager

- `GET /api/v1/admin/users`
- `GET /api/v1/admin/logs/views` (Analyze suspicious activity)

## 6. Development Roadmap

### Phase 1: Foundation (MVP)

- Setup Project Structure (Go + Vue Monorepo or separate).
- Database Design & Migration Setup.
- Authentication System (Register/Login/JWT).
- Basic Donor Profile CRUD.

### Phase 2: Search & Privacy

- Advanced Search Logic (Blood Group + Location).
- Phone Number View Logging.
- Privacy controls implementation.

### Phase 3: Geo & Optimization

- Geo-location sorting.
- Donation History & Cooldown Logic.
- Unit Testing & Deploy configurations.

## 7. Scalability & Performance without Redis

Since Redis is removed, we ensure performance through:

1.  **Database Indexing**: Aggressive indexing on `BloodGroup`, `District`, `Thana`.
2.  **Connection Pooling**: Proper configuration of `pgx` pool in GORM.
3.  **In-Memory Caching (Optional)**: For very static data (like list of Districts), use a global Go variable/map with a mutex, refreshed on app restart.
4.  **Stateless Backend**: The Go backend remains stateless (except for the minor in-memory config), allowing horizontal scaling behind a Load Balancer if needed.

## Verification Plan

### Automated Tests

- **Backend**: Use standard Go `testing` package.
  - Test Auth flows (Register login).
  - Test Search logic (ensure filters work).
  - Test Cooldown logic calculation.
- **Frontend**: Component testing with Vitest.

### Manual Verification

1.  **Register** two users (Donor A, Viewer B).
2.  **Search**: As Viewer B, search for Donor A's blood group.
3.  **View Contact**: Click to view contact, verify a log entry is created in `PhoneGroupViewLog`.
4.  **Cooldown**: Set Donor A's last donation date to 1 month ago -> Verify they do not appear in "Eligible" search results (if implemented) or show as "Unavailable".
