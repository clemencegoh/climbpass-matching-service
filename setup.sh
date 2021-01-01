# Sets up DB for use
sudo service postgresql start

# create database if not already
sudo -u postgres psql postgres -c "CREATE DATABASE climbpass WITH ENCODING 'UTF8'"