package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/liamt01/course-scope-cmu/backend/cmd/api"
	_ "github.com/lib/pq"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	port string
	ip   string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	email struct {
		host         string
		port         string
		hostUser     string
		hostPassword string
		frontendLink string
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "API server port")
	flag.StringVar(&cfg.ip, "::", os.Getenv("IP"), "API server IP address")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	// Read the connection pool settings from command-line flags into the config struct.
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// Read the email settings from command-line flags into the config struct.
	flag.StringVar(&cfg.email.host, "email-host", os.Getenv("EMAIL_HOST"), "SMTP server hostname")
	flag.StringVar(&cfg.email.port, "email-port", os.Getenv("EMAIL_PORT"), "SMTP server port")
	flag.StringVar(&cfg.email.hostUser, "email-host-user", os.Getenv("EMAIL_HOST_USER"), "SMTP server user")
	flag.StringVar(&cfg.email.hostPassword, "email-host-password", os.Getenv("EMAIL_HOST_PASSWORD"), "SMTP server password")
	flag.StringVar(&cfg.email.frontendLink, "email-frontend-link", os.Getenv("EMAIL_FRONTEND_LINK"), "Frontend link")

	flag.Parse()

	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	e.HTTPErrorHandler = api.CustomHTTPErrorHandler

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Secure())
	e.Use(middleware.CORS())

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 10 * time.Second,
	}))

	// Database connection
	db, err := openDB(cfg)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			e.Logger.Fatal(err)
		}
	}(db)

	h := api.NewHandler(db, cfg.email.host, cfg.email.port, cfg.email.hostUser, cfg.email.hostPassword,
		fmt.Sprintf("CourseScope CMU <%s>", cfg.email.hostUser), cfg.email.frontendLink)

	e.Use(echoprometheus.NewMiddleware("metrics"))
	e.GET("/metrics", h.RequireAdmin(echoprometheus.NewHandler()))

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(40)))
	e.Use(middleware.BodyLimit("2M"))

	e.Use(h.Authenticate)

	e.GET("/courses", h.ListCourses)
	e.GET("/courses/:id", h.RetrieveCourse)
	e.POST("/courses", h.RequireAdmin(h.CreateCourse))
	e.PUT("/courses/:id", h.RequireAdmin(h.UpdateCourse))
	e.DELETE("/courses/:id", h.RequireAdmin(h.DeleteCourse))

	e.GET("/instructors", h.ListInstructors)
	e.GET("/instructors/:id", h.RetrieveInstructor)
	e.POST("/instructors", h.RequireAdmin(h.CreateInstructor))
	e.PUT("/instructors/:id", h.RequireAdmin(h.UpdateInstructor))
	e.DELETE("/instructors/:id", h.RequireAdmin(h.DeleteInstructor))

	e.GET("/offerings", h.ListOfferings)
	e.GET("/offerings/:id", h.RetrieveOffering)
	e.POST("/offerings", h.RequireAdmin(h.CreateOffering))
	e.PUT("/offerings/:id", h.RequireAdmin(h.UpdateOffering))
	e.DELETE("/offerings/:id", h.RequireAdmin(h.DeleteOffering))

	e.GET("/ratings", h.ListRatings)
	e.GET("/ratings/my", h.RequireAuth(h.ListMyRatings))
	e.GET("/ratings/my-likes", h.RequireAuth(h.ListMyLikedRatings))
	e.GET("/ratings/my-dislikes", h.RequireAuth(h.ListMyDislikedRatings))
	e.GET("/ratings/:id", h.RetrieveRating)
	e.GET("ratings/stats", h.CalculateRatingStats)
	e.POST("/ratings", h.RequireActivatedUser(h.CreateRating))
	e.POST("/ratings/:id/like", h.RequireActivatedUser(h.LikeRating))
	e.DELETE("/ratings/:id/like", h.RequireActivatedUser(h.UndoLikeRating))
	e.POST("/ratings/:id/dislike", h.RequireActivatedUser(h.DislikeRating))
	e.DELETE("/ratings/:id/dislike", h.RequireActivatedUser(h.UndoDislikeRating))
	e.PUT("/ratings/:id", h.RequireActivatedUser(h.UpdateRating))
	e.DELETE("/ratings/:id", h.RequireActivatedUser(h.DeleteRating))

	e.GET("/users", h.ListUsers)
	e.GET("/users/:id", h.RetrieveUser)
	e.GET("/users/me", h.RequireAuth(h.RetrieveUserMe))
	e.GET("/users/stats/me", h.RequireAuth(h.CalculateUserStats))
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:id", h.RequireAdmin(h.UpdateUser))
	e.PUT("/users/activated/me", h.ActivateUser)
	e.PUT("/users/password/me", h.UpdateUserPassword)
	e.PUT("/users/username/me", h.RequireActivatedUser(h.UpdateUsername))
	e.DELETE("/users/me", h.RequireAuth(h.DeleteUser))

	e.POST("/tokens/auth", h.CreateAuthToken)
	e.POST("/tokens/activation", h.RequireAuth(h.SendActToken))
	e.POST("/tokens/password-reset", h.SendPwdResetToken)
	e.DELETE("/tokens/my-expired", h.DeleteMyExpiredTokens)

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(net.JoinHostPort(cfg.ip, cfg.port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	// Set the maximum idle timeout.
	db.SetConnMaxIdleTime(duration)

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5-second deadline, then this will return an error.
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}
