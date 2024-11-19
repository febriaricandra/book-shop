package services

import (
	"fmt"
	"strconv"

	"github.com/febriaricandra/book-shop/config"
	"github.com/febriaricandra/book-shop/internal/models"
	"github.com/febriaricandra/book-shop/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"time"
)

type UserService interface {
	RegisterUser(username, email, password string) error
	LoginUser(email, password string) (string, string, error) //return access token and refresh token
	RefreshToken(refreshToken string) (string, error)         // return access token
	VerifyToken(token string) (*JWTCustomClaims, error)       // return claims
}

var jwtSecret = config.LoadConfig().JWTSecret

type JWTCustomClaims struct {
	jwt.RegisteredClaims
	IsAdmin bool   `json:"is_admin"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	UserId  uint   `json:"user_id"`
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

// Password Hashing and Verification
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func checkHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *userService) RegisterUser(name, email, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		IsAdmin:  false,
	}

	return s.userRepo.CreateUser(user)
}

func (s *userService) LoginUser(email, password string) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !checkHashPassword(password, user.Password) {
		return "", "", errors.New("invalid password")
	}

	accessToken, err := s.generateToken(user, time.Hour*24)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateToken(user, time.Hour*24)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) generateToken(user *models.User, expiry time.Duration) (string, error) {
	now := time.Now()

	claims := JWTCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.New().String(),
		},
		IsAdmin: user.IsAdmin,
		Name:    user.Name,
		Email:   user.Email,
		UserId:  user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func (s *userService) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userIdString, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid user ID in claims")
	}
	userId, err := strconv.ParseUint(userIdString, 10, 64) // Parse string to uint
	if err != nil {
		return "", errors.New("failed to convert user ID to uint")
	}

	// Get the user by ID
	// userId := uint(claims["sub"].(float64))
	user, err := s.userRepo.GetUserById(uint(userId))
	if err != nil {
		return "", errors.New("user not found")
	}

	return s.generateToken(user, time.Minute*15)
}

func (s *userService) VerifyToken(tokenString string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTCustomClaims)

	if !ok {
		return nil, errors.New("unauthorized: invalid token claims")
	}
	if !token.Valid {
		return nil, errors.New("unauthorized: token not valid")
	}

	return claims, nil
}
