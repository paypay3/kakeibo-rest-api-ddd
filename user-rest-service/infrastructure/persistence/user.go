package persistence

import (
	"database/sql"

	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/domain/vo"

	"golang.org/x/xerrors"

	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/apierrors"
	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/domain/userdomain"
	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/infrastructure/persistence/datasource"
	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/infrastructure/persistence/db"
)

type userRepository struct {
	*db.RedisHandler
	*db.MySQLHandler
}

func NewUserRepository(redisHandler *db.RedisHandler, mysqlHandler *db.MySQLHandler) *userRepository {
	return &userRepository{
		redisHandler,
		mysqlHandler,
	}
}

func (r *userRepository) FindSignUpUserByUserID(userID userdomain.UserID) (*userdomain.SignUpUser, error) {
	query := `
        SELECT
            user_id,
            name,
            email,
            password
        FROM
            users
        WHERE
            user_id = ?`

	var signUpUserDto datasource.SignUpUser
	if err := r.MySQLHandler.Conn.QueryRowx(query, userID).StructScan(&signUpUserDto); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, apierrors.ErrUserNotFound
		}

		return nil, err
	}

	var userValidationError apierrors.UserValidationError

	userIDVo, err := userdomain.NewUserID(signUpUserDto.UserID)
	if err != nil {
		userValidationError.UserID = "ユーザーIDが正しくありません"
	}

	nameVo, err := userdomain.NewName(signUpUserDto.Name)
	if err != nil {
		userValidationError.Name = "名前が正しくありません"
	}

	emailVo, err := vo.NewEmail(signUpUserDto.Email)
	if err != nil {
		userValidationError.Email = "メールアドレスが正しくありません"
	}

	if userValidationError.UserID != "" ||
		userValidationError.Name != "" ||
		userValidationError.Email != "" {
		return nil, apierrors.NewBadRequestError(&userValidationError)
	}

	signUpUser := userdomain.NewSignUpUserFromDataSource(userIDVo, nameVo, emailVo)

	return signUpUser, nil
}

func (r *userRepository) FindSignUpUserByEmail(email vo.Email) (*userdomain.SignUpUser, error) {
	query := `
        SELECT
            user_id,
            name,
            email,
            password
        FROM
            users
        WHERE
            email = ?`

	var signUpUserDto datasource.SignUpUser
	if err := r.MySQLHandler.Conn.QueryRowx(query, email).StructScan(&signUpUserDto); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, apierrors.ErrUserNotFound
		}

		return nil, err
	}

	var userValidationError apierrors.UserValidationError

	userIDVo, err := userdomain.NewUserID(signUpUserDto.UserID)
	if err != nil {
		userValidationError.UserID = "ユーザーIDが正しくありません"
	}

	nameVo, err := userdomain.NewName(signUpUserDto.Name)
	if err != nil {
		userValidationError.Name = "名前が正しくありません"
	}

	emailVo, err := vo.NewEmail(signUpUserDto.Email)
	if err != nil {
		userValidationError.Email = "メールアドレスが正しくありません"
	}

	if userValidationError.UserID != "" ||
		userValidationError.Name != "" ||
		userValidationError.Email != "" {
		return nil, apierrors.NewBadRequestError(&userValidationError)
	}

	signUpUser := userdomain.NewSignUpUserFromDataSource(userIDVo, nameVo, emailVo)

	return signUpUser, nil
}

func (r *userRepository) CreateSignUpUser(signUpUser *userdomain.SignUpUser) error {
	query := `
        INSERT INTO users
            (user_id, name, email, password)
        VALUES
            (?,?,?,?)`

	if _, err := r.MySQLHandler.Conn.Exec(query, signUpUser.UserID(), signUpUser.Name(), signUpUser.Email(), signUpUser.Password()); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteSignUpUser(signUpUser *userdomain.SignUpUser) error {
	query := `
        DELETE
        FROM
            users
        WHERE
            user_id = ?`

	_, err := r.MySQLHandler.Conn.Exec(query, signUpUser.UserID())

	return err
}
