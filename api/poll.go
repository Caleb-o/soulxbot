package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/soulxburn/soulxbot/db"
	"github.com/soulxburn/soulxbot/twitch"
)

type StreamPoller struct {
	db        *db.Database
	twitchAPI twitch.ITwitchAPI
}

func NewStreamPoller(db *db.Database, twitchAPI twitch.ITwitchAPI) StreamPoller {
	return StreamPoller{db, twitchAPI}
}

func (sp StreamPoller) goliveHandler(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	apiKey := params.Get("key")

	user, ok := sp.db.FindUserByApiKey(apiKey)
	stream := sp.db.FindCurrentStream(user.ID)

	if ok && stream == nil {
		log.Printf("%s is now live!", user.DisplayName)
		stream = sp.db.InsertStream(user.ID, time.Now())
		res.WriteHeader(http.StatusAccepted)
		go sp.PollStreamStatus(stream, user)
	} else {
		log.Println("Go live not authorized")
		res.WriteHeader(http.StatusUnauthorized)
	}
}

// PollStreamStatus
func (sp StreamPoller) PollStreamStatus(stream *db.Stream, streamUser *db.User) {
	tick := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-tick.C:
			streamInfo, err := sp.twitchAPI.GetStream(streamUser.Username)
			if err != nil {
				log.Println("Error fetching stream info: ", err)
				continue
			}

			if stream.TWID == nil || stream.Title == nil {
				twid, err := strconv.Atoi(streamInfo.ID)
				if err == nil {
					sp.db.UpdateStreamInfo(stream.ID, twid, streamInfo.Title)
					stream = sp.db.FindStreamById(stream.ID)
				}
			}

			if streamInfo == nil {
				sp.db.UpdateStreamEndedAt(stream.ID, time.Now())
				tick.Stop()
				return
			}
		}
	}
}

func (sp StreamPoller) RestartStreamStatusPolls() {
	// Restart any streams that were live when the bot was last shut down
	streamsInProgress := sp.db.FindAllCurrentStreams()
	for _, stream := range streamsInProgress {
		user, ok := sp.db.FindUserByID(stream.UserId)
		if ok {
			go sp.PollStreamStatus(&stream, user)
		}
	}
}
