package main
/**
@apiGroup Crypto
@api {get} /coin/{crypto} Get
@apiName Get
@apiSuccess {String} name Name of the User.
@apiSampleRequest /coin/bitcoin
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "id": "bitcoin",
    "name": "Bitcoin",
    "symbol": "BTC",
    "price_usd": "8816.88",
    "price_btc": "1.0",
    "price_brl": "28742.0501263",
    "total_supply": "16918475.0",
    "max_supply": "21000000.0",
    "rank": "1",
    "last_updated": "1521023666",
    "market_cap_brl": "486271656511",
    "Calculate": 0
}

@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Not Found"
}
*/
/**
@apiGroup Crypto
@api {get} /coin/{crypto}/value Calculate
@apiName Calculate
@apiSuccess {String} name Name of the User.
@apiSampleRequest /coin/bitcoin/500
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "id": "bitcoin",
    "name": "Bitcoin",
    "symbol": "BTC",
    "price_usd": "8816.88",
    "price_btc": "1.0",
    "price_brl": "28742.0501263",
    "total_supply": "16918475.0",
    "max_supply": "21000000.0",
    "rank": "1",
    "last_updated": "1521023666",
    "market_cap_brl": "486271656511",
    "Calculate": 0.017396114675288323
}

@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Not Found"
}
*/
/**
@apiGroup Transaction
@api {get} /transaction/{id} Get
@apiName Get
@apiSuccess {String} name Name of the User.
@apiSampleRequest /transaction/5aa8f5a32f2eeb7f6b6fefbd
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "data": {
        "_id": "5aa8f5a32f2eeb7f6b6fefbd",
        "amount": "5000",
        "blockchain": {
            "address": "mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk",
            "oapaddress": "",
            "originaladdress": "",
            "private": "e43db0c6690667ea1770affc566b44e9762d7a915cd2d1337876f9c58c29d444",
            "pubkeys": [],
            "public": "020dcfa7c3a40faa705cd02c82352512c3c52823b72abd15c040d58d495ef58c0c",
            "scripttype": "",
            "wif": "cVENYbVLndDebRMbAqXZvGZbTHh7vCWgq8kdven24u2b8HKXX44S"
        },
        "coins": 0.17119981555961808,
        "created_at": "2018-03-14T07:12:51.897-03:00",
        "currency": "bitcoin",
        "status": "waiting",
        "user": {
            "bank": {
                "agency": "1234",
                "bank": "123",
                "number": "5678-9",
                "type": "Corrente"
            },
            "document": "12345678911",
            "email": "teste@teste.com"
        }
    }
}

@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "ID invalid"
}
*/
/**
@apiGroup Transaction
@api {post} /transaction/ Create
@apiName Create

@apiParam {String} amount     Amount send.
@apiHeader Content-Type (application/json).

@apiSampleRequest /transaction
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "data": {
        "_id": "5aa8f5a32f2eeb7f6b6fefbd",
        "amount": "5000",
        "blockchain": {
            "address": "mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk",
            "oapaddress": "",
            "originaladdress": "",
            "private": "e43db0c6690667ea1770affc566b44e9762d7a915cd2d1337876f9c58c29d444",
            "pubkeys": [],
            "public": "020dcfa7c3a40faa705cd02c82352512c3c52823b72abd15c040d58d495ef58c0c",
            "scripttype": "",
            "wif": "cVENYbVLndDebRMbAqXZvGZbTHh7vCWgq8kdven24u2b8HKXX44S"
        },
        "coins": 0.17119981555961808,
        "created_at": "2018-03-14T07:12:51.897-03:00",
        "currency": "bitcoin",
        "status": "waiting",
        "user": {
            "bank": {
                "agency": "1234",
                "bank": "123",
                "number": "5678-9",
                "type": "Corrente"
            },
            "document": "12345678911",
            "email": "teste@teste.com"
        }
    }
}

@apiErrorExample Error-Response:
HTTP/1.1 400 Not Found
{
	"message": "Bad Request"
}
*/
/**
@apiGroup Wallet
@api {get} /wallet/{address} Get
@apiName Get
@apiSampleRequest /wallet/mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "address": "mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt",
    "wallet": {},
    "hd_wallet": {},
    "total_received": 20000000,
    "total_sent": 0,
    "balance": 20000000,
    "unconfirmed_balance": 0,
    "final_balance": 20000000,
    "n_tx": 1,
    "unconfirmed_n_tx": 0,
    "final_n_tx": 1,
    "txs": [
        {
            "block_hash": "0000000096087949986e4fdd50d7d94d6f2ab2a13cc7499020b4ef092db36eba",
            "block_height": 1287820,
            "hash": "3f89f7ee5aff8df7afb560fd676e6d49d8e6ab38f84c8dee95bb6d3f17ea1280",
            "addresses": [
                "mgwMKYweY2zJFodgeYTo72kbsdNCCdFoTv",
                "mtmzAjHxQc9tm6pC7jN3hHL23LEUF1JXTv",
                "mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt"
            ],
            "total": 129990000,
            "fees": 10000,
            "size": 225,
            "preference": "high",
            "relayed_by": "194.135.92.123:18333",
            "received": "2018-03-09T22:29:15.727Z",
            "confirmed": "2018-03-09T22:46:01Z",
            "confirmations": 344,
            "confidence": 1,
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "inputs": [
                {
                    "prev_hash": "054e709658193f95d4363394a9ca143d32208508a566893c2f3a72f2e42e141a",
                    "output_value": 130000000,
                    "addresses": [
                        "mgwMKYweY2zJFodgeYTo72kbsdNCCdFoTv"
                    ],
                    "sequence": 4294967295,
                    "script_type": "pay-to-pubkey-hash",
                    "script": "47304402205af90898fefcc90608c0d710211d0205248bcef4269a8ea76930a1584ad52c6a02205a607f88bdd7f8a5439eb5d9aacdadc286f6573d6d93c1783153a1f96efc405001210343a12cf945c7e4a78b950eeb4825f9f73c72d9f6a10548fd2440805386825482",
                    "age": 1287756
                }
            ],
            "outputs": [
                {
                    "value": 20000000,
                    "addresses": [
                        "mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt"
                    ],
                    "script_type": "pay-to-pubkey-hash",
                    "script": "76a914ad9380dc9658c7a6c0b30f5401feef0252d2e6ec88ac"
                },
                {
                    "spent_by": "baa3c04d7b00497aadbe734d2bd8ccdbd650289653d5a784f1792a1278df44f2",
                    "value": 109990000,
                    "addresses": [
                        "mtmzAjHxQc9tm6pC7jN3hHL23LEUF1JXTv"
                    ],
                    "script_type": "pay-to-pubkey-hash",
                    "script": "76a9149171705010f263097d5f73557c24dce65ca760d888ac"
                }
            ]
        }
    ]
}

@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Not Found"
}
*/
/**
@apiGroup Wallet
@api {get} /wallet List Address
@apiName List Address
@apiSampleRequest /wallet
@apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
    "name": "velrino",
    "addresses": [
        "mwLjtxJF3jZBASVxHLZnUkuG4KGhprWMBt",
        "my8Mm95tuQrHaMFLA3HHENHtHZVZzhJqaC",
        "msCc5FD8aXzaRQKKgh4gpHWZ9sqCGWy3Wu",
        "mjqWQrfSCpX3Gfx2vnAnZxDHtVdgkB3Ygn",
        "my5PECr3oNttbsKMLtQJV18HMmonv1eB4o",
        "mvbH7S29cogAu28WEm8S5KTw3MJAYcK1jS",
        "mtgtFcxtSzjBYHRmGKz9zveX4d4N7M6Ljt",
        "mjfeCHhhFjFJ5pZpfAGbAfZc21bh9Ga4MT",
        "mrUPG8wPieTExA8EiLj86JkcSDquVfQQ4L",
        "mm5bJyxxoUwmxStem8NTxu41LCKmVj5vdF",
        "mrL9VxBbeEBdZKEABy499CJjwYnNHjXwix",
        "mvqAxdej9m5Hgk4TLwX9sdFsWd8Np1iHfy",
        "mj9EPb7HoBeebh3GDLZszk7Te8VHyWJHdk",
        "mrY2BX5GgZKoSR8YSwNmSU2rDgZ4tnM5ZV",
        "mrjE2DZHrWaUjHycBNLLaUYZwrQfbuzMHh",
        "mssTjS5UD5Cn1LprXRZMaWmydpEM5vwR5c",
        "mrdKWaYA51qygyBHVKEhqMJcezTW4XYeT5"
    ]
}

@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Not Found"
}
*/

