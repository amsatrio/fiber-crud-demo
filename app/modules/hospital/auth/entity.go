package auth

type Auth struct {
	MainToken    *string `form:"mainToken" json:"mainToken" xml:"mainToken"`
	RefreshToken *string `form:"refreshToken" json:"refreshToken" xml:"refreshToken"`
	ExpiredIn    *int64  `form:"expiredIn" json:"expiredIn" xml:"expiredIn"`
}

type AuthLoginRequest struct {
	Username *string `form:"username" json:"username" xml:"username"`
	Password *string `form:"password" json:"password" xml:"password"`
}

type AuthRegisterRequest struct {
	Username *string `form:"username" json:"username" xml:"username"`
	Password *string `form:"password" json:"password" xml:"password"`
}

type AuthRefreshTokenRequest struct {
	MainToken    *string `form:"mainToken" json:"mainToken" xml:"mainToken"`
	RefreshToken *string `form:"refreshToken" json:"refreshToken" xml:"refreshToken"`
}

type AuthResetPasswordRequest struct {
	Username    *string `form:"username" json:"username" xml:"username"`
	NewPassword *string `form:"newPassword" json:"newPassword" xml:"newPassword"`
}
