package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DisposableChat/api-core"
	"github.com/gofiber/fiber/v2"
)

func GetSession(c *fiber.Ctx) (string, interface{}, error) {
	var data core.GetSessionRequest
	err := c.BodyParser(&data)
	if err != nil {
		return "", nil, errors.New(core.BodyError)
	}

	if len(data.Token) < 1 {
		return "", nil, errors.New(core.InvalidTokenError)
	}

	password := core.GetHashPassword(core.LocalHashPassword, data.HashPassword)
	decrypt, err := core.Decrypt(password, data.Token)
	if err != nil {
		println(err.Error())
		return "", nil, errors.New(core.InternalServerError)
	}

	session, err := core.UnParseSession(decrypt)
	if err != nil {
		return "", nil, errors.New(core.InvalidTokenError)
	}

	if session.Expiration < time.Now().Unix() {
		return "", nil, errors.New(core.ExpiredTokenError)
	}

	return core.Authorized, session, nil
}

func CreateSession(c *fiber.Ctx) (string, interface{}, error) {
	var data core.CreateSessionRequest
	err := c.BodyParser(&data)
	if err != nil {
		return "", nil, errors.New(core.BodyError)
	}

	if len(data.PublicName) < 1 || len(data.PublicName) > 32 {
		return "", nil, errors.New(core.NameLengthError)
	}

	expirationDuration := time.Duration(data.Expiration * int64(time.Hour))
	if expirationDuration.Hours() > (24 * 7) || expirationDuration.Minutes() < 15 {
		return "", nil, errors.New(core.ExpirationSessionError)
	}

	id, err := core.RandomToken(fmt.Sprintf("%s.%d", data.PublicName, time.Now().UnixNano()))
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	session := core.Session{
		ID:         id,
		Expiration: time.Now().Add(expirationDuration).Unix(),
	}

	sessionString := core.ParseSession(session)
	password := core.GetHashPassword(core.LocalHashPassword, data.HashPassword)
	encrypt, err := core.Encrypt(password, sessionString)
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	user := core.User{
		ID:         id,
		PublicName: data.PublicName,
		Groups:    []string{},
		Expiration: time.Now().Add(expirationDuration).Unix(),
	}

	userString, err := core.ParseJSON(user)
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	err = Redis.Client.SetEX(context.Background(), fmt.Sprintf("user:%s", id), userString, expirationDuration).Err()
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	return core.Created, encrypt, nil
}