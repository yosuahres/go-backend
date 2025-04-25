package middleware

import (
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/yosuahres/go-backend/initializers"
    "github.com/yosuahres/go-backend/models"
)

func RequireAuth(c *gin.Context) {
    // Get the cookie from the request
    tokenString, err := c.Cookie("Authorization")
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Decode and validate the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Ensure the signing method is HMAC
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(os.Getenv("SECRET")), nil
    })

    if err != nil || !token.Valid {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Extract claims and validate them
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Check the expiration time
    if float64(time.Now().Unix()) > claims["exp"].(float64) {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Find the user with the token's subject (sub)
    var user models.User
    initializers.DB.First(&user, claims["sub"])
    if user.ID == 0 {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Attach the user to the request context
    c.Set("user", user)

    // Continue to the next middleware or handler
    c.Next()
}	