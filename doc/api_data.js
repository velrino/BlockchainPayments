define({ "api": [
  {
    "group": "Crypto",
    "type": "get",
    "url": "/coin/{crypto}/value",
    "title": "Calculate",
    "name": "Calculate",
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
          "content": "HTTP/1.1 200 OK\n{\n    \"id\": \"bitcoin\",\n    \"name\": \"Bitcoin\",\n    \"symbol\": \"BTC\",\n    \"price_usd\": \"8816.88\",\n    \"price_btc\": \"1.0\",\n    \"price_brl\": \"28742.0501263\",\n    \"total_supply\": \"16918475.0\",\n    \"max_supply\": \"21000000.0\",\n    \"rank\": \"1\",\n    \"last_updated\": \"1521023666\",\n    \"market_cap_brl\": \"486271656511\",\n    \"Calculate\": 0.017396114675288323\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/coin/bitcoin/500"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Crypto"
  },
  {
    "group": "Crypto",
    "type": "get",
    "url": "/coin/{crypto}",
    "title": "Get",
    "name": "Get",
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
          "content": "HTTP/1.1 200 OK\n{\n    \"id\": \"bitcoin\",\n    \"name\": \"Bitcoin\",\n    \"symbol\": \"BTC\",\n    \"price_usd\": \"8816.88\",\n    \"price_btc\": \"1.0\",\n    \"price_brl\": \"28742.0501263\",\n    \"total_supply\": \"16918475.0\",\n    \"max_supply\": \"21000000.0\",\n    \"rank\": \"1\",\n    \"last_updated\": \"1521023666\",\n    \"market_cap_brl\": \"486271656511\",\n    \"Calculate\": 0\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/coin/bitcoin"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Crypto"
  },
  {
    "group": "Transaction",
    "type": "post",
    "url": "/transaction/",
    "title": "Create",
    "name": "Create",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "amount",
            "description": "<p>Amount send.</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>(application/json).</p>"
          }
        ]
      }
    },
    "sampleRequest": [
      {
        "url": "/transaction"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"data\": {\n        \"_id\": \"5aa8f5a32f2eeb7f6b6fefbd\",\n        \"amount\": \"5000\",\n        \"blockchain\": {\n            \"address\": \"mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk\",\n            \"oapaddress\": \"\",\n            \"originaladdress\": \"\",\n            \"private\": \"e43db0c6690667ea1770affc566b44e9762d7a915cd2d1337876f9c58c29d444\",\n            \"pubkeys\": [],\n            \"public\": \"020dcfa7c3a40faa705cd02c82352512c3c52823b72abd15c040d58d495ef58c0c\",\n            \"scripttype\": \"\",\n            \"wif\": \"cVENYbVLndDebRMbAqXZvGZbTHh7vCWgq8kdven24u2b8HKXX44S\"\n        },\n        \"coins\": 0.17119981555961808,\n        \"created_at\": \"2018-03-14T07:12:51.897-03:00\",\n        \"currency\": \"bitcoin\",\n        \"status\": \"waiting\",\n        \"user\": {\n            \"bank\": {\n                \"agency\": \"1234\",\n                \"bank\": \"123\",\n                \"number\": \"5678-9\",\n                \"type\": \"Corrente\"\n            },\n            \"document\": \"12345678911\",\n            \"email\": \"teste@teste.com\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 400 Not Found\n{\n\t\"message\": \"Bad Request\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Transaction"
  },
  {
    "group": "Transaction",
    "type": "get",
    "url": "/transaction/{id}",
    "title": "Get",
    "name": "Get",
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
          "content": "HTTP/1.1 200 OK\n{\n    \"data\": {\n        \"_id\": \"5aa8f5a32f2eeb7f6b6fefbd\",\n        \"amount\": \"5000\",\n        \"blockchain\": {\n            \"address\": \"mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk\",\n            \"oapaddress\": \"\",\n            \"originaladdress\": \"\",\n            \"private\": \"e43db0c6690667ea1770affc566b44e9762d7a915cd2d1337876f9c58c29d444\",\n            \"pubkeys\": [],\n            \"public\": \"020dcfa7c3a40faa705cd02c82352512c3c52823b72abd15c040d58d495ef58c0c\",\n            \"scripttype\": \"\",\n            \"wif\": \"cVENYbVLndDebRMbAqXZvGZbTHh7vCWgq8kdven24u2b8HKXX44S\"\n        },\n        \"coins\": 0.17119981555961808,\n        \"created_at\": \"2018-03-14T07:12:51.897-03:00\",\n        \"currency\": \"bitcoin\",\n        \"status\": \"waiting\",\n        \"user\": {\n            \"bank\": {\n                \"agency\": \"1234\",\n                \"bank\": \"123\",\n                \"number\": \"5678-9\",\n                \"type\": \"Corrente\"\n            },\n            \"document\": \"12345678911\",\n            \"email\": \"teste@teste.com\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/transaction/5aa8f5a32f2eeb7f6b6fefbd"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"ID invalid\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Transaction"
  },
  {
    "group": "Wallet",
    "type": "get",
    "url": "/wallet/{address}",
    "title": "Get",
    "name": "Get",
    "sampleRequest": [
      {
        "url": "/wallet/mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"address\": \"mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt\",\n    \"wallet\": {},\n    \"hd_wallet\": {},\n    \"total_received\": 20000000,\n    \"total_sent\": 0,\n    \"balance\": 20000000,\n    \"unconfirmed_balance\": 0,\n    \"final_balance\": 20000000,\n    \"n_tx\": 1,\n    \"unconfirmed_n_tx\": 0,\n    \"final_n_tx\": 1,\n    \"txs\": [\n        {\n            \"block_hash\": \"0000000096087949986e4fdd50d7d94d6f2ab2a13cc7499020b4ef092db36eba\",\n            \"block_height\": 1287820,\n            \"hash\": \"3f89f7ee5aff8df7afb560fd676e6d49d8e6ab38f84c8dee95bb6d3f17ea1280\",\n            \"addresses\": [\n                \"mgwMKYweY2zJFodgeYTo72kbsdNCCdFoTv\",\n                \"mtmzAjHxQc9tm6pC7jN3hHL23LEUF1JXTv\",\n                \"mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt\"\n            ],\n            \"total\": 129990000,\n            \"fees\": 10000,\n            \"size\": 225,\n            \"preference\": \"high\",\n            \"relayed_by\": \"194.135.92.123:18333\",\n            \"received\": \"2018-03-09T22:29:15.727Z\",\n            \"confirmed\": \"2018-03-09T22:46:01Z\",\n            \"confirmations\": 344,\n            \"confidence\": 1,\n            \"ver\": 1,\n            \"vin_sz\": 1,\n            \"vout_sz\": 2,\n            \"inputs\": [\n                {\n                    \"prev_hash\": \"054e709658193f95d4363394a9ca143d32208508a566893c2f3a72f2e42e141a\",\n                    \"output_value\": 130000000,\n                    \"addresses\": [\n                        \"mgwMKYweY2zJFodgeYTo72kbsdNCCdFoTv\"\n                    ],\n                    \"sequence\": 4294967295,\n                    \"script_type\": \"pay-to-pubkey-hash\",\n                    \"script\": \"47304402205af90898fefcc90608c0d710211d0205248bcef4269a8ea76930a1584ad52c6a02205a607f88bdd7f8a5439eb5d9aacdadc286f6573d6d93c1783153a1f96efc405001210343a12cf945c7e4a78b950eeb4825f9f73c72d9f6a10548fd2440805386825482\",\n                    \"age\": 1287756\n                }\n            ],\n            \"outputs\": [\n                {\n                    \"value\": 20000000,\n                    \"addresses\": [\n                        \"mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt\"\n                    ],\n                    \"script_type\": \"pay-to-pubkey-hash\",\n                    \"script\": \"76a914ad9380dc9658c7a6c0b30f5401feef0252d2e6ec88ac\"\n                },\n                {\n                    \"spent_by\": \"baa3c04d7b00497aadbe734d2bd8ccdbd650289653d5a784f1792a1278df44f2\",\n                    \"value\": 109990000,\n                    \"addresses\": [\n                        \"mtmzAjHxQc9tm6pC7jN3hHL23LEUF1JXTv\"\n                    ],\n                    \"script_type\": \"pay-to-pubkey-hash\",\n                    \"script\": \"76a9149171705010f263097d5f73557c24dce65ca760d888ac\"\n                }\n            ]\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Wallet"
  },
  {
    "group": "Wallet",
    "type": "get",
    "url": "/wallet",
    "title": "List Address",
    "name": "List_Address",
    "sampleRequest": [
      {
        "url": "/wallet"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"name\": \"velrino\",\n    \"addresses\": [\n        \"mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt\",\n        \"my8Mm95tuQrHaMFLA3HHENHtHZVZzhJqaC\",\n        \"msCc5FD8aXzaRQKKgh4gpHWZ9sqCGWy3Wu\",\n        \"mjqWQrfSCpX3Gfx2vnAnZxDHtVdgkB3Ygn\",\n        \"my5PECr3oNttbsKMLtQJV18HMmonv1eB4o\",\n        \"mvbH7S29cogAu28WEm8S5KTw3MJAYcK1jS\",\n        \"mtgtFcxtSzjBYHRmGKz9zveX4d4N7M6Ljt\",\n        \"mjfeCHhhFjFJ5pZpfAGbAfZc21bh9Ga4MT\",\n        \"mrUPG8wPieTExA8EiLj86JkcSDquVfQQ4L\",\n        \"mm5bJyxxoUwmxStem8NTxu41LCKmVj5vdF\",\n        \"mrL9VxBbeEBdZKEABy499CJjwYnNHjXwix\",\n        \"mvqAxdej9m5Hgk4TLwX9sdFsWd8Np1iHfy\",\n        \"mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk\",\n        \"mrY2BX5GgZKoSR8YSwNmSU2rDgZ4tnM5ZV\",\n        \"mrjE2DZHrWaUjHycBNLLaUYZwrQfbuzMHh\",\n        \"mssTjS5UD5Cn1LprXRZMaWmydpEM5vwR5c\",\n        \"mrdKWaYA51qygyBHVKEhqMJcezTW4XYeT5\"\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Wallet"
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
