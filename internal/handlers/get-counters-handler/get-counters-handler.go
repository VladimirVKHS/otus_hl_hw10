package get_counters_handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"otus_sn_counters/internal/constants"
	httpHelper "otus_sn_counters/internal/helpers/http"
	"otus_sn_counters/internal/logger"
	"otus_sn_counters/internal/models/counter"
	"strconv"
)

func GetCountersHandler(w http.ResponseWriter, r *http.Request) {
	requestId := r.Context().Value(constants.RequestIDKey).(string)
	logger.Info(fmt.Sprintf("Get counters. Request ID = %s", requestId))

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	c := counter.Counter{}

	err := counter.GetCounter(r.Context(), id, &c)
	if err != nil {
		c.UserId = id
	}

	httpHelper.JsonResponse(w, c.ToResponse())
}
