package bvr

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/angelfluffyookami/247BVR/modules/common/global"
)

func GetUserByID(ID string) (UserStruct, error) {

	hadError := true

	var req *http.Response

	retryCount := 0

	var err error

	for hadError {

		req, err = http.Get(global.Config.APIEndpoint + "users/" + ID)

		if err != nil {

			if retryCount == 4 {

				time.Sleep(time.Duration(int64(retryCount*30)) * time.Second)
				logging.Err().Alert().Message("HTTP GET request error retryCount exceeded for: /user/<id>. Is server up? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.").Add()
				return UserStruct{}, err
			}

			hadError = true

			retryCount += 1
		} else {
			hadError = false
		}

	}

	var user UserStruct

	err = json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		logging.Err().Alert().Message("JSON Unmarshal error retryCount exceeded for: /online. Are API definitions up to date? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.")

		hadError = true
		return UserStruct{}, err
	}

	return user, nil
}

func GetKillsByUserID(killerID string) ([]KillStruct, error) {

	hadError := true

	var req *http.Response

	retryCount := 0

	var err error

	for hadError {

		req, err = http.Get(global.Config.APIEndpoint + "kills?killerId=" + killerID)

		if err != nil {

			if retryCount == 4 {

				time.Sleep(time.Duration(int64(retryCount*30)) * time.Second)
				logging.Err().Alert().Message("HTTP GET request error retryCount exceeded for: /user/<id>. Is server up? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.").Add()
				return []KillStruct{}, err
			}

			hadError = true

			retryCount += 1
		} else {
			hadError = false
		}

	}

	var kills []KillStruct

	err = json.NewDecoder(req.Body).Decode(&kills)

	if err != nil {
		logging.Err().Alert().Message("JSON Unmarshal error retryCount exceeded for: /online. Are API definitions up to date? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.")

		hadError = true
		return []KillStruct{}, err
	}

	return kills, nil
}

func GetDeathsByUserID(victimID string) ([]DeathStruct, error) {

	hadError := true

	var req *http.Response

	retryCount := 0

	var err error

	for hadError {

		req, err = http.Get(global.Config.APIEndpoint + "deaths?victimId=" + victimID)

		if err != nil {

			if retryCount == 4 {

				time.Sleep(time.Duration(int64(retryCount*30)) * time.Second)
				logging.Err().Alert().Message("HTTP GET request error retryCount exceeded for: /user/<id>. Is server up? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.").Add()
				return []DeathStruct{}, err
			}

			hadError = true

			retryCount += 1
		} else {
			hadError = false
		}

	}

	var deaths []DeathStruct

	err = json.NewDecoder(req.Body).Decode(&deaths)

	if err != nil {
		logging.Err().Alert().Message("JSON Unmarshal error retryCount exceeded for: /online. Are API definitions up to date? Cache refresh paused until heartbeat detected. Further charts will be generated with last server snapshot.")

		hadError = true
		return []DeathStruct{}, err
	}

	return deaths, nil
}
