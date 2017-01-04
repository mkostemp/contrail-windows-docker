package hns

import (
	"encoding/json"
	"errors"

	"github.com/Microsoft/hcsshim"
	log "github.com/Sirupsen/logrus"
)

func CreateHNSNetwork(configuration *hcsshim.HNSNetwork) (string, error) {
	log.Infoln("Creating HNS network")
	configBytes, err := json.Marshal(configuration)
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	log.Debugln("Config:", string(configBytes))
	response, err := hcsshim.HNSNetworkRequest("POST", "", string(configBytes))
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	return response.Id, nil
}

func DeleteHNSNetwork(hnsID string) error {
	log.Infoln("Deleting HNS network", hnsID)
	_, err := hcsshim.HNSNetworkRequest("DELETE", hnsID, "")
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func ListHNSNetworks() ([]hcsshim.HNSNetwork, error) {
	log.Infoln("Listing HNS networks")
	nets, err := hcsshim.HNSListNetworkRequest("GET", "", "")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return nets, nil
}

func GetHNSNetwork(hnsID string) (*hcsshim.HNSNetwork, error) {
	log.Infoln("Getting HNS network", hnsID)
	net, err := hcsshim.HNSNetworkRequest("GET", hnsID, "")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return net, nil
}

func GetHNSNetworkByName(name string) (*hcsshim.HNSNetwork, error) {
	log.Infoln("Getting HNS network by name:", name)
	nets, err := hcsshim.HNSListNetworkRequest("GET", "", "")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	for _, n := range nets {
		if n.Name == name {
			return &n, nil
		}
	}
	return nil, errors.New("HNS network not found by name")
}

func CreateHNSEndpoint(configuration *hcsshim.HNSEndpoint) (string, error) {
	log.Infoln("Creating HNS endpoint")
	configBytes, err := json.Marshal(configuration)
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	log.Debugln("Config: ", string(configBytes))
	response, err := hcsshim.HNSEndpointRequest("POST", "", string(configBytes))
	if err != nil {
		return "", err
	}
	return response.Id, nil
}

func DeleteHNSEndpoint(endpointID string) error {
	log.Infoln("Deleting HNS endpoint", endpointID)
	_, err := hcsshim.HNSEndpointRequest("DELETE", endpointID, "")
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func GetHNSEndpoint(endpointID string) (*hcsshim.HNSEndpoint, error) {
	log.Infoln("Getting HNS endpoint", endpointID)
	endpoint, err := hcsshim.HNSEndpointRequest("GET", endpointID, "")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return endpoint, nil
}
