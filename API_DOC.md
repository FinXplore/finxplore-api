# FinXplore API Documentation

Base URL: `/api/v1`

## Authentication

### 1. Register User
**Endpoint:** `POST /api/v1/auth/register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123",  // Min 6 characters
  "first_name": "John",
  "last_name": "Doe",
  "dialing_code": "+1",
  "mobile_number": "1234567890",
  "currency": "USD"
}
```

**Response (201 Created):**
```json
{
  "full_name": "John Doe",
  "role": "user"
}
```

**Error (400/409):**
```json
{
  "error": "Error message details"
}
```

---

### 2. Login User
**Endpoint:** `POST /api/v1/auth/login`

**Request Body:**
```json
{
  "dialing_code": "+1",
  "mobile_number": "1234567890",
  "password": "securepassword123"
}
```

**Response (201 Created):**
```json
{
  "full_name": "John Doe",
  "role": "user",
  "user_code": "USR-UUID-123",
  "access_token": "jwt.access.token...",
  "refresh_token": "jwt.refresh.token..."
}
```

---

### 3. Refresh Access Token
**Endpoint:** `POST /api/v1/auth/refresh`

**Request Body:**
```json
{
  "dialing_code": "+1",
  "mobile_number": "1234567890",
  "refresh_token": "jwt.refresh.token..."
}
```

**Response (200 OK):**
```json
{
  "access_token": "new.jwt.access.token..."
}
```

---

### 4. Get Current User Profile (Protected)
**Endpoint:** `GET /api/v1/auth/me`

**Headers:**
`Authorization: Bearer <access_token>`

**Response (200 OK):**
```json
{
  "message": "You are authorized!",
  "user_code": "USR-UUID-123",
  "role": "user"
}
```

---

## Stock Market Data
All stock routes are protected and require the `Authorization` header.

### 1. Search Stocks
**Endpoint:** `GET /api/v1/stock/search`

**Query Parameters:**
- `q`: Search query string (e.g., "Apple")

**Response (200 OK):**
```json
[
  {
    "ticker": "AAPL",
    "name": "Apple Inc.",
    "type": "EQUITY",
    "exchange": "NASDAQ"
  },
  ...
]
```

---

### 2. Get Stock Details
**Endpoint:** `GET /api/v1/stock/:ticker`

**Path Parameters:**
- `ticker`: Stock symbol (e.g., "AAPL")

**Response (200 OK):**
```json
{
  "ticker": "AAPL",
  "name": "Apple Inc.",
  "exchange": "NASDAQ",
  "currency": "USD",
  "sector": "Technology",
  "industry": "Consumer Electronics",
  "description": "Apple Inc. designs, manufactures, and markets...",
  "website": "https://www.apple.com",
  "market_cap": 3000000000000,
  "pe_ratio": 28.5,
  "dividend_yield": 0.55,
  "last_updated": "2024-03-20T10:00:00Z"
}
```

**Error (404):**
```json
{
  "error": "Stock not found"
}
```

---

### 3. Get Stock Chart Data
**Endpoint:** `GET /api/v1/stock/:ticker/chart`

**Path Parameters:**
- `ticker`: Stock symbol (e.g., "AAPL")

**Query Parameters:**
- `period`: Time range (default "1m"). Allowed values depend on provider (e.g., "1d", "1mo", "1y").

**Response (200 OK):**
```json
[
  {
    "date": "2024-03-19",
    "open": 170.5,
    "high": 175.2,
    "low": 170.0,
    "close": 174.8,
    "volume": 50000000
  },
  ...
]
```

---

## Health Check
**Endpoint:** `GET /health`

**Response (200 OK):**
```json
{
  "status": "up"
}
```
