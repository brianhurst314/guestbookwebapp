# guestbookwebapp

# index.html
The index.html file is just a simple webpage where anyone can enter a name and then it will be saved to the database.
The names from the database will be displayed on the page and can be changed or deleted.

# webserver.go
The webserver.go file is the code that listens to requests from the webpage and performs CRUD(create, read, update, delete)
actions as necessary with the database. 

# script.js, styles.css
These files are used to decorate and add some simple functionality to the webpage. I was having some issues with trying to link
to these files from the html document and so for now I have included their contents at the end of the html file.

# Dockerfile
This is the docker file that I used to create the container image of the webserver so that it can be deployed with docker. This is 
the first docker file I have created and I am sure there is room for improvement, but it does what I intended and so it will
do for now.

# go.mod, go.sum
These describe the dependencies required by the webserver.go file. I am still pretty new to go and this is the first time
I have needed these files. 

# The following are the steps I took to get things up and running

I used docker to create a custom network.
$docker network create my_app_net

I created a custom docker volume.
$docker volume create --name mysqldata

I created a mysql container with access to the custom network and volume.
You must replace the "Your Password" with the password you choose to access your database.
Also if you named your volume and network differently than I have, then use your data.
$docker run -d --name mysql -v mysqldata:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={"Your Password"} --network my_app_net mysql

I used the docker file to create a new image.
The docker file needs to be in the directory of where html and webserver.go files are located.
If the pwd is not where the docker file is located then you need to replace "." with the 
full path to the docker file.
$docker build -t myapp .

Now myapp is an image that I can use to create a second container.
The network needs to be the same as the mysql. The first part of the link is the name
of the mysql container, and the second part is the name that you use in the webserver.go
in order to connect to the db. You can choose different port numbers.
$docker run --name goweb --network my_app_net --link mysql:mydb -p 8080:8081 myapp

With these containers up and running I was able to access the webpage from localhost:8080 and so 
far it is working as intended.

My next step from here is to create a simple kubernetes cluster with a master node and two worker
nodes. I want to put one container on one worker node and the other container on the second
worker node. I intend to use persistent storage for the database data, as well as the index.html as 
well. That way I can easily make changes to the html file with immediate update to the server.
