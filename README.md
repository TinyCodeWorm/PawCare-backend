# PawCare-backend
PawCare-backend is backend of PawCare project.

1.after git clone the project, please add a file: constants.go on your project folder, add the content as below:


package main


const (

	ES_URL      = "http://34.135.235.101:9200"
	
	ES_USERNAME = "elastic"
	
	ES_PASSWORD = "lyhK0FMIzzffLiAt073J"
)


remember to change the user name and password!


2.when you use "go run ." to start the server, the indexs will be created automaticly as below:

	"pawcare_user"
	"pet"
	"food"
	"reaction"
	"pet_reaction"

If the index exists, we do nothing, otherwise, we will create the index.
For the index "reaction", we already put data in it, you can use it directly.

