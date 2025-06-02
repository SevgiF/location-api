package benchmark

import (
	"bytes"
	"encoding/json"
	"github.com/SevgiF/location-api/internal/adapters/test/helpers/location_helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/http/httptest"
	"testing"
)

func BenchmarkCreate(b *testing.B) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.App().Post("/v1/location/add", locationHandler.AddLocation)

	assertionRequest := location_helpers.GetRequestsResponsesForCreateLocation()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		req := assertionRequest[rand.Intn(len(assertionRequest))]

		request := httptest.NewRequest("POST", "/v1/location/add", nil)
		if req.Request != nil {
			jsonBody, err := json.Marshal(req.Request)
			if err != nil {
				panic(err)
			}

			request = httptest.NewRequest("POST", "/v1/location/add", bytes.NewReader(jsonBody))
			request.Header.Set("Content-Type", "application/json")
		}
		_, _ = c.App().Test(request, -1)
	}
	b.StopTimer()
	defer Clean()
}
