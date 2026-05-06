package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"runtime"

	_ "fiber-crud-demo/docs"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/initializer"
	"fiber-crud-demo/middleware"
	"fiber-crud-demo/modules/health"
	"fiber-crud-demo/modules/hello_world"

	"fiber-crud-demo/modules/hospital/m_admin"
	"fiber-crud-demo/modules/hospital/m_bank"
	"fiber-crud-demo/modules/hospital/m_biodata"
	"fiber-crud-demo/modules/hospital/m_biodata_address"
	"fiber-crud-demo/modules/hospital/m_biodata_attachment"
	"fiber-crud-demo/modules/hospital/m_blood_group"
	"fiber-crud-demo/modules/hospital/m_courier"
	"fiber-crud-demo/modules/hospital/m_courier_type"
	"fiber-crud-demo/modules/hospital/m_customer"
	"fiber-crud-demo/modules/hospital/m_customer_member"
	"fiber-crud-demo/modules/hospital/m_customer_relation"
	"fiber-crud-demo/modules/hospital/m_doctor"
	"fiber-crud-demo/modules/hospital/m_doctor_education"
	"fiber-crud-demo/modules/hospital/m_education_level"
	"fiber-crud-demo/modules/hospital/m_location"
	"fiber-crud-demo/modules/hospital/m_location_level"
	"fiber-crud-demo/modules/hospital/m_medical_facility"
	"fiber-crud-demo/modules/hospital/m_medical_facility_category"
	"fiber-crud-demo/modules/hospital/m_medical_facility_schedule"
	"fiber-crud-demo/modules/hospital/m_medical_item"
	"fiber-crud-demo/modules/hospital/m_medical_item_category"
	"fiber-crud-demo/modules/hospital/m_medical_item_segmentation"
	"fiber-crud-demo/modules/hospital/m_menu"
	"fiber-crud-demo/modules/hospital/m_menu_role"
	"fiber-crud-demo/modules/hospital/m_payment_method"
	"fiber-crud-demo/modules/hospital/m_role"
	"fiber-crud-demo/modules/hospital/m_specialization"
	"fiber-crud-demo/modules/hospital/m_user"
	"fiber-crud-demo/modules/hospital/m_wallet_default_nominal"
	"fiber-crud-demo/modules/hospital/t_appointment"
	"fiber-crud-demo/modules/hospital/t_appointment_cancellation"
	"fiber-crud-demo/modules/hospital/t_appointment_done"
	"fiber-crud-demo/modules/hospital/t_appointment_reschedule_history"
	"fiber-crud-demo/modules/hospital/t_courier_discount"
	"fiber-crud-demo/modules/hospital/t_current_doctor_specialization"
	"fiber-crud-demo/modules/hospital/t_customer_chat"
	"fiber-crud-demo/modules/hospital/t_customer_chat_history"
	"fiber-crud-demo/modules/hospital/t_customer_custom_nominal"
	"fiber-crud-demo/modules/hospital/t_customer_registered_card"
	"fiber-crud-demo/modules/hospital/t_customer_va"
	"fiber-crud-demo/modules/hospital/t_customer_va_history"
	"fiber-crud-demo/modules/hospital/t_customer_wallet"
	"fiber-crud-demo/modules/hospital/t_customer_wallet_top_up"
	"fiber-crud-demo/modules/hospital/t_customer_wallet_withdraw"
	"fiber-crud-demo/modules/hospital/t_doctor_office"
	"fiber-crud-demo/modules/hospital/t_doctor_office_schedule"
	"fiber-crud-demo/modules/hospital/t_doctor_office_treatment"
	"fiber-crud-demo/modules/hospital/t_doctor_office_treatment_price"
	"fiber-crud-demo/modules/hospital/t_doctor_treatment"
	"fiber-crud-demo/modules/hospital/t_medical_item_purchase"
	"fiber-crud-demo/modules/hospital/t_medical_item_purchase_detail"
	"fiber-crud-demo/modules/hospital/t_reset_password"
	"fiber-crud-demo/modules/hospital/t_token"
	"fiber-crud-demo/modules/hospital/t_treatment_discount"

	"fiber-crud-demo/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializer.LoadEnvironmentVariables()
	initializer.LoggerInit()
	initializer.InitializeDatabase()
}

