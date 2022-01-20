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

Later... 

<h3>To do</h3>

- Finish delete function by checking time match started and delete when finished

<h3>Packages used</h3>

1. [Go-kit](https://github.com/go-kit/kit)
2. [Gorm](https://gorm.io)
3. [Gorilla Mux](https://github.com/gorilla/mux)
