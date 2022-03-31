# 🦕 Golephant

## Installation
```bash
# Clone the repository
git clone https://github.com/Almazatun/golephant.git
# Enter into the directory
cd golephant/
# Install the dependencies
go mod download
```

### Starting the application

```bash
# Build and Up
$ docker-compose up --build -d
```
## Endpoints

### User

| Description | http | path |
|:--:|:--:|:--|
| register | POST | /BASE_URL/user/register |
| login  | PUT | /BASE_URL/user/login |
| update | PATCH | /BASE_URL/user/:userId |
### UML diagram
<img src="./assets/uml-golephant.png">