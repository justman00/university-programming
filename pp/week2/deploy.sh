#!/bin/bash

heroku container:push web -a pp-week2 
heroku container:release web -a pp-week2
heroku ps:scale web=1 -a pp-week2
heroku open -a pp-week2