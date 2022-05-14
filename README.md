# 🦕 Golephant

Golephant is a backend service for searching jobs and finding companies that are hiring candidates for their open positions. The service also allows people to create their own resumes to represent them-self.

## Installation
```bash
# Clone the repository
git clone https://github.com/Almazatun/golephant.git
# Enter into the directory
cd golephant/
# Install the dependencies
go mod download
```
### Parameters
In order to run the server make the following steps:
```bash
# Step 1 (Create .env file)
$ touch .env
# Step 2 (ENV variables)
* Include some env variables:
    # DB
     * DB_PG
     * DB_DATABASE
     * DB_USER
     * DB_PASSWORD
     * DB_HOST
     * DB_PORT
    # DB_EXTENSIONS
     * POSTGRES_EXTENSIONS
    # JWT
     * JWT_SECRET_KEY
    # COOKIE
     * SET_COOKIE_PATH
    # SMTP(GMAIL)
     * SMTP_MAIL_FROM
     * SMTP_MAIL_PASSWORD
     * SMTP_MAIL_PORT
     * SMTP_MAIL_HOST
```

### Run app with docker-compose
```bash
# Build and Up
$ docker-compose up --build -d
# Stop
$ docker-compose down
# Build and Up with Swagger UI
$ docker-compose --profile swaggerapi up --build
```

## Checking API documents with swagger UI
Browse to http://localhost:3002
You can see all the documented endpoints in UI
## Endpoints

### User

| Description | http | path |
|:--:|:--:|:--|
| REGISTER | POST | BASE_URL/user/register |
| LOGIN  | PUT | BASE_URL/user/login |
| UPDATE_USER_DATA | PATCH | BASE_URL/user/:userId |
| AUTH_ME | POST | BASE_URL/authMe/:userId |
| GET_LINK_RESET_PASSWORD | POST | BASE_URL/user/:userId/resetPassword |
| RESET_PASSWORD | PUT | BASE_URL/user/:userId/resetPassword/:token |

### Resume

| Description | http | path |
|:--:|:--:|:--|
| LIST | GET | BASE_URL/resumes/user/:userId |
| CREATE | POST | BASE_URL/resume/user/:userId |
| DELETE  | PUT | BASE_URL/resume/:resumeId|
| BASIC_INFO | PUT | BASE_URL/resume/:resumeId/user/:userId/basicInfo |
| ABOUT_ME | PUT | BASE_URL/resume/:resumeId/user/:userId/aboutMe |
| CITIZENSHIP | PUT | BASE_URL/resume/:resumeId/user/:userId/citizenship |
| DESIRED_POSITION | PUT | BASE_URL/resume/:resumeId/user/:userId/desiredPosition |
| TAGS | PUT | BASE_URL/resume/:resumeId/user/:userId/tags |

### User Education

| Description | http | path |
|:--:|:--:|:--|
| CREATE_AND_UPDATE | PUT | BASE_URL/resume/:resumeId/user/:userId/userEducation |
| DELETE | DELETE | BASE_URL/resume/:resumeId/userEducation/:userEducationId |

### User Experiences

| Description | http | path |
|:--:|:--:|:--|
| CREATE_AND_UPDATE | PUT | BASE_URL/resume/:resumeId/user/:userId/userExperiences |
| DELETE | DELETE | BASE_URL/resume/:resumeId/userExperience/:userExperienceId |
### UML diagram
<img src="./assets/uml-golephant.png">