# Copy this file to .env and change the secrets below.
# .env will be ignored by .gitignore. Please don't check secrets into the repository!

# Secrets (fill in locally)
# We suggest generating passwords using the command `openssl rand -base64 32`
POSTGRES_DB=*************
POSTGRES_USER=********
POSTGRES_PASSWORD=*********************

POSTGRES_DB_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}

# POSTGRES_USER2=********
# POSTGRES_PASSWORD2=*********************

# Astro (presentation layer) and API services
ASTRO_PORT=4321
API_PORT=3333
