# .env
# ────────────────────────────────────────────────
# Option A: Recommended - simplest & cleanest
# Just use this one line and comment out the rest
DB_URL=postgresql://postgres:password@localhost:5432/mydb?sslmode=disable&TimeZone=Asia/Dhaka

# ────────────────────────────────────────────────
# Option B: individual variables (used if DB_URL is missing)
# DB_HOST=localhost
# DB_USER=postgres
# DB_PASSWORD=password
# DB_NAME=mydb
# DB_PORT=5432

# ────────────────────────────────────────────────
# Common optional settings you might want later

# APP_PORT=8080
# JWT_SECRET=super-long-random-secret-ChangeMe1234567890
# ENVIRONMENT=development
