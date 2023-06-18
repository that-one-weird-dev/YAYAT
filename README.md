# YAYAT
Yet another: Yet another twitter clone
You're probably wondering where the "C" went, well... talking about technologies

# Techonologies
The backend is written in go using [fiber](https://gofiber.io/) as the web framework and [gorm](https://gorm.io/) for the ORM.
As for the frontend is just svelte with typescript, tailwindcss and svelte.

# Running in dev
### Backend
You can run it with air or the default go compiler
- ***Air:*** `air`
- ***Go:*** `go run main.go`

Also you need to have a postgres server running

### Frontend
Just `pnpm run dev` or `npm run dev` or *proceeds to list all the 1.487.437 node package managers in existance*

### The easy way
Or you could just `docker compose up -d`

# Deployment
Also just `docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d`

