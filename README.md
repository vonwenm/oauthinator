# oauthinator

This is a simple golang based oauth2 server which runs on google appengine.

It uses the new appengine api, and tooling which makes it easier to develop along side other golang projects within the one `$GOPATH`.

# endpoints

* POST `/user` create a new user account, this requires an oauth2 application.

# development

Pretty much as follows:

```
cd app
gcloud --verbosity debug preview app run app.yaml
```

```
curl -H "Content-Type: application/json" -X POST -d '{"login": "wolfeidau", "email": "mark@wolfe.id.au", "name": "Mark Wolfe"}' http://localhost:8080/users
```

# disclaimer

I am currently road testing the new app engine apis so this is very much an experiment at the moment.

# license

This code is Copyright (c) 2014 Mark Wolfe and licensed under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included `LICENSE` file for more details.
