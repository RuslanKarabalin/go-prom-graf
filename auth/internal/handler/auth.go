package handler

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	db *pgxpool.Pool
}

func New(pgPool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		db: pgPool,
	}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	ctx := context.Background()
	authUser := new(AuthUser)

	if err := c.Bind().Body(authUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tx, err := h.db.Begin(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	insertUser := `insert into users(email, password_hash) values($1, $2) returning id`
	insertRole := `insert into user_roles(user_id, role) values($1, 'user')`

	var id uuid.UUID

	err = tx.QueryRow(ctx, insertUser, authUser.Email, hash).Scan(&id)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == "23505" {
			return c.Status(fiber.StatusBadRequest).SendString("User already exists")
		}
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	_, err = tx.Exec(ctx, insertRole, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).SendString(id.String())
}

// func Login(c *fiber.Ctx) error {

// }

// func Refresh(c *fiber.Ctx) error {

// }
