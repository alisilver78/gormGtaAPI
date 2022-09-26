## Golang API With Gorm and sqlite database
### Handlers
### Get- GetData handler function
this method is against /gta endpoint.
GetData finds all records in database and returns them as response.
#### Post-InsertData handler function
this method is against /gta endpoint.
InsertData creates on or more records.
data struct must be between two pair of square brackets.
###### Example of inserting one data:
[
	{
		"year": 2013,
		"name": "Grand Theft Auto V",
		"city": "Los Santos"
	}
]
###### Example of inserting one data:
[
	{
		"year": 1999,
		"name": "Grand Theft Auto 2",
		"city": "Anywhere, USA"
	},
	{
		"year": 2001,
		"name": "Grand Theft Auto III",
		"city": "Liberty City"
	},
	{
		"year": 2002,
		"name": "Grand Theft Auto: Vice City",
		"city": "Vice City"
	},
	{
		"year": 2004,
		"name": "Grand Theft Auto: San Andreas",
		"city": "state of San Andreas"
	},
	{
		"year": 2008,
		"name": "Grand Theft Auto IV",
		"city": "Liberty City"
	},
	{
		"year": 2013,
		"name": "Grand Theft Auto V",
		"city": "Los Santos"
	}
]
