# CLI-go-sport-Api

<h1>Sports API</h1>

<h3>Description</h3>

<p>
A well designed Microservice following good patterns!
The service fetch a external api (sports API) giving
informations like odds for each website offering betting service.
</p>

<h3>Get your API key</h3>

[API](https://the-odds-api.com/)

<h3>Create Config</h3>

<p>
Create a file inside config folder which is called app.env
</p>

```
dsn=host=localhost port=5432 user=Username dbname=Database password=Password sslmode=disable
url=YourUrlFromSportsAPIwithKey
```

<h3>Endpoints</h3>

1. Fetch endpoints, send a request to the api and save the response in the Database
```
127.0.0.1:8080/fetch
```

2. Get, take an identifier and return data and sites from this id
```
127.0.0.1:8080/sports/{id}
```

3. Get All, does not take any arguments and return all the data and sites
```
127.0.0.1:8080/sports
```

4. Update, not working for the moment but should update odds
```
127.0.0.1:8080/sports/update
```

5. Delete if the times is different from the actual time
```
127.0.0.1:8080/sports/delete
```

<h3>Packages used</h3>

1. [Go-kit](https://github.com/go-kit/kit)
2. [Gorm](https://gorm.io)
3. [Gorilla Mux](https://github.com/gorilla/mux)
