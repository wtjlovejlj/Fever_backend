package routers

import (
	"Fever_backend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/fever/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "I'll love you till I die")
	})
	//登录
	v1.POST("/login", controller.LoginHandler)
	//查看所有发热人员的名单
	v1.GET("/Fare_all", controller.FareListHandler)
	// 创新新的发热人员信息
	v1.POST("/addFare", controller.AddMessage)
	//修改信息
	v1.POST("/updateFare/:feverId", controller.UpdateMessage)
	//删除信息
	v1.POST("/deleteFare/:feverId", controller.DeleteMessage)

	//登录验证token
	v1.Use(controller.JWTAuthMiddleware())
	{
		//修改密码
		v1.POST("/change_password", controller.ChangePasswordHandler)
		//加权限
		v1.POST("/add_casbin", controller.AddCasbin)

	}

	//医生名单相关接口
	v2 := r.Group("/doctor/v1")

	//登录验证token
	v2.Use(controller.JWTAuthMiddleware(), controller.AuthCheckRole())
	{
		//查看所有医生的名单
		v2.GET("/information_all", controller.DoctorListHandler) //66666666666666666666666666666666666
		//添加医生
		v2.POST("/add_doc", controller.AddDoctorHandler)
		//修改医生的信息
		v2.POST("/change_doc", controller.ChangeDoctorHandler)
		//删除医生
		v2.POST("/delete_doc", controller.DeleteDoctorHandler)
		//修改当前医生管理员的信息
		v2.POST("/update_myself", controller.UpdateMyMessage)
	}

	//医生名单相关接口
	v3 := r.Group("/county/v1")

	//登录验证token
	v3.Use(controller.JWTAuthMiddleware(), controller.AuthCheckRole())
	{
		//查看所有区县的名单
		v3.GET("/district_all", controller.DistrictListHandler)
		//修改区县的信息
		v3.POST("/change_con", controller.ChangeCountyHandler)
		//删除区县的信息
		v3.POST("/delete_con", controller.DeleteCountyHandler)
		//通过区县名查找区县的所有医院
		v3.POST("/find_con_his", controller.FindCountyHandler)

	}

	v4 := r.Group("/hospital/v1")
	v4.Use(controller.JWTAuthMiddleware(), controller.AuthCheckRole())
	{
		//增加院长
		v4.POST("/addHospitalHandler", controller.AddHospitalHandler)
		//展示医院
		v4.GET("/hospitalAdminAll", controller.DetailHandle) //555555555555555555555555555555
		//点击院长获取本院的所有医生（点击人员按钮）
		v4.GET("/show", controller.ShowDoctorHandle)
		//删除院长
		v4.POST("/delete", controller.DeleteDetailHandle)
		//修改院长
		v4.POST("/hospitalAdmin", controller.UpdateDetailHandle)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
