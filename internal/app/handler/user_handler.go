package handler

import (
	"database/sql"
	"github/ardaberrun/credit-app-go/internal/app/model"
	"github/ardaberrun/credit-app-go/internal/app/service"
	"github/ardaberrun/credit-app-go/internal/app/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


type UserHandler struct {
	userService service.IUserService
	router *gin.Engine
}

func InitializeUserHandler(us service.IUserService, r *gin.Engine) *UserHandler {
	return &UserHandler{userService: us, router: r};
}

func (h *UserHandler) RegisterRoutes() {
	user := h.router.Group("/users")
	{
		user.POST("/register", h.CreateUser);
		user.POST("/login", h.Login);
		user.GET("/", h.GetUsers);
		user.GET("/:id", h.GetUserById);
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var cur model.UserRequest;

	if err := c.ShouldBindJSON(&cur); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	user, err := h.userService.GetUserByEmail(cur.Email);
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"});

		return;
	}

	hpwd, err := utils.HashPassword(cur.Password);
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	newUser := &model.User{
		Email: cur.Email,
		HashedPassword: hpwd,
		AccountNumber: int64(rand.Intn(1000000)),
		Balance: 0,
		CreatedAt: time.Now().UTC(),
	};


	if err := h.userService.CreateUser(newUser); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"});
}

func (h *UserHandler) Login(c *gin.Context) {
	var cur model.UserRequest;

	if err := c.ShouldBindJSON(&cur); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	user, err := h.userService.GetUserByEmail(cur.Email);
	if user == nil || err == sql.ErrNoRows {
		c.JSON(http.StatusBadGateway, gin.H{"error": "User not found"});

		return;
	} else if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	isCompared := true;

	if (user.RoleName == "admin") {
		if user.HashedPassword != cur.Password {
			isCompared = false;
		}
	} else {
		isCompared = utils.ComparePassword(user.HashedPassword, cur.Password);
	}

	if !isCompared {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"});

		return;
	}

	token, err := utils.GenerateJWT(user.Id, user.RoleName);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create jwt"});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "User login successfully", "token": token});
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	userRole := claims.(jwt.MapClaims)["user_role"].(string);

	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	users, err := h.userService.GetUsers();
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": users});
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	claims, exists := c.Get("claims");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;
	}

	uid, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"});

		return;
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64));
	userRole := claims.(jwt.MapClaims)["user_role"].(string);

	if userRole != "admin" && userId != uid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"});

		return;		
	}

	user, err := h.userService.GetUserById(uid);
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"});

		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": user});
}