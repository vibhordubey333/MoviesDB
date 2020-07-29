# MoviesDB
Microservice written in Golang , where user can save new movies as well as rate, comment them . Functionality is similar to IMDB website where guest can see the movies , logged in users can comment , rate , admin can save new movies . For DB MongoDB is used.

# Run.

0.Clone the project.

1.Inside project directory, fire command **"go run ./cmd/movies-d-b-server/main.go --port=7777**".

2.Swagger is present in "api" directory. All functionalities are stored under directories like dbconnection , service,handlers are self explanatory.

3.Testcases are present under testcases directory.

4.Refer cURL_Requests.rtf for testing the API.

5.Future scope for this microservice.

    -Add update service.
    -Contanerization using Docker.
    -Minikube.
    -Introduce Makefile.
    -Improve testcases.
 

