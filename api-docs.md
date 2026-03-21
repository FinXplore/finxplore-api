# FinXplore API Documentation

> **Base URL:** `/api/v1`
> **Authentication:** Protected routes require `Authorization: Bearer <access_token>` header.

---

## 🔐 Auth Routes (`/auth`)

### 1. Register User
- **Endpoint:** `POST /auth/register`
- **Status:** ✅ Exists
- **Description:** Register a new user and retrieve tokens.
- **Request JSON:**
```json
{
  "email": "user@example.com",
  "password": "strongPassword123",
  "first_name": "John",
  "last_name": "Doe",
  "dialing_code": "+91",
  "mobile_number": "9876543210",
  "currency": "INR"
}
```
- **Response JSON (201 Created):**
```json
{
  "full_name": "John Doe",
  "role": "user",
  "user_code": "USR123",
  "access_token": "ey...",
  "refresh_token": "ey..."
}
```

### 2. Login User
- **Endpoint:** `POST /auth/login`
- **Status:** ✅ Exists
- **Description:** Login and receive JWT tokens.
- **Request JSON:**
```json
{
  "email": "user@example.com",
  "password": "strongPassword123"
}
```
- **Response JSON (201 Created):**
```json
{
  "full_name": "John Doe",
  "role": "user",
  "user_code": "USR123",
  "access_token": "ey...",
  "refresh_token": "ey..."
}
```

### 3. Refresh Token
- **Endpoint:** `POST /auth/refresh`
- **Status:** ✅ Exists
- **Description:** Get a new access token using a valid refresh token.
- **Request JSON:**
```json
{
  "email": "user@example.com",
  "refresh_token": "ey..."
}
```
- **Response JSON (200 OK):**
```json
{
  "access_token": "ey..."
}
```

### 4. Logout User (Protected)
- **Endpoint:** `POST /auth/logout`
- **Status:** ✅ Exists
- **Description:** Invalidate session.
- **Request Body:** None
- **Response JSON (200 OK):**
```json
{
  "message": "logged out successfully"
}
```

### 5. Get Current User Profile (Protected)
- **Endpoint:** `GET /auth/me`
- **Status:** ✅ Exists
- **Request Body:** None
- **Response JSON (200 OK):**
```json
{
  "user_code": "USR123",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "avatar_url": "https://...",
  "currency": "INR",
  "role": "user",
  "risk_tolerance": "moderate"
}
```

### 6. Update User Profile (Protected)
- **Endpoint:** `PUT /auth/me`
- **Status:** ✅ Exists
- **Request JSON:**
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "avatar_url": "https://...",
  "currency": "USD"
}
```
- **Response JSON (200 OK):**
```json
{
  "user_code": "USR123",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "avatar_url": "https://...",
  "currency": "USD",
  "role": "user",
  "risk_tolerance": "moderate"
}
```

### 7. Change Password (Protected)
- **Endpoint:** `POST /auth/change-password`
- **Status:** ✅ Exists
- **Request JSON:**
```json
{
  "current_password": "oldPassword123",
  "new_password": "newStrongPassword!#@"
}
```
- **Response JSON (200 OK):**
```json
{
  "message": "password updated successfully"
}
```

---

## 📈 Stock Routes (`/stock`)

### 8. Search Stocks (Protected)
- **Endpoint:** `GET /stock/search?q=Apple`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "ticker": "AAPL",
    "name": "Apple Inc.",
    "type": "EQUITY",
    "exchange": "NASDAQ"
  }
]
```

### 9. Get Stock Details (Protected)
- **Endpoint:** `GET /stock/:ticker`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
{
  "ticker": "AAPL",
  "name": "Apple Inc.",
  "exchange": "NASDAQ",
  "currency": "USD",
  "sector": "Technology",
  "industry": "Consumer Electronics",
  "description": "Apple Inc. designs, manufactures...",
  "website": "http://www.apple.com",
  "market_cap": 2500000000000,
  "pe_ratio": 25.5,
  "dividend_yield": 0.5,
  "last_updated": "2024-03-20T10:00:00Z"
}
```

### 10. Get Stock Chart Data (Protected)
- **Endpoint:** `GET /stock/:ticker/chart?period=1d`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "date": "2024-01-01",
    "open": 150.0,
    "high": 155.0,
    "low": 149.0,
    "close": 154.5,
    "volume": 1200000
  }
]
```

### 11. Get Stock Quote (Protected)
- **Endpoint:** `GET /stock/:ticker/quote`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
{
  "ticker": "AAPL",
  "price": 174.32,
  "change": 1.45,
  "change_pct": 0.84,
  "volume": 52000000,
  "last_updated": "2024-03-20T10:00:00Z"
}
```

### 12. Get Stock Indicators
- **Endpoint:** `GET /stock/:ticker/indicators?type=RSI&period=14`
- **Status:** 🔨 Build

### 13. Get Stock News (Protected)
- **Endpoint:** `GET /stock/:ticker/news?limit=20`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "headline": "Apple announces new product",
    "source": "Bloomberg",
    "url": "https://...",
    "published_at": "2024-03-20T08:00:00Z"
  }
]
```

### 14. Get Stock Financials (Protected)
- **Endpoint:** `GET /stock/:ticker/financials?type=income`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "period_end_date": "2023-12-31",
    "metrics": {
      "netIncome": 25000000,
      "revenue": 100000000
    }
  }
]
```

---

## 🌍 Market Routes (`/market`)

### 15. Get Market Indices (Protected)
- **Endpoint:** `GET /market/indices`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "name": "NIFTY 50",
    "value": 22000.50,
    "change": 150.25,
    "change_pct": 0.68,
    "last_updated": "2024-03-20T10:00:00Z"
  }
]
```

