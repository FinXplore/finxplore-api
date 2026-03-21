# FinXplore API Task List

> **Base URL:** `/api/v1`
> **Legend:** ✅ Exists · 🔨 Build · 🔌 WebSocket

---

## 🔐 Authentication & Users

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 1 | `POST` | `/auth/register` | Register new user with email, password, name, mobile, currency | ✅ Exists |
| 2 | `POST` | `/auth/login` | Login with dialing code + mobile + password, returns JWT tokens | ✅ Exists |
| 3 | `POST` | `/auth/refresh` | Exchange refresh token for new access token | ✅ Exists |
| 4 | `GET`  | `/auth/me` | Get current authenticated user profile and role | ✅ Exists |
| 5 | `POST` | `/auth/logout` | Invalidate refresh token server-side, end session | 🔨 Build |
| 6 | `PUT`  | `/auth/me` | Update profile — first/last name, avatar URL, preferred currency | 🔨 Build |
| 7 | `POST` | `/auth/change-password` | Change password (requires current password + new password) | 🔨 Build |

**Notes:**
- All protected routes require `Authorization: Bearer <access_token>` header
- Logout should blacklist the refresh token to prevent reuse
- Profile update should support partial updates (PATCH semantics acceptable)

---

## 📈 Stock Data

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 8  | `GET` | `/stock/search?q=Apple` | Search stocks by name or ticker symbol, returns ticker + exchange | ✅ Exists |
| 9  | `GET` | `/stock/:ticker` | Full stock details — price, P/E ratio, sector, market cap, description | ✅ Exists |
| 10 | `GET` | `/stock/:ticker/chart?period=1d` | OHLCV candlestick data for given timeframe (1d, 1w, 1m, 1y) | ✅ Exists |
| 11 | `GET` | `/stock/:ticker/quote` | Latest snapshot quote — current price, bid, ask, % change, volume | 🔨 Build |
| 12 | `GET` | `/stock/:ticker/indicators?type=RSI&period=14` | Computed technical indicators — RSI, MACD, EMA, Bollinger Bands | 🔨 Build |
| 13 | `GET` | `/stock/:ticker/news?limit=20` | News articles for ticker — headline, source, URL, published timestamp | 🔨 Build |
| 14 | `GET` | `/stock/:ticker/financials?type=income` | Financial statements — income statement, balance sheet, cash flow | 🔨 Build |

**Notes:**
- `/chart` period values should support: `1d`, `5d`, `1mo`, `3mo`, `6mo`, `1y`, `5y`
- `/indicators` type values: `RSI`, `MACD`, `EMA`, `SMA`, `BB` (Bollinger Bands)
- `/financials` type values: `income`, `balance`, `cashflow`
- Consider using Polygon.io, Twelve Data, or Alpha Vantage as data providers
- `/quote` is the REST fallback; WebSocket (see Realtime section) handles live updates

---

## 🌍 Market Overview

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 15 | `GET` | `/market/indices` | Major indices — NIFTY 50, SENSEX, NASDAQ, S&P 500, DOW, FTSE | 🔨 Build |
| 16 | `GET` | `/market/gainers?exchange=NSE&limit=10` | Top gaining stocks for today by exchange | 🔨 Build |
| 17 | `GET` | `/market/losers?exchange=NSE&limit=10` | Top losing stocks for today by exchange | 🔨 Build |
| 18 | `GET` | `/market/volume-leaders?exchange=NSE&limit=10` | Highest volume stocks traded today | 🔨 Build |
| 19 | `GET` | `/market/sentiment` | Market sentiment score — fear/greed index with label and value 0–100 | 🔨 Build |

**Notes:**
- `exchange` query param values: `NSE`, `BSE`, `NASDAQ`, `NYSE`, `ALL`
- Indices response should include: name, value, change, change%, last updated
- Sentiment response: `{ score: 72, label: "Greed", trend: "rising" }`
- Consider caching these responses for 60s to reduce provider API calls
- Gainers/losers/volume-leaders can be powered by a scheduled job that refreshes every minute

