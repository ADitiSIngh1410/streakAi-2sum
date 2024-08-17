
**API Endpoint**
Endpoint: /find-pairs
Method: POST

**Request Structure**
{
  "numbers": [1, 2, 3, 4, 5],
  "target": 6
}


**Response Structure**
{
  "solutions": [
    [0, 4],
    [1, 3]
  ]
}


**How to Run**
Clone the repository.
Install the necessary dependencies.
Run the server.
Send a POST request to /find-pairs with the required JSON structure to receive the pairs of indices that match the target sum.
