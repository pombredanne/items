package common

import (
	"zonstfe_api/common/controllers"
	"zonstfe_api/common/middlewares"
	"zonstfe_api/common/my_context"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func InitRouter(r *chi.Mux, context *my_context.Context, mw *middlewares.MiddleWares) *chi.Mux {
	upload := controllers.NewFileController(context)
	user := controllers.NewUserController(context)
	r.Group(func(r chi.Router) {
		r.Post("/upload", upload.FileUpload)
		//后台不开启注册功能
		if context.RoleID == 1 || context.RoleID==3{
			r.Post("/signup", user.Signup)
			r.Post("/signup/{email}", user.SignupEmail)
		}
		r.Post("/login", user.Login)

	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(context.TokenAuth))
		r.Use(mw.Authenticator)
		r.Post("/logout", user.Logout)
	})
	return r
}
