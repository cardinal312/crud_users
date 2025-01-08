// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/cardinal312/crud_users/handlers"
// 	"github.com/cardinal312/crud_users/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func SetupTestDB() {
// 	dsn := "host=localhost user=postgres password=example dbname=exampledb port=5432 sslmode=disable"
// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to the database")
// 	}
// 	db.AutoMigrate(&models.User{})
// }

// func TestCreateUser(t *testing.T) {
// 	e := echo.New()

// 	user := models.User{
// 		Firstname: "John",
// 		Lastname:  "Doe",
// 		Email:     "john.doe@example.com",
// 		Age:       30,
// 	}

// 	// Mock the request
// 	jsonData, _ := json.Marshal(user)
// 	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonData))
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Call the handler
// 	if assert.NoError(t, handlers.CreateUser(db)(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 	}
// }

// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/cardinal312/crud_users/handlers" // Путь к функции CreateUser
// 	"github.com/cardinal312/crud_users/models"   // Путь к функции CreateUser
// 	"github.com/jinzhu/gorm"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Мок для gorm.DB
// type MockDB struct {
// 	mock.Mock
// }

// //var db *gorm.DB

// func (m *MockDB) Create(value interface{}) *gorm.DB {
// 	args := m.Called(value)
// 	return args.Get(0).(*gorm.DB)
// }

// // Тест для успешного создания пользователя
// func TestCreateUser_Success(t *testing.T) {
// 	e := echo.New()
// 	reqBody := map[string]interface{}{
// 		"username": "testuser",
// 		"email":    "testuser@example.com",
// 		"password": "password123",
// 	}
// 	reqBodyBytes, _ := json.Marshal(reqBody)
// 	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(reqBodyBytes))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Создаем мок db
// 	mockDB := new(MockDB)
// 	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

// 	// Создаем хендлерs
// 	handler := handlers.CreateUser(mockDB)

// 	// Выполняем запрос
// 	if assert.NoError(t, handler(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		var resp models.User
// 		if err := json.NewDecoder(rec.Body).Decode(&resp); assert.NoError(t, err) {
// 			assert.Equal(t, "testuser", resp)
// 			assert.Equal(t, "testuser@example.com", resp.Email)
// 		}
// 	}

// 	// Проверяем, что метод Create был вызван
// 	mockDB.AssertExpectations(t)
// }

// // Тест для ошибки при привязке данных
// func TestCreateUser_BindError(t *testing.T) {
// 	e := echo.New()
// 	reqBody := map[string]interface{}{
// 		"username": "testuser",
// 		// Пропускаем поле email, чтобы вызвать ошибку привязки
// 		"password": "password123",
// 	}
// 	reqBodyBytes, _ := json.Marshal(reqBody)
// 	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(reqBodyBytes))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Создаем мок db
// 	mockDB := new(MockDB)

// 	// Создаем хендлер
// 	handler := handlers.CreateUser(mockDB)

// 	// Выполняем запрос
// 	if assert.NoError(t, handler(c)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)

// 		var resp models.Response
// 		if err := json.NewDecoder(rec.Body).Decode(&resp); assert.NoError(t, err) {
// 			assert.Equal(t, "Error", resp.Status)
// 			assert.Equal(t, "Could not add the user", resp.Message)
// 		}
// 	}

// 	// Проверяем, что метод Create не был вызван
// 	mockDB.AssertExpectations(t)
// }

// // Тест для ошибки при создании пользователя в базе данных
// func TestCreateUser_DBError(t *testing.T) {
// 	e := echo.New()
// 	reqBody := map[string]interface{}{
// 		"username": "testuser",
// 		"email":    "testuser@example.com",
// 		"password": "password123",
// 	}
// 	reqBodyBytes, _ := json.Marshal(reqBody)
// 	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(reqBodyBytes))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Создаем мок db
// 	mockDB := new(MockDB)
// 	mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: assert.AnError})

// 	// Создаем хендлер
// 	handler := handlers.CreateUser(mockDB)

// 	// Выполняем запрос
// 	if assert.NoError(t, handler(c)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)

// 		var resp models.Response
// 		if err := json.NewDecoder(rec.Body).Decode(&resp); assert.NoError(t, err) {
// 			assert.Equal(t, "Error", resp.Status)
// 			assert.Equal(t, "Could not create the user", resp.Message)
// 		}
// 	}

// 	// Проверяем, что метод Create был вызван
// 	mockDB.AssertExpectations(t)
// }
