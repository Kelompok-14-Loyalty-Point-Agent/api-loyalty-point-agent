# api-loyalty-point-agent

### Connect to SSH

```
ssh -i "capstoneAPI.pem" ec2-user@ec2-3-0-59-152.ap-southeast-1.compute.amazonaws.com
```

### Connect to MySQL through docker-compose

1. Use docker-compose exec to get into container service
    ```
    docker-compose exec mysql bash
    ```

2. Login to mysql (Password: root123)
    ```
    mysql -u root -p -P3307 api_loyalty_point_agent_db
    ```

3. Run any sql command
    ```
    SHOW DATABASES;
    ```

### API IPv4 Adrress

Use this link to test API (Currently):

https://3.0.59.152.nip.io


### Local Use

Without Docker:

1. Change .env file (Based on your local mysql configure):
    ```
    DB_HOST="localhost" 
    DB_PORT="3306"
    DB_USERNAME="root"
    DB_PASSWORD=""
    DB_NAME="api_loyalty_point_agent_db"
    ```

2. Run this code in terminal:
    ```
    go run main.go
    ```

With Docker (**make user docker desktop is running**):

1. Change .env file:
    ```
    DB_HOST="mysql" 
    DB_PORT="3306"
    DB_USERNAME="root"
    DB_PASSWORD="root123"
    DB_NAME="api_loyalty_point_agent_db"
    ```

2. Run this code in terminal:
    ```
    docker-compose up -d
    ```
    

