package server

import (
	"app/src/bundles/quasar-fire/controller"
	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/service"
	router "app/src/http"
)

var (
	topSecretSrv         service.TopSecretService        = service.NewTopSecretService()
	fireQuasarController controller.FireQuasarController = controller.NewFireQuasarController(topSecretSrv)
)

// topsecret response payload
// swagger:response topSecretResp
type swaggTopSecretRes struct {
	// in:body
	Body models.TopSecretResponse
}

func GetTopSecretRoute(httpRouter router.Router) {
	// swagger:operation POST /topsecret Fire_Quasar topsecret
	// ---
	// summary: topsecret.
	// description: Gets the location of the ship and the message it emits.
	// parameters:
	// - name: topsecret
	//   in: body
	//   required: true
	//   schema:
	//     type: object
	//     required:
	//       - satellites
	//     properties:
	//       satellites:
	//         type: array
	//         items:
	//           type: object
	//           properties:
	//             name:
	//               type: string
	//               summary: Satellite name
	//               example: "Kenoby"
	//             distance:
	//               type: number
	//               format: float
	//               summary: Distance from transmitter to satellite
	//               example: 100.0
	//             message:
	//               type: array
	//               items:
	//                 type: string
	//               summary: Message to issue
	//               example: ["this", "is", "a", "message"]
	// responses:
	//   "200":
	//     "$ref": "#/responses/topSecretResp"
	//   "400":
	//     summary: Message error
	//     "$ref": "#/responses/badRequest"
	//   "404":
	//     summary: Message error
	//     "$ref": "#/responses/notFoundReq"
	httpRouter.POST("/topsecret", fireQuasarController.TopSecret)
}

func GetTopSecretSplitRoute(httpRouter router.Router) {
	// swagger:operation GET /topsecretsplit Fire_Quasar topsecretsplit
	// ---
	// summary: topsecretsplit.
	// description: Gets the location of the ship and the message it emits.
	// parameters:
	// - name: notificationTurn
	//   in: body
	//   required: true
	//   schema:
	//     type: object
	//     required:
	//       - distance
	//       - message
	//     properties:
	//       distance:
	//         type: number
	//         format: float
	//         summary: Distance from transmitter to satellite
	//         example: 100.0
	//       message:
	//         type: array
	//         items:
	//           type: string
	//         summary: Message to issue
	//         example: ["this", "is", "a", "message"]
	// responses:
	//   "200":
	//     "$ref": "#/responses/topSecretResp"
	//   "400":
	//     summary: Message error
	//     "$ref": "#/responses/badRequest"
	//   "404":
	//     summary: Message error
	//     "$ref": "#/responses/notFoundReq"
	httpRouter.GET("/topsecret_split/{satellite_name}", fireQuasarController.TopSecretSplit)
}
