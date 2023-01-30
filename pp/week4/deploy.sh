#!/bin/bash

heroku container:push web -a practica-2023
heroku container:release web -a practica-2023
heroku ps:scale web=1 -a practica-2023
heroku open -a practica-2023