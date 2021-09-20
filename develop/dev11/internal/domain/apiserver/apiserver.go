package apiserver

import (
	"context"
	"encoding/json"
	"go-hw2/Development/task11/internal/domain/models"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"

	"github.com/google/uuid"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar"
)

type ProcessType int

const (
	DAY ProcessType = iota
	WEEK
	MONTH
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type ApiServer struct {
	log             logrus.FieldLogger
	calendarService calendar.Srv
	host            string
	server          *http.Server
}

func New(log logrus.FieldLogger, host string, calendarService calendar.Srv, addMiddlewares []func(http.Handler) http.Handler) *ApiServer {
	service := &ApiServer{
		log:             log,
		calendarService: calendarService,
		host:            host,
	}
	router := mux.NewRouter()
	router.Use(service.ContentTypeMiddleware)
	router.Use(service.LoggingMiddleware)

	for _, m := range addMiddlewares {
		router.Use(m)
	}

	// endpoints
	router.HandleFunc("/create_event", service.CreateEvent).Methods("POST")
	router.HandleFunc("/update_event/{id}", service.UpdateEvent).Methods("POST")
	router.HandleFunc("/delete_event/{id}", service.DeleteEventByID).Methods("POST")
	router.HandleFunc("/events_for_day/{id}", service.FindDayEvents).Methods("GET")
	router.HandleFunc("/events_for_week/{id}", service.FindWeekEvents).Methods("GET")
	router.HandleFunc("/events_for_month/{id}", service.FindMonthEvents).Methods("GET")

	server := http.Server{
		Addr:         host,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	service.server = &server
	return service
}

func (c *ApiServer) ListenAndServe() error {
	notifyCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-notifyCh
		c.log.Println("Stopping server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c.server.SetKeepAlivesEnabled(false)
		if err := c.server.Shutdown(ctx); err != nil {
			errorCh <- err
		}
		close(errorCh)
	}()

	c.log.Printf("Start server on %s....", c.host)
	if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	if err := <-errorCh; err != nil {
		return err
	}
	c.log.Println("Stop server successfully")
	return nil
}

func (c *ApiServer) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapperWriter := NewLoggingResponseWriter(w)
		next.ServeHTTP(wrapperWriter, r)
		latency := time.Since(start)
		c.log.Printf("%s %s %s %s %dms %d %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			latency.Microseconds(),
			wrapperWriter.statusCode,
			r.UserAgent())
	})
}

func (c *ApiServer) ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (c *ApiServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	eventReq := api.CreateEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&eventReq)
	if err != nil {
		c.log.Printf("failed to get body %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ev := unmarshalEvent(eventReq.Event)
	id, err := c.calendarService.CreateEvent(context.Background(), ev)
	if err != nil {
		c.log.Printf("failed to create event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&api.CreateEventResponse{
		ID: id.String(),
	})

	if err != nil {
		c.log.Printf("error in sending response : %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *ApiServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventReq := api.EventUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&eventReq)
	if err != nil {
		c.log.Printf("failed to get body %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rID := mux.Vars(r)["id"]
	eID, err := uuid.Parse(rID)
	if err != nil {
		c.log.Printf("failed to parse id %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ev := unmarshalEvent(eventReq.Event)
	err = c.calendarService.UpdateEvent(context.Background(), eID, ev)
	if err != nil {
		c.log.Printf("failed to update event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ApiServer) DeleteEventByID(w http.ResponseWriter, r *http.Request) {
	rID := mux.Vars(r)["id"]
	eID, err := uuid.Parse(rID)
	if err != nil {
		c.log.Printf("failed to parse id %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.calendarService.DeleteEvent(context.Background(), eID)
	if err != nil {
		c.log.Printf("failed to delete event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ApiServer) FindDayEvents(w http.ResponseWriter, r *http.Request) {
	c.processDaysRequest(w, r, DAY)
}

func (c *ApiServer) FindWeekEvents(w http.ResponseWriter, r *http.Request) {
	c.processDaysRequest(w, r, WEEK)
}

func (c *ApiServer) FindMonthEvents(w http.ResponseWriter, r *http.Request) {
	c.processDaysRequest(w, r, MONTH)
}

func (c *ApiServer) processDaysRequest(w http.ResponseWriter, r *http.Request, t ProcessType) {
	sDay := r.FormValue("from")
	var err error
	var events []*models.Event
	startDay, err := time.Parse("2006-01-02T15:04:05-0700", sDay)
	if err != nil {
		c.log.Printf("failed to parse from day %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch t {
	case DAY:
		events, err = c.calendarService.FindDayEvents(context.Background(), startDay)
	case WEEK:
		events, err = c.calendarService.FindWeekEvents(context.Background(), startDay)
	case MONTH:
		events, err = c.calendarService.FindMonthEvents(context.Background(), startDay)
	}
	if err != nil {
		c.log.Printf("failed to find events %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resEv := make([]*api.EventDTO, 0, len(events))
	for _, ev := range events {
		mEv := marshalEvent(ev)
		resEv = append(resEv, mEv)
	}
	err = json.NewEncoder(w).Encode(&api.FindMonthEventsResponse{
		Events: resEv,
	})
	if err != nil {
		c.log.Printf("failed to write response %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func unmarshalEvent(ev *models.EventDTO) *models.Event {
	return &models.Event{
		ID:          ev.ID,
		Date:        ev.Date,
		Description: ev.Description,
		OwnerID:     ev.OwnerID,
	}
}

func marshalEvent(mEv *models.Event) *models.EventDTO {
	return &api.EventDTO{
		ID:          mEv.ID,
		Date:        mEv.Date,
		Description: mEv.Description,
		OwnerID:     mEv.OwnerID,
	}
}
