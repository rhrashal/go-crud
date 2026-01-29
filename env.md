# .env
# ────────────────────────────────────────────────
# Option A: Recommended - simplest & cleanest
# Just use this one line and comment out the rest
DB_URL=postgresql://postgres:abcd@localhost:5432/myappdb?sslmode=disable&TimeZone=Asia/Dhaka

# ────────────────────────────────────────────────
# Option B: individual variables (used if DB_URL is missing)
# DB_HOST=localhost
# DB_USER=postgres
# DB_PASSWORD=abcd
# DB_NAME=myappdb
# DB_PORT=5432

# ────────────────────────────────────────────────
# Common optional settings you might want later

# APP_PORT=8080
# JWT_SECRET=super-long-random-secret-ChangeMe1234567890
# ENVIRONMENT=development
