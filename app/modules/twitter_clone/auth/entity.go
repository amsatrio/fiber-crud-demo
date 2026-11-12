package auth

type Auth struct {
	MainToken    *string `form:"mainToken" json:"mainToken" xml:"mainToken"`
	RefreshToken *string `form:"refreshToken" json:"refreshToken" xml:"refreshToken"`
	ExpiredIn    *int64  `form:"expiredIn" json:"expiredIn" xml:"expiredIn"`
}

type AuthLoginRequest struct {
	Username *string `form:"username" json:"username" xml:"username" validate:"required"`
	Password *string `form:"password" json:"password" xml:"password" validate:"required"`
}

type AuthRegisterRequest struct {
	Username    *string `form:"username" json:"username" xml:"username" validate:"required,max=30"`
	Email       *string `form:"email" json:"email" xml:"email" validate:"required,max=255,email"`
	Password    *string `form:"password" json:"password" xml:"password" validate:"required,min=6,max=255"`
	DisplayName *string `form:"displayName" json:"displayName" xml:"displayName" validate:"max=50"`
}

type AuthRefreshTokenRequest struct {
	MainToken    *string `form:"mainToken" json:"mainToken" xml:"mainToken"`
	RefreshToken *string `form:"refreshToken" json:"refreshToken" xml:"refreshToken"`
}

type AuthResetPasswordRequest struct {
	Username    *string `form:"username" json:"username" xml:"username"`
	NewPassword *string `form:"newPassword" json:"newPassword" xml:"newPassword"`
}
