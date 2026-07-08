# 🚗 Parkora

**Parkora** is a smart parking management system built with **Go** that helps users find, reserve, and manage parking spaces efficiently. The system provides secure authentication, role-based authorization, parking zone management, and reservation handling.

---

# 🚀 Features

## 🔐 Authentication & Authorization

- User registration and login
- JWT-based authentication
- Role-based access control
- Secure password hashing using bcrypt
- Protected routes with middleware

---

## 🅿️ Parking Zone Management

- Create parking zones
- Update parking zones
- Delete parking zones
- View all available parking zones
- Track parking capacity dynamically
- Support multiple vehicle types:

  - 🚗 Car
  - 🏍️ Bike
  - 🚚 Truck
  - 🚌 Bus

---

## 📅 Reservation Management

- Create parking reservations
- View reservations
- Update reservation details
- Cancel reservations
- Manage reservation status

Reservation statuses:

```
active
cancelled
completed
```

---

# 🛠️ Technology Stack

## Backend

- Go (Golang)
- Echo Framework
- GORM ORM
- PostgreSQL
- JWT Authentication
- Go Playground Validator

## Development Tools

- Air (Live Reload)
- Git
- GitHub
- Postman

---

# 📂 Project Structure

```
Parkora/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   ├── user/
│   ├── parking/
│   ├── reservations/
│   ├── middleware/
│   ├── config/
│   └── response/
│
├── .env
├── go.mod
├── go.sum
├── air.toml
└── README.md
```

---

# ⚙️ Installation & Setup

## 1. Clone Repository

```bash
git clone https://github.com/yourusername/parkora.git
```

Go to project directory:

```bash
cd parkora
```

---

## 2. Install Dependencies

```bash
go mod download
```

---

## 3. Environment Configuration

Create a `.env` file in the root directory.

```env
APP_PORT=5000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=parkora

JWT_SECRET=your_secret_key
```

---

## 4. Run Application

### Using Go

```bash
go run ./cmd/server
```

### Using Air

```bash
air
```

Server will start:

```
http://localhost:5000
```

---

# 🔑 User Roles

## 👤 Driver

Drivers can:

- Register an account
- Login
- View their own reservations
- Create reservations
- Update reservation information
- Cancel their own reservations


Restrictions:

- Cannot update other users' reservations
- Cannot change parking zone during reservation update

---

## 👨‍💼 Admin

Admins can:

- Manage parking zones
- View all reservations
- Update any reservation
- Modify reservation status
- Manage system resources

---

# 📌 API Endpoints

## Authentication

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/auth/register` | Register user |
| POST | `/api/v1/auth/login` | Login user |

---

# Parking Zones

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/parking-zones` | Create parking zone |
| GET | `/api/v1/parking-zones` | Get all parking zones |
| GET | `/api/v1/parking-zones/:id` | Get single parking zone |
| PUT | `/api/v1/parking-zones/:id` | Update parking zone |
| DELETE | `/api/v1/parking-zones/:id` | Delete parking zone |

---

# Reservations

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/reservations` | Create reservation |
| GET | `/api/v1/reservations` | Get reservations |
| GET | `/api/v1/reservations/:id` | Get reservation |
| PATCH | `/api/v1/reservations/:id` | Update reservation |
| DELETE | `/api/v1/reservations/:id` | Cancel reservation |

---

# 🧠 Business Rules

- Drivers can only manage their own reservations.
- Admins can manage all reservations.
- Cancelled reservations cannot be modified by drivers.
- Admins can modify cancelled reservations.
- Completed reservation status cannot be changed.
- Parking availability is calculated dynamically.
- A reservation requires an available parking spot.
- User permissions are controlled using JWT roles.

---

# 🔒 Security

- Password hashing with bcrypt
- JWT authentication
- Role-based authorization
- Request validation
- Protected API routes
- Secure database queries using GORM

---

# 🧪 API Testing

Recommended tools:

- Postman
- Thunder Client

---

# 📈 Future Improvements

- Online payment integration
- QR code based parking entry system
- Real-time parking availability
- Notification system
- Booking history
- Mobile application
- GPS-based parking search

---

# 👨‍💻 Author

**Opi Korim**

Backend Developer | Go Developer

---

# 📄 License

This project is licensed under the MIT License.