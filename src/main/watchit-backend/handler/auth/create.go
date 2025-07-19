package auth

import (
	"net/http"
	"watchit/httpx/encryption"
	"watchit/httpx/infra/constants"
	"watchit/httpx/infra/store/postgres/models"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"

	"github.com/google/uuid"
)

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload *CreateUserPayload

	if err := httpx.HttpParse(r, &payload); err != nil {
		return httperr.BadRequest(err.Error())
	}

	if err := httpx.Validate.Struct(payload); err != nil {
		return httperr.BadRequest("not all fields are filled in")
	}

	exists, err := h.Store.Users.Get_UserProfileByUsername(ctx, payload.Username)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	if exists != nil {
		return httperr.BadRequest("this username is already registered")
	}

	tx, err := h.Db.BeginTx(ctx, nil)
	if err != nil {
		return httperr.Db(ctx, httperr.Err_DbNetwork)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	uuid := uuid.NewString()

	userCoreModel := &models.UserCore{
		UserUUID: uuid,
	}

	if userCoreErr := h.Store.Users.Create_UserCore(ctx, tx, userCoreModel); userCoreErr != nil {
		return httperr.Db(ctx, userCoreErr)
	}

	userProfileModel := &models.UserProfile{
		UserUUID: uuid,
		Username: payload.Username,
	}

	if userProfileErr := h.Store.Users.Create_UserProfile(ctx, tx, userProfileModel); userProfileErr != nil {
		return httperr.Db(ctx, userProfileErr)
	}

	userLimitModel := &models.UserLimit{
		UserUUID: uuid,
		LimitId:  constants.LimitFree,
	}

	if userLimitErr := h.Store.Users.Create_UserLimit(ctx, tx, userLimitModel); userLimitErr != nil {
		return httperr.Db(ctx, userLimitErr)
	}

	if err := tx.Commit(); err != nil {
		return httperr.Conflict("couldn't save data")
	}

	token, err := encryption.Encrypt(uuid)
	if err != nil {
		return httperr.InternalServerError(err.Error())
	}

	httpx.HttpResponse(w, r, http.StatusCreated, token)
	return nil
}
