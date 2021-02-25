# guestbookwebapp

This is a project that I want to use to demonstrate some things that I have been learning

My intention is to use Golang to create a webserver, then design a simple guestbook webpage
that gives the option to add names to a list. The current list will be displayed. I want to 
use a mysql database to store the names. The Golang webserver will communicate with the mysql
database to perform all CRUD(Create, Read, Update, Delete) operations. The webpage will need
options that allow deleting and updating names in the list. I also wanted to try to be able
to sort the list of names in ascending or descending order.

Once all of this is working as intended, I will then create a new docker image using the
Golang webserver I have created. Then I will spin up a kubernetes cluster with a pod for
the webserver and a pod for the database. I want to utilize some type of persistent data
storage for the database and for the html webpage. Then I want to expose the webpage so that 
it can be accessed from any device.

I feel like I have the knowledge, now I just need to apply it