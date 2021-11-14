#!/bin/bash

heroku container:push web -a pp-week4
heroku container:release web -a pp-week4
heroku ps:scale web=1 -a pp-week4
heroku open -a pp-week4