### 16. Get Market Gainers (Protected)
- **Endpoint:** `GET /market/gainers?exchange=NSE`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "ticker": "RELIANCE",
    "name": "Reliance Industries",
    "price": 2980.00,
    "change": 120.00,
    "change_pct": 4.19,
    "volume": 1500000
  }
]
```

### 17. Get Market Losers (Protected)
- **Endpoint:** `GET /market/losers?exchange=NSE`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "ticker": "TCS",
    "name": "Tata Consultancy Services",
    "price": 3950.00,
    "change": -50.00,
    "change_pct": -1.25,
    "volume": 800000
  }
]
```

### 18. Get Market Volume Leaders (Protected)
- **Endpoint:** `GET /market/volume-leaders?exchange=NSE`
- **Status:** ✅ Exists
- **Response JSON (200 OK):**
```json
[
  {
    "ticker": "HDFCBANK",
    "name": "HDFC Bank",
    "price": 1450.00,
    "change": 10.00,
    "change_pct": 0.69,
    "volume": 50000000
  }
]
```

### 19. Get Market Sentiment
- **Endpoint:** `GET /market/sentiment`
- **Status:** 🔨 Build
- **Expected Response JSON:**
```json
{
  "score": 72,
  "label": "Greed",
  "trend": "rising"
}
```

---

## ⚡ Realtime Quotes (`/realtime`)

### 20. Live WebSocket Stream
- **Endpoint:** `WS /realtime/quotes?token=<access_token>`
- **Status:** 🔌 WebSocket
- **Request (Subscribe):**
```json
{ "action": "subscribe", "tickers": ["AAPL", "RELIANCE", "INFY"] }
```
- **Response Stream (Tick):**
```json
{ 
  "ticker": "AAPL", 
  "price": 174.32, 
  "change": 1.45, 
  "change_pct": 0.84, 
  "volume": 52000000, 
  "ts": "2024-03-20T10:00:00Z" 
}
```

---

## ⭐ Watchlists (`/watchlist`)
*All Status: 🔨 Build*

### 21. Get User Watchlists
- **Endpoint:** `GET /watchlist`
- **Response JSON (Expected):**
```json
[
  {
    "id": "wl_123",
    "name": "My Tech Stocks",
    "tickers": [
      {
        "ticker": "AAPL",
        "price": 174.32,
        "change_pct": 0.84
      }
    ]
  }
]
```

### 22. Create Watchlist
- **Endpoint:** `POST /watchlist`
- **Request JSON:**
```json
{ "name": "My Tech Stocks" }
```

### 23. Add Stock to Watchlist
- **Endpoint:** `POST /watchlist/:id/stocks`
- **Request JSON:**
```json
{ "ticker": "AAPL" }
```

### 24. Reorder Watchlist
- **Endpoint:** `PUT /watchlist/:id/reorder`
- **Request JSON:**
```json
{ "order": ["TSLA", "AAPL", "MSFT"] }
```

---

## 🧪 Paper Trading (`/paper`)
*All Status: 🔨 Build*

### 25. Place Order
- **Endpoint:** `POST /paper/orders`
- **Request JSON:**
```json
{
  "ticker": "AAPL",
  "type": "BUY",
  "quantity": 10,
  "order_type": "MARKET"
}
```

### 26. Get Portfolio
- **Endpoint:** `GET /paper/portfolio`
- **Response JSON:**
```json
{
  "cash_balance": 87450.00,
  "portfolio_value": 12550.00,
  "total_value": 100000.00,
  "unrealized_pnl": 550.00,
  "unrealized_pnl_pct": 0.55,
  "starting_balance": 100000.00
}
```

### 27. Get Performance Stats
- **Endpoint:** `GET /paper/performance/stats`
- **Response JSON:**
```json
{
  "total_trades": 42,
  "winning_trades": 28,
  "losing_trades": 14,
  "win_rate": 66.67,
  "best_trade": { "ticker": "TSLA", "pnl": 1240.00 },
  "worst_trade": { "ticker": "COIN", "pnl": -380.00 },
  "total_return_pct": 8.3
}
```

---

## 🔔 Alerts (`/alerts`)
*All Status: 🔨 Build*

### 28. Create Alert
- **Endpoint:** `POST /alerts`
- **Request JSON:**
```json
{
  "ticker": "AAPL",
  "type": "PRICE_TARGET",
  "condition": "ABOVE",
  "value": 200.00,
  "notify_via": ["push", "in_app"]
}
```

### 29. Get Alerts History
- **Endpoint:** `GET /alerts/history?limit=50`
- **Response JSON:**
```json
[
  {
    "id": "alt_123",
    "ticker": "AAPL",
    "type": "PRICE_TARGET",
    "condition": "ABOVE",
    "value": 200.00,
    "triggered_at": "2024-03-20T10:00:00Z"
  }
]
```

---

## 🛠 Utility (`/health`, `/user`)

### 30. Health Check
- **Endpoint:** `GET /health`
- **Status:** ✅ Exists
- **Response JSON:**
```json
{ "status": "up" }
```

### 31. Update User Preferences
- **Endpoint:** `PUT /user/preferences`
- **Status:** 🔨 Build
- **Request JSON:**
```json
{ 
  "theme": "dark", 
  "default_exchange": "NSE", 
  "currency": "INR" 
}
```
