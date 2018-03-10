package main
/**
* @api {get} /api/calculate/{crypto} List
* @apiName Calculate
* @apiGroup Info
*
* @apiSuccess {String} name Name of the User.
* @apiSampleRequest /api/calculate
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
[
	{
		"id": 1,
	}
]
* @apiError CoinNotFound The Users was not found.
*
* @apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Coin Not Found"
}
*/
