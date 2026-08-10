package v1

// ActivityCategory is the shared, cross-service vocabulary for ActivityEvent.Category.
type ActivityCategory string

// ActivitySubcategory is the shared, cross-service vocabulary for ActivityEvent.Subcategory.
type ActivitySubcategory string

const (
	ActivityCategoryAuth     ActivityCategory = "AUTH"     // auth-service: login, logout, 2FA, signup
	ActivityCategoryUser     ActivityCategory = "USER"     // auth-service: password/otp/account setting changes
	ActivityCategoryInvestor ActivityCategory = "INVESTOR" // investor-service: investor profile & broker enrollment
)

// Subcategories for ActivityCategoryAuth.
const (
	ActivitySubcategoryLogin          ActivitySubcategory = "LOGIN"
	ActivitySubcategoryLoginFailed    ActivitySubcategory = "LOGIN_FAILED"
	ActivitySubcategoryLogout         ActivitySubcategory = "LOGOUT"
	ActivitySubcategorySignupComplete ActivitySubcategory = "SIGNUP_COMPLETE"
)

// Subcategories for ActivityCategoryUser.
const (
	ActivitySubcategoryPasswordChange       ActivitySubcategory = "PASSWORD_CHANGE"
	ActivitySubcategoryPasswordResetRequest ActivitySubcategory = "PASSWORD_RESET_REQUEST"
	ActivitySubcategoryPasswordReset        ActivitySubcategory = "PASSWORD_RESET"
	ActivitySubcategoryOtpChange            ActivitySubcategory = "OTP_CHANGE"
	ActivitySubcategoryOtpReset             ActivitySubcategory = "OTP_RESET"
)

// Subcategories for ActivityCategoryInvestor.
const (
	ActivitySubcategoryInvestorProfileUpdated ActivitySubcategory = "INVESTOR_PROFILE_UPDATED"
)
