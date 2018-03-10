define({ "api": [
  {
    "type": "get",
    "url": "/api/calculate/{crypto}",
    "title": "List",
    "name": "Calculate",
    "group": "Info",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the User.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n[\n\t{\n\t\t\"id\": 1,\n\t}\n]",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/api/calculate"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "CoinNotFound",
            "description": "<p>The Users was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Coin Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Info"
  }
] });
