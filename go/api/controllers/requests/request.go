package requests

// Requête personnalisée que l'on va binder sur le model de "User" enregistré en base de données
type SignupRequest struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Address   string `form:"address" json:"address" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Photo     string `form:"photo" json:"photo"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// Requête pour pouvoir se loguer
type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type EmailRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type ResetPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type InvitationGroupRequest struct {
	Email   string `form:"email" json:"email" binding:"required"`
	GroupID uint64 `form:"group_id" json:"group_id" binding:"required"`
}

type GroupRequest struct {
	Budget float32 `json:"budget"`
	Nom    string  `json:"nom"`
}

type UpdateBudgetRequest struct {
	Budget float32 `json:"budget"`
}
