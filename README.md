 Dynamic Schema Generation

 We start by creating a Tenant and the Schema will be created automatically. The name of the Schema will
 be the same as the Tenant name. A Tenant only has one Schema to interact with throught the API.

 Specs:
  * Open Api 3 specs were use to document the API
  * Swagger UI generated with the Open Api 3 specs 

 Comparisons:
  * Schemas are like Databases
  * Collections are like Tables
  * Documents are like Rows

 Architecture and design patterns:
  * Domain Drvien design (DDD)
  * Hexagonal Architecture
  * 4 Layered Architecture
  * Event Driven Architecture (EDA)

 Steps:
   * Fill necessary env variables
   * Run docker-compose file to start necessary services
   * Run go run . in root folder to start application