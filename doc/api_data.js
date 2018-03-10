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
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_root_go_src_github_com_velrino_BlockchainPayments_doc_main_js",
    "groupTitle": "_root_go_src_github_com_velrino_BlockchainPayments_doc_main_js",
    "name": ""
  }
] });
