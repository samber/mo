package mo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo_Success(t *testing.T) {
	is := assert.New(t)

	result := Do(func() string {
		return "Hello, World!"
	})

	is.False(result.IsError())
	is.Equal("Hello, World!", result.MustGet())
}

func TestDo_Error(t *testing.T) {
	is := assert.New(t)

	result := Do(func() string {
		panic(errors.New("something went wrong"))
	})

	is.True(result.IsError())
	is.EqualError(result.Error(), "something went wrong")
}

func TestDo_ComplexSuccess(t *testing.T) {
	is := assert.New(t)

	validateBooking := func(params map[string]string) Result[map[string]string] {
		if params["guest"] != "" && params["roomType"] != "" {
			return Ok(params)
		}
		return Err[map[string]string](errors.New("validation failed"))
	}

	createBooking := func(guest string) Result[string] {
		if guest != "" {
			return Ok("Booking Created for: " + guest)
		}
		return Err[string](errors.New("booking creation failed"))
	}

	assignRoom := func(booking string, roomType string) Result[string] {
		if roomType != "" {
			return Ok("Room Assigned: " + roomType + " for " + booking)
		}
		return Err[string](errors.New("room assignment failed"))
	}

	bookRoom := func(params map[string]string) Result[[]string] {
		return Do(func() []string {
			values := validateBooking(params).MustGet()
			booking := createBooking(values["guest"]).MustGet()
			room := assignRoom(booking, values["roomType"]).MustGet()
			return []string{booking, room}
		})
	}

	params := map[string]string{
		"guest":    "Foo Bar",
		"roomType": "Suite",
	}

	result := bookRoom(params)
	is.False(result.IsError())
	is.Equal([]string{"Booking Created for: Foo Bar", "Room Assigned: Suite for Booking Created for: Foo Bar"}, result.MustGet())
}

func TestDo_ComplexError(t *testing.T) {
	is := assert.New(t)

	validateBooking := func(params map[string]string) Result[map[string]string] {
		if params["guest"] != "" && params["roomType"] != "" {
			return Ok(params)
		}
		return Err[map[string]string](errors.New("validation failed"))
	}

	createBooking := func(guest string) Result[string] {
		if guest != "" {
			return Ok("Booking Created for: " + guest)
		}
		return Err[string](errors.New("booking creation failed"))
	}

	assignRoom := func(booking string, roomType string) Result[string] {
		if roomType != "" {
			return Ok("Room Assigned: " + roomType + " for " + booking)
		}
		return Err[string](errors.New("room assignment failed"))
	}

	bookRoom := func(params map[string]string) Result[[]string] {
		return Do(func() []string {
			values := validateBooking(params).MustGet()
			booking := createBooking(values["guest"]).MustGet()
			room := assignRoom(booking, values["roomType"]).MustGet()
			return []string{booking, room}
		})
	}

	params := map[string]string{
		"guest":    "",
		"roomType": "Suite",
	}

	result := bookRoom(params)
	is.True(result.IsError())
	is.EqualError(result.Error(), "validation failed")
}
