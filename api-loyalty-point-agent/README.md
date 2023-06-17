# api-loyalty-point-agent

### Key Pair

https://drive.google.com/file/d/1242WI-erRCbZGjgtHbMGi5yP8kSGbGqB/view?usp=share_link

### Connect to SSH

```
ssh -i "capstoneAPI.pem" ec2-user@ec2-13-229-84-45.ap-southeast-1.compute.amazonaws.com
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

Use this link to test API:

http://13.229.84.45


