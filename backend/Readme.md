# Hao

Hao is a web application using MongoDB as its database, containerized with Docker.

## Routes

| Method | Endpoint            | Description              |
|--------|---------------------|--------------------------|
| GET    | `/hotSearches`      | Get all hot searches     |
| GET    | `/categories`       | Get all categories       |
| GET    | `/items`            | Get all items            |
| GET    | `/itemCategories`   | Get all item categories  |
| GET    | `/siteItem`         | Get all site items       |
| GET    | `/commonSiteItem`   | Get all common site items|
| GET    | `/topNews`          | Get all top news         |
| GET    | `/govSites`         | Get all government sites |
| GET    | `/hotList`          | Get all hot lists        |
| GET    | `/hotTab`           | Get all hot tabs         |
| GET    | `/topListItems`     | Get all top list items   |
| GET    | `/topList`          | Get all top lists        |

## Technologies

- **Backend**: Go (Gin Framework)
- **Database**: MongoDB
- **Containerization**: Docker

## Setup

1. Clone the repository:
   ```bash
   git clone http://206.238.118.2:10880/game/hao.git
   cd hao

2. go to cd backend/project
   then run  docker-compose up --build

3. Configure the environment:
    Create a .env file and specify your MongoDB connection settings.

4. Adding Dummy Data for Testing
Before each route call,make sure the function TestAddMany(t *testing.T) is executed to add each dummy data to the database. This function is located in backend/api/db.

5. Run the service:
    go run .