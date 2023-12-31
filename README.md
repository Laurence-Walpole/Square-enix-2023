# SQUARE ENIX - BackEnd Developer Test

The candidate will implement a microservice in a language of their choice exposing an API that will allow distributed processing of data in a database.

The app receives the DB connection configuration via environment variables.
It will expose the following endpoints (either REST or GraphQL), returning a JSON response:

- `start`
  Will start the processing
  
- `stats`
  Will return the (approximate) number of elements already processed
  
The app will be deployed in an unspecified number of instances and there is no service discovery mechanism.

Upon receiving a request on the `start` endpoint the app:
- checks if the process was already started, and return a 429 status for REST or `false` for GraphQL, if that's the case
- if it wasn't started, mark it as started, then return 202 for REST or `true` for GraphQL
- each instance must then process batches of items in the DB, avoiding duplicate processing, until all the items have been processed

Upon receiving a request on the `stats` endpoint the app:
- checks if the process is in progress or finished and return the number of already processed items (with a 200 status for REST)
- if the process is not started, return 412 for REST or `null` for GraphQL

The DB can be any of (candidate's choice):
- MySQL/MariaDB
- PostgreSQL
- MongoDB

The DB can ideally contain millions of items and the structure can be defined by the candidate.
The processing function can be chosen by the candidate, but must be done in the app, not with DB functions.
For example: convert a text field to lowercase, or set to 0 any negative number in a numeric field, or sum two numeric fields and store the result in a third field.

The candidate can choose how to implement the synchronisation among instances of the microservice.

The project should contain the migration to create the DB table or index (for SQL) or a description of the structure and the indexes (for MongoDB)
The resulting codebase should be versioned with git, and the whole directory, including the `.git` folder, should be zipped and sent back. Do not include vendored dependencies in the archive.


Optional bonus features:

- New instances of the microservice could be spawn during the processing. They should join the processing.

- Add a `pause` endpoint that will stop the process and store the last item processed

- Modify `start` to restart from (after) the last item processed if `pause` was invoked previously.

