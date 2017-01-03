package driver

import (
	"encoding/json"
	"errors"

	"github.com/Microsoft/hcsshim"
	"github.com/Sirupsen/logrus"
)

func CreateHNSNetwork(configuration *hcsshim.HNSNetwork) (string, error) {
	configBytes, err := json.Marshal(configuration)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	logrus.Debugln("[HNS] Request: POST, , ", string(configBytes))
	response, err := hcsshim.HNSNetworkRequest("POST", "", string(configBytes))
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Try restarting HNS: Restart-Service hns")
		return "", err
	}
	return response.Id, nil
}

func DeleteHNSNetwork(hnsID string) error {
	logrus.Debugln("[HNS] Request: DELETE, ", hnsID, ", ")
	logrus.Debugln(ListHNSNetworks())
	_, err := hcsshim.HNSNetworkRequest("DELETE", hnsID, "")
	logrus.Debugln(ListHNSNetworks())

	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Try restarting HNS: Restart-Service hns")
		return err
	}
	return nil
}

func ListHNSNetworks() ([]hcsshim.HNSNetwork, error) {
	logrus.Debugln("[HNS] Request: GET, , ")
	nets, err := hcsshim.HNSListNetworkRequest("GET", "", "")
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("[HNS] Try restarting HNS: Restart-Service hns")
		return nil, err
	}
	return nets, nil
}

func GetHNSNetwork(hnsID string) (*hcsshim.HNSNetwork, error) {
	logrus.Debugln("[HNS] Request: GET, ", hnsID, ", ")
	net, err := hcsshim.HNSNetworkRequest("GET", hnsID, "")
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Try restarting HNS: Restart-Service hns")
		return nil, err
	}
	return net, nil
}

func GetHNSNetworkByName(name string) (*hcsshim.HNSNetwork, error) {
	logrus.Debugln("[HNS] Request: GET, , ")
	nets, err := hcsshim.HNSListNetworkRequest("GET", "", "")
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Try restarting HNS: Restart-Service hns")
		return nil, err
	}
	for _, n := range nets {
		if n.Name == name {
			return &n, nil
		}
	}
	return nil, errors.New("[HNS] Network not found by name")
}
