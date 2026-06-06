package hospital

import (
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/auth"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_admin"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_bank"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_biodata"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_biodata_address"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_biodata_attachment"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_blood_group"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_courier"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_courier_type"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_customer"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_customer_member"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_customer_relation"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_doctor"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_doctor_education"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_education_level"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_location"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_location_level"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_facility"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_facility_category"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_facility_schedule"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_item"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_item_category"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_medical_item_segmentation"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_menu"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_menu_role"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_payment_method"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_role"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_specialization"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_user"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_wallet_default_nominal"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_appointment"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_appointment_cancellation"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_appointment_done"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_appointment_reschedule_history"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_courier_discount"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_current_doctor_specialization"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_chat"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_chat_history"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_custom_nominal"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_registered_card"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_va"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_va_history"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_wallet"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_wallet_top_up"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_customer_wallet_withdraw"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_doctor_office"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_doctor_office_schedule"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_doctor_office_treatment"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_doctor_office_treatment_price"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_doctor_treatment"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_medical_item_purchase"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_medical_item_purchase_detail"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_reset_password"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_token"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/t_treatment_discount"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	var validate = validator.New()

	api := app.Group("/v1/hospital")

	auth.GetRouter(api, validate)

	m_admin.GetRouter(api, validate)
	m_bank.GetRouter(api, validate)
	m_biodata.GetRouter(api, validate)
	m_biodata_address.GetRouter(api, validate)
	m_blood_group.GetRouter(api, validate)
	m_courier.GetRouter(api, validate)
	m_customer.GetRouter(api, validate)
	m_customer_member.GetRouter(api, validate)
	m_customer_relation.GetRouter(api, validate)
	m_doctor.GetRouter(api, validate)
	m_doctor_education.GetRouter(api, validate)
	m_education_level.GetRouter(api, validate)
	m_location.GetRouter(api, validate)
	m_location_level.GetRouter(api, validate)
	m_medical_facility.GetRouter(api, validate)
	m_medical_facility_category.GetRouter(api, validate)
	m_medical_facility_schedule.GetRouter(api, validate)
	m_medical_item.GetRouter(api, validate)
	m_medical_item_category.GetRouter(api, validate)
	m_medical_item_segmentation.GetRouter(api, validate)
	m_menu.GetRouter(api, validate)
	m_menu_role.GetRouter(api, validate)
	m_payment_method.GetRouter(api, validate)
	m_role.GetRouter(api, validate)
	m_specialization.GetRouter(api, validate)
	m_user.GetRouter(api, validate)
	m_biodata_attachment.GetRouter(api, validate)
	m_wallet_default_nominal.GetRouter(api, validate)
	t_appointment.GetRouter(api, validate)
	t_appointment_cancellation.GetRouter(api, validate)
	t_appointment_done.GetRouter(api, validate)
	t_appointment_reschedule_history.GetRouter(api, validate)
	t_current_doctor_specialization.GetRouter(api, validate)
	t_customer_chat.GetRouter(api, validate)
	t_customer_chat_history.GetRouter(api, validate)
	t_customer_custom_nominal.GetRouter(api, validate)
	t_customer_registered_card.GetRouter(api, validate)
	t_customer_va.GetRouter(api, validate)
	t_customer_va_history.GetRouter(api, validate)
	t_customer_wallet.GetRouter(api, validate)
	t_customer_wallet_top_up.GetRouter(api, validate)
	t_doctor_office.GetRouter(api, validate)
	t_doctor_office_schedule.GetRouter(api, validate)
	t_doctor_office_treatment.GetRouter(api, validate)
	t_doctor_office_treatment_price.GetRouter(api, validate)
	t_doctor_treatment.GetRouter(api, validate)
	t_medical_item_purchase.GetRouter(api, validate)
	t_medical_item_purchase_detail.GetRouter(api, validate)
	t_reset_password.GetRouter(api, validate)
	t_token.GetRouter(api, validate)
	t_treatment_discount.GetRouter(api, validate)
	m_courier_type.GetRouter(api, validate)
	t_courier_discount.GetRouter(api, validate)
	t_customer_wallet_withdraw.GetRouter(api, validate)

}
