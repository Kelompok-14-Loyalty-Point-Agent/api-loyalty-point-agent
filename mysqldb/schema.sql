CREATE DATABASE IF NOT EXISTS api_loyalty_point_agent_db;
ALTER TABLE stocks MODIFY COLUMN last_top_up DATETIME NULL;
