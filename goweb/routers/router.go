// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact 1834327029@qq.com
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/controllers"
)

// 相关路由配置: https://beego.me/docs/mvc/controller/router.md

func init() {
	// 命名空间配置路由
	//初始化 游客命名空间
	visitor :=
		web.NewNamespace("/visitor/v1",
			web.NSRouter("/login", &controllers.VisitorController{}, "options,post:Login"),
			web.NSRouter("/register", &controllers.VisitorController{}, "post:Register"),
		)
	web.AddNamespace(visitor)

	// 登录用户命名空间
	user :=
		web.NewNamespace("/user/v1",
			web.NSCond(controllers.Identify),
			web.NSRouter("/version", &controllers.MainController{}),
			web.NSRouter("/logout", &controllers.UserController{}, "get:Logout"),
			web.NSNamespace("/course",
				web.NSRouter("/info", &controllers.CourseController{}, "get,post:CourseInfo"),
				web.NSRouter("/create", &controllers.CourseController{}, "post:CreateCourse"),
				web.NSRouter("/show", &controllers.CourseController{}, "get:GetCourse"),
				web.NSRouter("/export", &controllers.CourseController{}, "get:ExportCourse"),
				web.NSRouter("/choose", &controllers.CourseController{}, "get,post:ChooseCourse"),
				web.NSRouter("/grade", &controllers.CourseController{}, "get,post:CourseGrade"),
				web.NSRouter("/class", &controllers.CourseController{}, "get:CourseClass"),
			),
			web.NSNamespace("/student",
				web.NSRouter("/info", &controllers.StudentController{}, "get,post:StudentInfo"),
				web.NSRouter("/create", &controllers.StudentController{}, "post:CreateStudent"),
				web.NSRouter("/list", &controllers.StudentController{}, "get:GetStudentList"),
				web.NSRouter("/sort", &controllers.StudentController{}, "get:GetGradeSortStudentList"),
			),
			web.NSNamespace("/class",
				web.NSRouter("/info", &controllers.ClassController{}, "get,post:ClassInfo"),
				web.NSRouter("/create", &controllers.ClassController{}, "post:CreateClass"),
				web.NSRouter("/list", &controllers.ClassController{}, "get:GetClassList"),
				web.NSRouter("/course", &controllers.ClassController{}, "get,post:ClassCourse"),
				web.NSRouter("/gradeSort", &controllers.ClassController{}, "get:ClassSort"),
			),
			web.NSNamespace("/teacher",
				web.NSRouter("/info", &controllers.TeacherController{}, "get,post:TeacherInfo"),
				web.NSRouter("/create", &controllers.TeacherController{}, "post:CreateTeacher"),
				web.NSRouter("/list", &controllers.TeacherController{}, "get:GetTeachers"),
			),
			//web.NSNamespace("/class",
			//	web.NSRouter("/info", &controllers.ClassController{}, "get,post:ClassInfo"),
			//	web.NSRouter("/export", &controllers.ClassController{}, "get:ExportCourse"),
			//	web.NSRouter("/choose", &controllers.ClassController{}, "get:ChooseCourse"),
			//	web.NSRouter("/grade", &controllers.ClassController{}, "get,post:CourseGrade"),
			//),
		)
	web.AddNamespace(user)

	//初始化 admin命名空间
	admin :=
		web.NewNamespace("/admin/v1",
			web.NSCond(func(ctx *context.Context) bool {
				return true
			}),
			web.NSGet("/get", func(ctx *context.Context) {
				ctx.Output.Body([]byte("管理员get"))
			}),
			web.NSRouter("/version", &controllers.AdminController{}, "get:ShowAPIVersion"),
			web.NSNamespace("/show",
				web.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("展示管理员"))
				}),
			),
		)
	//注册 namespace
	web.AddNamespace(admin)
}
