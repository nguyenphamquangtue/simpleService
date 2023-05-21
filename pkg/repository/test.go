package repository

// import (
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Userzz struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Name     string `gorm:"size:255"`
// 	Age      int
// 	Password string `gorm:"size:255"`
// }

// func main() {
// 	// Thông tin kết nối
// 	dsn := "host=localhost user=your_username password=your_password dbname=your_database port=5432 sslmode=disable"

// 	// Khởi tạo kết nối với cơ sở dữ liệu
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}

// 	// Đồng bộ hóa cấu trúc bảng
// 	err = db.AutoMigrate(&User{})
// 	if err != nil {
// 		log.Fatal("Failed to migrate database:", err)
// 	}

// 	// Khởi tạo instance của Echo
// 	e := echo.New()

// 	// Route xử lý GET "/users"
// 	e.GET("/users", func(c echo.Context) error {
// 		// Truy vấn dữ liệu từ bảng "users"
// 		var users []User
// 		result := db.Find(&users)
// 		if result.Error != nil {
// 			log.Println("Failed to fetch users:", result.Error)
// 			return c.String(http.StatusInternalServerError, "Failed to fetch users")
// 		}

// 		// Xóa mật khẩu trước khi trả về dữ liệu
// 		for i := range users {
// 			users[i].Password = ""
// 		}

// 		// Trả về danh sách người dùng
// 		return c.JSON(http.StatusOK, users)
// 	})

// 	// Route xử lý POST "/users"
// 	e.POST("/users", func(c echo.Context) error {
// 		// Lấy thông tin người dùng từ request body
// 		var user User
// 		if err := c.Bind(&user); err != nil {
// 			return c.String(http.StatusBadRequest, "Invalid request body")
// 		}

// 		// Mã hóa mật khẩu
// 		hashedPassword, err := hashPassword(user.Password)
// 		if err != nil {
// 			log.Println("Failed to hash password:", err)
// 			return c.String(http.StatusInternalServerError, "Failed to create user")
// 		}

// 		// Lưu thông tin người dùng vào cơ sở dữ liệu
// 		user.Password = hashedPassword
// 		result := db.Create(&user)
// 		if result.Error != nil {
// 			log.Println("Failed to create user:", result.Error)
// 			return c.String(http.StatusInternalServerError, "Failed to create user")
// 		}

// 		// Xóa mật khẩu trước khi trả về dữ liệu
// 		user.Password = ""

// 		// Trả về người dùng đã được tạo
// 		return c.JSON(http.StatusCreated, user)
// 	})

// 	// Route xử lý POST "/login"
// 	e.POST("/login", func(c echo.Context) error {
// 		// Lấy thông tin đăng nhập từ request body
// 		var loginData struct {
// 			Username string `json:"username"`
// 			Password string `json:"password"`
// 		}
// 		if err := c.Bind(&loginData); err != nil {
// 			return c.String(http.StatusBadRequest, "Invalid request body")
// 		}

// 		// Truy vấn người dùng từ cơ sở dữ liệu
// 		var user User
// 		result := db.Where("name = ?", loginData.Username).First(&user)
// 		if result.Error != nil {
// 			log.Println("Failed to find user:", result.Error)
// 			return c.String(http.StatusUnauthorized, "Invalid username or password")
// 		}

// 		// So sánh mật khẩu
// 		if err := comparePassword(user.Password, loginData.Password); err != nil {
// 			log.Println("Failed to compare password:", err)
// 			return c.String(http.StatusUnauthorized, "Invalid username or password")
// 		}

// 		// Xóa mật khẩu trước khi trả về dữ liệu
// 		user.Password = ""

// 		// Trả về thông tin người dùng đã đăng nhập thành công
// 		return c.JSON(http.StatusOK, user)
// 	})

// 	// Chạy server Echo trên cổng 8080
// 	e.Start(":8080")
// }

// // Hash mật khẩu
// func hashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPassword), nil
// }

// // So sánh mật khẩu với hash đã được lưu trữ
// func comparePassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }
