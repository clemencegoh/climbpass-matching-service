# Sets up DB for use
sudo service postgresql start

# create database if not already
sudo -u postgres psql postgres -c "CREATE DATABASE climbpass WITH ENCODING 'UTF8'"

# To delete data from table:
# TRUNCATE TABLE gym_models;
# ALTER SEQUENCE gym_models_id_seq RESTART WITH 1;