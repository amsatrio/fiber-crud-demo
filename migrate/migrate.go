package main

import (
	"log"

	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/modules/german_vocabulary/nomen"
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
	"gorm.io/gorm"
)

func main() {
	initializer.LoadEnvironmentVariables()
	initializer.InitializeDatabaseGerman()
	initializer.InitializeDatabaseHospital()

	migrateGerman()
	migrateHospital()

	log.Println("migration completed successfully")
}

func migrateGerman() {
	if initializer.DB_GERMAN == nil {
		log.Fatal("DB_GERMAN is not connected")
	}

	log.Println("migrating DB_GERMAN...")
	if err := initializer.DB_GERMAN.AutoMigrate(
		&nomen.Noun{},
	); err != nil {
		log.Fatalf("DB_GERMAN migration failed: %v", err)
	}
	log.Println("DB_GERMAN migration done")
}

func migrateHospital() {
	if initializer.DB_HOSPITAL == nil {
		log.Fatal("DB_HOSPITAL is not connected")
	}

	log.Println("migrating DB_HOSPITAL...")
	if err := autoMigrate(initializer.DB_HOSPITAL,
		&m_admin.MAdmin{},
		&m_bank.MBank{},
		&m_biodata.MBiodata{},
		&m_biodata_address.MBiodataAddress{},
		&m_biodata_attachment.MBiodataAttachment{},
		&m_blood_group.MBloodGroup{},
		&m_courier.MCourier{},
		&m_courier_type.MCourierType{},
		&m_customer.MCustomer{},
		&m_customer_member.MCustomerMember{},
		&m_customer_relation.MCustomerRelation{},
		&m_doctor.MDoctor{},
		&m_doctor_education.MDoctorEducation{},
		&m_education_level.MEducationLevel{},
		&m_location.MLocation{},
		&m_location_level.MLocationLevel{},
		&m_medical_facility.MMedicalFacility{},
		&m_medical_facility_category.MMedicalFacilityCategory{},
		&m_medical_facility_schedule.MMedicalFacilitySchedule{},
		&m_medical_item.MMedicalItem{},
		&m_medical_item_category.MMedicalItemCategory{},
		&m_medical_item_segmentation.MMedicalItemSegmentation{},
		&m_menu.MMenu{},
		&m_menu_role.MMenuRole{},
		&m_payment_method.MPaymentMethod{},
		&m_role.MRole{},
		&m_specialization.MSpecialization{},
		&m_user.MUser{},
		&m_wallet_default_nominal.MWalletDefaultNominal{},
		&t_appointment.TAppointment{},
		&t_appointment_cancellation.TAppointmentCancellation{},
		&t_appointment_done.TAppointmentDone{},
		&t_appointment_reschedule_history.TAppointmentRescheduleHistory{},
		&t_courier_discount.TCourierDiscount{},
		&t_current_doctor_specialization.TCurrentDoctorSpecialization{},
		&t_customer_chat.TCustomerChat{},
		&t_customer_chat_history.TCustomerChatHistory{},
		&t_customer_custom_nominal.TCustomerCustomNominal{},
		&t_customer_registered_card.TCustomerRegisteredCard{},
		&t_customer_va.TCustomerVa{},
		&t_customer_va_history.TCustomerVaHistory{},
		&t_customer_wallet.TCustomerWallet{},
		&t_customer_wallet_top_up.TCustomerWalletTopUp{},
		&t_customer_wallet_withdraw.TCustomerWalletWithdraw{},
		&t_doctor_office.TDoctorOffice{},
		&t_doctor_office_schedule.TDoctorOfficeSchedule{},
		&t_doctor_office_treatment.TDoctorOfficeTreatment{},
		&t_doctor_office_treatment_price.TDoctorOfficeTreatmentPrice{},
		&t_doctor_treatment.TDoctorTreatment{},
		&t_medical_item_purchase.TMedicalItemPurchase{},
		&t_medical_item_purchase_detail.TMedicalItemPurchaseDetail{},
		&t_reset_password.TResetPassword{},
		&t_token.TToken{},
		&t_treatment_discount.TTreatmentDiscount{},
	); err != nil {
		log.Fatalf("DB_HOSPITAL migration failed: %v", err)
	}
	log.Println("DB_HOSPITAL migration done")
}

func autoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
