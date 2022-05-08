package update_counters_handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"otus_sn_counters/internal/constants"
	httpHelper "otus_sn_counters/internal/helpers/http"
	"otus_sn_counters/internal/logger"
	"otus_sn_counters/internal/models/counter"
	"strconv"
)

func UpdateCountersHandler(w http.ResponseWriter, r *http.Request) {
	requestId := r.Context().Value(constants.RequestIDKey).(string)
	logger.Info(fmt.Sprintf("Update counters. Request ID = %s", requestId))

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	var request counter.CounterUpdateRequest
	if err := json.Unmarshal(data, &request); err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	validationErrs := validate.Struct(request)
	if validationErrs != nil {
		httpHelper.ValidationErrorResponse(w, validationErrs.Error())
		return
	}

	if err := request.CheckAndApply(r.Context(), id); err != nil {
		logger.Error("Check and apply error: " + err.Error())
		httpHelper.InternalServerErrorResponse(w)
		return
	}

	c := counter.Counter{}
	if err := counter.GetCounter(r.Context(), id, &c); err != nil {
		logger.Error("Get counter error: " + err.Error())
		httpHelper.InternalServerErrorResponse(w)
		return
	}

	httpHelper.JsonResponse(w, c.ToResponse())
}
