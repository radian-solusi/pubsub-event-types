package v1

// NotificationCategory is the shared, cross-service vocabulary for NotificationEvent.Category.
type NotificationCategory string

// NotificationSubcategory is the shared, cross-service vocabulary for NotificationEvent.Subcategory.
type NotificationSubcategory string

// Notification categories already in production use.
const (
	NotificationCategoryNotification NotificationCategory = "NOTIFICATION"
)

// Subcategories for NotificationCategoryNotification.
const (
	NotificationSubcategoryPasswordReset         NotificationSubcategory = "PASSWORD_RESET"
	NotificationSubcategoryEmailVerification     NotificationSubcategory = "EMAIL_VERIFICATION"
	NotificationSubcategoryOtpSms                NotificationSubcategory = "OTP_SMS"
	NotificationSubcategoryOtpResetRequest       NotificationSubcategory = "OTP_RESET_REQUEST"
	NotificationSubcategoryInvestorProfileUpdate NotificationSubcategory = "INVESTOR_PROFILE_UPDATE"
)
