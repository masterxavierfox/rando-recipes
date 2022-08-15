# Welcome to Rando Recipes

An Go webapp that loads a random recipe from a list of recipes.

## Starting the Application

Rando Recipes is built of [buffalo](https://gobuffalo.io/documentation/getting_started/installation/) which ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```console
buffalo dev
```
#### Deploy To heroku
1. Creat Heroku account. [Start Here](https://signup.heroku.com/login).
2. Install heroku-cli then run below commands. [Start here](https://devcenter.heroku.com/articles/heroku-cli):

To login with heroku cli and create a new app run:

```console
heroku container:login && heroku create
```

To build and deploy to heroku using docker run:
```console
heroku container:push web && heroku container:release web

```

[For more Herkou info go here.](https://devcenter.heroku.com/articles/container-registry-and-runtime)

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Rando Recipes Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io) and [Powered by themealdb.com](https://www.themealdb.com/)
