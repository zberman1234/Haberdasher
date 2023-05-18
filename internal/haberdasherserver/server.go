package haberdasherserver

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	//	Local import statements follow the format:
	//	"<relative module path>/<package>"
	pb "github.com/example/rpc/haberdasher"

	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (*pb.Hat, error) {

	if timeHours() > 13 && timeHours() < 14 {
		httpStatus := twirp.ServerHTTPStatusFromErrorCode(twirp.Unavailable)
		return nil, twirp.Unavailable.Error("Error " + strconv.Itoa(httpStatus) + " " + http.StatusText(httpStatus) + "\n Haberdasher is out for lunch\n")
	}
	// if timeHours() > 13 && timeHours() < 14 {
	// 	return nil, twirp.NewError(twirp.Unavailable, "Server is down for lunch")
	// }
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(5)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(4)],
	}, nil
}

// returns the current time in hours, as a float64
// where 1:30PM = 13.5
func timeHours() float64 {
	currentTime := time.Now()

	hour := float64(currentTime.Hour())
	minute := float64(currentTime.Minute())
	proportion := hour + minute/60.0

	return proportion
}