---

## ⚡ Realtime Quotes

| # | Type | Endpoint | Description | Status |
|---|------|----------|-------------|--------|
| 20 | `WS`  | `ws://api/v1/realtime/quotes` | Live price stream — subscribe/unsubscribe to tickers, receive tick updates | 🔌 WebSocket |
| 21 | `GET` | `/stock/:ticker/quote` | REST snapshot fallback — latest price, change, volume (see #11 above) | 🔨 Build |

**WebSocket Protocol:**

Subscribe message (client → server):
```json
{ "action": "subscribe", "tickers": ["AAPL", "RELIANCE", "INFY"] }
```

Tick message (server → client):
```json
{ "ticker": "AAPL", "price": 174.32, "change": 1.45, "change_pct": 0.84, "volume": 52000000, "ts": "2024-03-20T10:00:00Z" }
```

Unsubscribe message (client → server):
```json
{ "action": "unsubscribe", "tickers": ["AAPL"] }
```

**Notes:**
- WebSocket connection requires a valid JWT passed as query param: `?token=<access_token>`
- Server should support at least 50 concurrent subscriptions per connection
- Heartbeat ping/pong every 30s to keep connection alive
- On reconnect, client should re-send subscribe message with full ticker list

---

## ⭐ Watchlists

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 22 | `GET`  | `/watchlist` | Get all watchlists for current user with tickers and latest prices | 🔨 Build |
| 23 | `POST` | `/watchlist` | Create a new named watchlist | 🔨 Build |
| 24 | `PUT`  | `/watchlist/:id` | Rename watchlist | 🔨 Build |
| 25 | `DELETE` | `/watchlist/:id` | Delete a watchlist and all its tickers | 🔨 Build |
| 26 | `POST` | `/watchlist/:id/stocks` | Add ticker to watchlist | 🔨 Build |
| 27 | `DELETE` | `/watchlist/:id/stocks/:ticker` | Remove ticker from watchlist | 🔨 Build |
| 28 | `PUT`  | `/watchlist/:id/reorder` | Reorder tickers via drag-drop (accepts sorted array of tickers) | 🔨 Build |

**Request/Response Examples:**

`POST /watchlist` body:
```json
{ "name": "My Tech Stocks" }
```

`POST /watchlist/:id/stocks` body:
```json
{ "ticker": "AAPL" }
```

`PUT /watchlist/:id/reorder` body:
```json
{ "order": ["TSLA", "AAPL", "MSFT", "GOOGL"] }
```

**Notes:**
- `GET /watchlist` response should embed latest quote data per ticker to avoid N+1 fetches on the frontend
- Each user should be capped at a reasonable watchlist limit (e.g. 10 watchlists, 50 stocks each)
- Reorder endpoint should validate that the provided ticker list matches exactly the existing tickers

---

## 🧪 Paper Trading

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 29 | `GET`  | `/paper/portfolio` | Virtual portfolio — cash balance, holdings, total value, unrealized PnL | 🔨 Build |
| 30 | `POST` | `/paper/orders` | Place a simulated buy or sell order at current market price | 🔨 Build |
| 31 | `GET`  | `/paper/orders?status=filled` | Order history — all past paper trades with fill price and timestamp | 🔨 Build |
| 32 | `DELETE` | `/paper/orders/:id` | Cancel a pending paper order | 🔨 Build |
| 33 | `GET`  | `/paper/positions` | Current open positions — ticker, qty, avg cost, current price, PnL | 🔨 Build |
| 34 | `GET`  | `/paper/performance/daily` | Daily PnL history array — used to render equity curve chart | 🔨 Build |
| 35 | `GET`  | `/paper/performance/stats` | Aggregate stats — win rate, best trade, worst trade, total return % | 🔨 Build |
| 36 | `GET`  | `/paper/badges` | Badges and achievements earned in training mode | 🔨 Build |
| 37 | `POST` | `/paper/reset` | Reset paper account to default starting balance, clear all trades | 🔨 Build |

**Request/Response Examples:**

`POST /paper/orders` body:
```json
{
  "ticker": "AAPL",
  "type": "BUY",
  "quantity": 10,
  "order_type": "MARKET"
}
```

`GET /paper/portfolio` response:
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

`GET /paper/performance/stats` response:
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

**Notes:**
- Starting balance should be configurable per user (default: ₹1,00,000 or $10,000 based on currency)
- Order execution uses the latest available price from `/stock/:ticker/quote`
- `order_type` values: `MARKET` only for MVP (add `LIMIT` and `STOP` later)
- Badges examples: "First Trade", "10 Trades", "First Profitable Week", "No Loss Streak x5"
- Reset should soft-delete trade history (keep for analytics) or archive it

---

## 🔔 Alerts & Notifications

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 38 | `GET`  | `/alerts` | Get all active price and % change alerts for current user | 🔨 Build |
| 39 | `POST` | `/alerts` | Create price target or % change alert for a ticker | 🔨 Build |
| 40 | `PUT`  | `/alerts/:id` | Update alert threshold, condition, or toggle active/paused | 🔨 Build |
| 41 | `DELETE` | `/alerts/:id` | Delete an alert permanently | 🔨 Build |
| 42 | `GET`  | `/alerts/history?limit=50` | History of triggered alerts for notification center | 🔨 Build |

**Request/Response Examples:**

`POST /alerts` body:
```json
{
  "ticker": "AAPL",
  "type": "PRICE_TARGET",
  "condition": "ABOVE",
  "value": 200.00,
  "notify_via": ["push", "in_app"]
}
```

Alert types: `PRICE_TARGET`, `PERCENT_CHANGE`
Condition values: `ABOVE`, `BELOW`
Notify via: `push`, `in_app`, `email`

**Notes:**
- Alert evaluation should run on a background job that polls prices every 60s
- Once triggered, alert should auto-pause (not fire repeatedly) unless user re-enables
- `GET /alerts/history` powers the notification center dropdown in the navbar
- Push notifications require a device token registration endpoint (add if building mobile app)

---

## 🛠 Utility

| # | Method | Endpoint | Description | Status |
|---|--------|----------|-------------|--------|
| 43 | `GET` | `/health` | Server health check — returns `{ status: "up" }` | ✅ Exists |
| 44 | `GET` | `/user/preferences` | Get user preferences — theme, default exchange, currency display | 🔨 Build |
| 45 | `PUT` | `/user/preferences` | Update user preferences | 🔨 Build |

**Notes:**
- Preferences payload: `{ theme: "dark" | "light", default_exchange: "NSE", currency: "INR" }`

---

## 📊 Summary

| Category | Total | ✅ Exists | 🔨 To Build | 🔌 WebSocket |
|----------|-------|-----------|-------------|--------------|
| Authentication | 7 | 4 | 3 | — |
| Stock Data | 7 | 3 | 4 | — |
| Market Overview | 5 | 0 | 5 | — |
| Realtime Quotes | 2 | 0 | 1 | 1 |
| Watchlists | 7 | 0 | 7 | — |
| Paper Trading | 9 | 0 | 9 | — |
| Alerts | 5 | 0 | 5 | — |
| Utility | 3 | 1 | 2 | — |
| **Total** | **45** | **8** | **36** | **1** |

---

## 🗺 Recommended Build Order

### Phase 1 — Core foundation
1. Auth gaps (logout, profile update, change password)
2. Market overview endpoints (indices, gainers, losers)
3. Stock quote snapshot + news + financials

### Phase 2 — Engagement features
4. Realtime WebSocket for live price streaming
5. Watchlist CRUD + reorder

### Phase 3 — Training platform
6. Paper trading — portfolio, orders, positions
7. Paper trading — performance stats + badges

### Phase 4 — Retention features
8. Alerts — create, manage, history
9. User preferences
10. Push notifications (device token registration)

---

*Generated for FinXplore · Base URL: `/api/v1` · Auth: Bearer JWT*
