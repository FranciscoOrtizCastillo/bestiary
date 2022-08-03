# A GraphQL API for Medieval Beasts! Written in Golang with graphql-go

from 
https://www.apollographql.com/blog/graphql/golang/building-a-graphql-server-in-go-with-go-graphql/

```
go mod tidy 
go run .
```

```
curl -X POST -H "Content-Type: application/json" --data '{ "query": "{ beastList {id name } }" }' http://localhost:8081/graphql
````

### Queries

```
query ExampleQuery {
  beastList {
    id
    description
  }
}
```
>
>Response :
>
>{
  "data": {
    "beastList": [
      {
        "id": "1",
        "description": "A beast"
      },
      {
        "id": "2",
        "description": "Another beast"
      }
    ]
  }
}
{
  "data": {
    "beastList": [
      {
        "description": "A serpent with two heads, one at either end",
        "id": 1
      },
      {
        "description": "A fierce beast with a single long horn",
        "id": 2
      }, .....

```  
query Beast {
  beast(name: "Amphisbaena") {
    description
    id
    name
  }
}
```
>Response :
>
>{
  "data": {
    "beast": {
      "description": "A serpent with two heads, one at either end",
      "id": 1,
      "name": "Amphisbaena"
    }
  }
}

### Mutations

```
mutation RootMutation {
  addBeast(
    name: "Hydra",
    description: "Hydra") 
  {
    description
    id
  }
}
````
>Response :
>
>{
  "data": {
    "addBeast": {
      "description": "Hydra",
      "id": 6
    }
  }
}
