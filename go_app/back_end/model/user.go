package model

// User represents the user model from Auth0.
type User struct {
	Email         string       `json:"email"`
	Picture       string       `json:"picture"`
	Name          string       `json:"name"`
	Nickname      string       `json:"nickname"`
	Metadata      UserMetadata `json:"user_metadata"`
	EmailVerified bool         `json:"email_verified"`
	UserID        string       `json:"user_id"`
}

// UserMetadata represents the user_metadata field's structure from Auth0.
type UserMetadata struct {
	Courses []string `json:"courses"`
}
