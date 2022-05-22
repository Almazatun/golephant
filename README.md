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

### Auth

| Description | http | path |
|:--:|:--:|:--|
| REGISTER_COMPANY | POST | BASE_URL/auth/register/company |
| LOGIN_COMPANY  | PUT | BASE_URL/auth/login/company |
| REGISTER_USER | POST | BASE_URL/auth/register/user |
| LOGIN_USER  | PUT | BASE_URL/auth/login/user |
| ME  | PUT | BASE_URL/auth/me |

### User

| Description | http | path |
|:--:|:--:|:--|
| UPDATE_USER_DATA | PATCH | BASE_URL/user/:userId |
| GET_LINK_RESET_PASSWORD | POST | BASE_URL/user/:userId/resetPassword |
| RESET_PASSWORD | PUT | BASE_URL/user/:userId/resetPassword/:token |

### Company

| Description | http | path |
|:--:|:--:|:--|


### Company position

| Description | http | path |
|:--:|:--:|:--|
| ADD | POST | BASE_URL/company/:companyId/position  |
| UPDATE_RESPONSOBILITIES  | PUT | BASE_URL/company/:companyId/position/:positionId/reponsobilities  |
| UPDATE_REQUIREMENTS  | PUT | BASE_URL/company/:companyId/position/:positionId/requirements  |
| UPDATE  | PUT | BASE_URL/company/:companyId/position/:positionId  |
| UPDATE_STATUS  | PUT | BASE_URL/company/:companyId/position/:positionId/status  |
| DELETE  | DELETE | BASE_URL/company/:companyId/position/:positionId  |

### Company address

| Description | http | path |
|:--:|:--:|:--|
| ADD | POST | BASE_URL/company/:companyId/address  |
| DELETE  | DELETE | BASE_URL/company/:companyId/address/:companyAddressId  |
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

### User education in resume

| Description | http | path |
|:--:|:--:|:--|
| CREATE_AND_UPDATE | PUT | BASE_URL/resume/:resumeId/user/:userId/userEducation |
| DELETE | DELETE | BASE_URL/resume/:resumeId/userEducation/:userEducationId |

### User experiences in resume

| Description | http | path |
|:--:|:--:|:--|
| CREATE_AND_UPDATE | PUT | BASE_URL/resume/:resumeId/user/:userId/userExperiences |
| DELETE | DELETE | BASE_URL/resume/:resumeId/userExperience/:userExperienceId |
### UML diagram
<img src="./assets/uml-golephant.png">