# Go Blog

[Live Demo](https://lewiswilliams.info/)

---

# Deployment
copy .env.example to .env and edit the variables accordingly

## Deploying for local development

#### Deploy MySQL database & phpmyadmin
```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

#### Running the application

1. Run the application with `make dev`
2. Migrate the database with `make migrate-up`
3. Seed the database with `make seed`

## Deploying for production

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```