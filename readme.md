To get the project running first, install the dependencies 
	1) go : https://golang.org/doc/install
	2) MongoDB community edition
	
Next Initialize the database
	1) Open a mongo shell 
	2) use the cd() command to switch the working direcotry to the "mongo_shell" folder. 
		ex: cd(path\\to\\your\\local\\branch\\mongo_shell")
	3) type "use WebServiceDb" to create and switch over to the WebServiceDb within the shell context 	
	4) use the load() command to load the "InitializeDatabase.js" file, and create the MongoDB collections
		ex: load("InitializeDatabase.js")

1) CD to DataModels, and "go build"
2) CD to Singletons, and "go build"
3) CD to repositories, and "go build" 
4) CD to Controllers, and "go build" 
5) CD to WebService and "go build" 


// *** Notes on go archeticture ***
// I may have overbuilt, and/or created more packages than in conventional by go idiomatic standards. 
// I am not overlly committed to the folder structure and package import pattern I've layed out here. 
//	- ie, if it would be more conventional to condense all struct definitions into a single file, I would be fine with that. 

//*** Notes on Mongo Schema ***
// I've applied priniciles of normalized relational database design to this mongo schema.  The decision to use this relational approach, over the embedded approach, was made primarily to avoid having the semester start/end dates defined in multiple places.  This drastically improves perfomance when updating semester start/end dates, as the dates do not need to be updated in every user document. (It also eliminates the possiblity of having different users with different start/end dates for the same meal plan) It does however reduce the read performance on the user data, as getting the user's meal plan info now consists of one database call to get the user's meal plan, and another database call to get the semester start/end dates associated with that meal plan.  Therefore, despite having decided to go with a relationally designed schema, I fully acknwoledge that the embedded approach has some performance benifits, and should be considered. 