package api

import (
	"github.com/gofiber/fiber/v2"
	"go-hotel-reservation-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HotelHandler handles HTTP requests related to hotels.
type HotelHandler struct {
	store *db.Store
}

// NewHotelHandler creates a new instance of HotelHandler.
func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

// HandleGetRooms handles the retrieval of rooms for a specific hotel.
func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

// HandleGetHotel handles the retrieval of a specific hotel.
func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

// ResourceResp represents the response structure for resource endpoints.
type ResourceResp struct {
	Results int         `json:"results"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
}

// HotelQueryParams represents the query parameters for hotel endpoints.
type HotelQueryParams struct {
	db.Pagination
	Rating int `json:"rating"`
}

// HandleGetHotels handles the retrieval of hotels based on query parameters.
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{
		"rating": params.Rating,
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrNotResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(params.Page),
	}
	return c.JSON(resp)
}
