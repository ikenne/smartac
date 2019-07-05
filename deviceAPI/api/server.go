package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"

	"deviceAPI/lib"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "github.com/rs/cors"
)

var (
	status map[string]struct{}
)

const (
	co2Limit    = 9
	statusLimit = 150
	sensorLimit = 500
)

type APIServer struct {
	db     lib.DeviceDB
	http   *http.Server
	Router *mux.Router
}

func NewServer(port string, db lib.DeviceDB) *APIServer {
	s := &APIServer{
		db:     db,
		http:   &http.Server{Addr: ":" + port},
		Router: mux.NewRouter(),
	}

	allowedOrig := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type"})
	allowedMethods := handlers.AllowedMethods([]string{"POST, GET, OPTIONS, PUT, DELETE"})

	handler := handlers.CORS(allowedOrig, allowedHeaders, allowedMethods)(s.Router)
	s.http.Handler = handler

	s.registerHandlers()

	buildHealthStatus()

	return s
}

func buildHealthStatus() {
	status = make(map[string]struct{})
	unhealthy := []string{"needs_service", "needs_new_filter", "gas_leak"}

	for _, v := range unhealthy {
		status[v] = struct{}{}
	}
}

func (s *APIServer) Start() {
	ln, err := net.Listen("tcp", s.http.Addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error starting server %s", err))
	}

	go func() {
		err := s.http.Serve(ln)
		if err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("Server error %s", err))
		}
	}()

	fmt.Println("server started")
}

func (s *APIServer) Stop() {
	err := s.http.Shutdown(context.Background())
	if err != nil {
		log.Fatal("Shut down failed")
	}
}

func (s *APIServer) registerHandlers() {
	get, post, put, delete := http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete
	s.Router.Path("/device").Methods(get).HandlerFunc(s.getAllDevices)
	s.Router.Path("/device/service").Methods(get).HandlerFunc(s.getUnHealthyDevices)
	s.Router.Path("/device/co2limit").Methods(get).HandlerFunc(s.getUnHealthyCO2Devices)

	s.Router.Path("/device").Methods(post).HandlerFunc(s.createDevice)
	s.Router.Path("/device/{id}").Methods(delete).HandlerFunc(s.deleteDevice)
	s.Router.Path("/device/{id}").Methods(put).HandlerFunc(s.updateDevice)
	s.Router.Path("/device/{id}").Methods(get).HandlerFunc(s.getDevice)

}

func (s *APIServer) getAllDevices(w http.ResponseWriter, r *http.Request) {
	s.responseJSON(w, http.StatusOK, s.db.GetAllDevices())
}

func (s *APIServer) getDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	a, err := s.db.GetDevice(key)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.responseJSON(w, http.StatusOK, a)
}

func (s *APIServer) createDevice(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var d lib.Device
	err := json.Unmarshal(reqBody, &d)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateDevice(d)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.db.AddDevice(d)
	s.responseJSON(w, http.StatusCreated, d)
}

func validateDevice(d lib.Device) (err error) {
	var sb strings.Builder

	if reflect.DeepEqual(d, lib.Device{}) {
		sb.WriteString("No values in device input")
		err = errors.New(sb.String())
		return
	}

	if d.SerialNumber == "" {
		sb.WriteString("No serial number")
	}

	if d.RegistrationDate.IsZero() {
		sb.WriteString("No registeration date")
	}

	if d.FirmwareVersion == "" {
		sb.WriteString("No firmware")
	}

	if len(d.Status) >= statusLimit {
		sb.WriteString(fmt.Sprintf("Status longer than limit %v", statusLimit))
	}

	if sb.Len() > 0 {
		err = errors.New(sb.String())
	}

	return err
}

func (s *APIServer) updateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var d lib.Device
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &d)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())

		return
	}

	err = validateUpdate(d)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	ua, err := s.db.UpdateDevice(key, d)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())

		return
	}

	s.responseJSON(w, http.StatusOK, ua)
}

func validateUpdate(d lib.Device) (err error) {
	var sb strings.Builder

	if reflect.DeepEqual(d, lib.Device{}) {
		sb.WriteString("No values in device input")
		err = errors.New(sb.String())
		return
	}

	if len(d.Temperature) > sensorLimit {
		sb.WriteString(fmt.Sprintf("Temperature sensor values more than limit %v", sensorLimit))
	}

	if len(d.CO2) > sensorLimit {
		sb.WriteString(fmt.Sprintf("CO2 sensor values  more than limit %v", sensorLimit))
	}

	if len(d.Humidity) > sensorLimit {
		sb.WriteString(fmt.Sprintf("Humidity sensor values more than limit %v", sensorLimit))
	}

	if sb.Len() > 0 {
		err = errors.New(sb.String())
	}

	return err
}

func (s *APIServer) deleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	err := s.db.DeleteDevice(key)
	if err != nil {
		s.responsError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.responseJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) getUnHealthyDevices(w http.ResponseWriter, r *http.Request) {
	s.responseJSON(w, http.StatusOK, s.db.GetUnHealthyDevices(status))
}

func (s *APIServer) getUnHealthyCO2Devices(w http.ResponseWriter, r *http.Request) {
	s.responseJSON(w, http.StatusOK, s.db.GetCO2LimitDevice(co2Limit))
}
