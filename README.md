Description
- Repo follows clean architecture to ensure maintainability and scalability.

Features
- Employment management
- GoGin for building APIs
- MySQL database
- Gorm for ORM

Prerequisites
- Docker

Installation
1. Clone the repository:
    - git clone https://github.com/LongDuong123/Test.git

2. Create a [.env] file in the root directory and add the following environment variables:
    - MYSQL_ROOT_PASSWORD=rootpassword
    - MYSQL_DATABASE=mydatabase
    - MYSQL_USER=myuser
    - MYSQL_PASSWORD=mypassword
    - MYSQL_PORT=3306
    - MYSQL_HOST=mysql

Running the Application
1. Build and start the application using Docker Compose:
    - docker-compose up 
2. The application will be available at [http://localhost:8080].