func config() fiber.Config {
	htmlEngine := html.New("./web/templates", ".html")
	return fiber.Config{
		// Prefork:               true,
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     4 * 1024 * 1024,
		// DisableStartupMessage: true,
		ServerHeader: "Fiber Audio Management",
		AppName:      "Fiber Audio Management v0.0.1",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			res := &response.Response{}
			res.ErrMessage(c.Path(), code, err.Error())

			return c.Status(code).JSON(res)
		},
		Views:       htmlEngine,
		ViewsLayout: "layouts/main",
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	app := fiber.New(config())

	// ### Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // v3 expects a slice of strings for origins
		AllowHeaders: []string{"*"},
	}))
	// app.Use(cache.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	// app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swaggo.HandlerDefault)

	app.Get("/public/*", static.New("./web/public"))

	// ### Routes
	routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve on "+host+" port "+port)

	// ### Run
	log.Fatal(app.Listen(host + ":" + port))
}

func routes(app *fiber.App) {

	// HEALTH
	health_api := app.Group("/v1/health")
	health_api.Get("/status", health.Status)

	// HELLO WORLD
	hello_world_api := app.Group("/v1/hello-world")
	hello_world_api.Get("", hello_world.HelloWorld)
	hello_world_api.Get("/path/:message", hello_world.HelloWorldPath)
	hello_world_api.Get("/query", hello_world.HelloWorldQuery)
	hello_world_api.Post("/payload", hello_world.HelloWorldPayload)
	hello_world_api.Get("/error/:type", hello_world.HelloWorldError)

	var validate = validator.New()

	m_admin.GetRouter(app, validate)
	m_bank.GetRouter(app, validate)
	m_biodata.GetRouter(app, validate)
	m_biodata_address.GetRouter(app, validate)
	m_blood_group.GetRouter(app, validate)
	m_courier.GetRouter(app, validate)
	m_customer.GetRouter(app, validate)
	m_customer_member.GetRouter(app, validate)
	m_customer_relation.GetRouter(app, validate)
	m_doctor.GetRouter(app, validate)
	m_doctor_education.GetRouter(app, validate)
	m_education_level.GetRouter(app, validate)
	m_location.GetRouter(app, validate)
	m_location_level.GetRouter(app, validate)
	m_medical_facility.GetRouter(app, validate)
	m_medical_facility_category.GetRouter(app, validate)
	m_medical_facility_schedule.GetRouter(app, validate)
	m_medical_item.GetRouter(app, validate)
	m_medical_item_category.GetRouter(app, validate)
	m_medical_item_segmentation.GetRouter(app, validate)
	m_menu.GetRouter(app, validate)
	m_menu_role.GetRouter(app, validate)
	m_payment_method.GetRouter(app, validate)
	m_role.GetRouter(app, validate)
	m_specialization.GetRouter(app, validate)
	m_user.GetRouter(app, validate)
	m_biodata_attachment.GetRouter(app, validate)
	m_wallet_default_nominal.GetRouter(app, validate)
	t_appointment.GetRouter(app, validate)
	t_appointment_cancellation.GetRouter(app, validate)
	t_appointment_done.GetRouter(app, validate)
	t_appointment_reschedule_history.GetRouter(app, validate)
	t_current_doctor_specialization.GetRouter(app, validate)
	t_customer_chat.GetRouter(app, validate)
	t_customer_chat_history.GetRouter(app, validate)
	t_customer_custom_nominal.GetRouter(app, validate)
	t_customer_registered_card.GetRouter(app, validate)
	t_customer_va.GetRouter(app, validate)
	t_customer_va_history.GetRouter(app, validate)
	t_customer_wallet.GetRouter(app, validate)
	t_customer_wallet_top_up.GetRouter(app, validate)
	t_doctor_office.GetRouter(app, validate)
	t_doctor_office_schedule.GetRouter(app, validate)
	t_doctor_office_treatment.GetRouter(app, validate)
	t_doctor_office_treatment_price.GetRouter(app, validate)
	t_doctor_treatment.GetRouter(app, validate)
	t_medical_item_purchase.GetRouter(app, validate)
	t_medical_item_purchase_detail.GetRouter(app, validate)
	t_reset_password.GetRouter(app, validate)
	t_token.GetRouter(app, validate)
	t_treatment_discount.GetRouter(app, validate)
	m_courier_type.GetRouter(app, validate)
	t_courier_discount.GetRouter(app, validate)
	t_customer_wallet_withdraw.GetRouter(app, validate)

}
