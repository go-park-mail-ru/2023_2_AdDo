package microservices_init

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"strconv"
)

func RegisterInConsul(port int, serviceName, serviceHost string, logger *logrus.Logger) string {
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		logger.Fatalf("error while creating consul client %s", err.Error())
	}

	serviceID := "API_" + serviceHost + strconv.Itoa(port)

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    port,
		Address: serviceHost,
	})
	if err != nil {
		logger.Fatalln("cant add service to consul", err)
	}

	logger.Infoln("registered in consul", serviceID)

	return serviceID
}

func UnRegisterInConsul(serviceId string, logger *logrus.Logger) {
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		logger.Fatalf("error while creating consul client %s", err.Error())
	}

	err = consul.Agent().ServiceDeregister(serviceId)
	if err != nil {
		logger.Fatalln("cant delete service from consul", err)
	}
	logger.Fatalf("sevice deleted from consul %s", serviceId)
}
