package models

type UserFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Occupation string `json:"occupation"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

type ProfileFormatter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Email: user.Email,
		Token: token,
	}

	return formatter
}

func FormatProfile(profile User) ProfileFormatter {
	formatter := ProfileFormatter{
		Email:    profile.Email,
		Name:     profile.Name,
		ImageURL: "this url not found",
	}

	return formatter
}
