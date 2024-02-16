package route

var (
	ErrTryLoginAgain             = response("Unable to login with those details - please try again!")
	ErrTryRegisterAgain          = response("Unable to register with those details - please try again!")
	ErrTryOrganizationAgain      = response("Unable to create or access that organization - please try again!")
	ErrTryProfileAgain           = response("Unable to create or access user profile - please try again!")
	ErrTryResendAgain            = response("Unable to resend email - please try again!")
	ErrMemberNotFound            = response("Team member with the specified ID was not found.")
	ErrMissingOrganizationName   = response("Organization name is required.")
	ErrMissingOrganizationDomain = response("Organization domain is required.")
	ErrOrganizationNotFound      = response("Organization with the specified ID was not found.")
	ErrLogBackIn                 = response("Logged out of your account - please log back in!")
	ErrVerifyEmail               = response("Please verify your email address and try again!")
	ErrInvalidEmail              = response("Please enter a valid email address.")
	ErrVerificationFailed        = response("Email verification failed. Please contact support@datum.net for assistance.")
	ErrSendPasswordResetFailed   = response("Unable to send password reset email. Please contact support@datum.net for assistance.")
	ErrPasswordResetFailed       = response("Unable to reset your password. Please contact support@datum.net for assistance.")
	ErrRequestNewInvite          = response("Invalid invitation link - please request a new one!")
	ErrSomethingWentWrong        = response("Oops - something went wrong!")
	ErrBadResendRequest          = response("Unable to resend email - please update request and try again.")
	ErrRequestNewReset           = response("Unable to reset your password - please request a new password reset.")

	AllResponses = map[string]struct{}{}
)

// response creates a standard error message to ensure uniqueness and testability for external packages
func response(msg string) string {
	if _, ok := AllResponses[msg]; ok {
		panic("duplicate error response defined: " + msg)
	}

	AllResponses[msg] = struct{}{}

	return msg
}

// StatusError decodes an error response from datum.
type StatusError struct {
	StatusCode int
	Reply      Reply
}

// Reply contains standard fields that are used for generic API responses and errors.
type Reply struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Unverified bool   `json:"unverified,omitempty"`
}
