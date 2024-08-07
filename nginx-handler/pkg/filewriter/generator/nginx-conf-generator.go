package generator

import "fmt"

type FWDProxyDescriptor struct {
	ListenPort, RedirectPort int
	Name                     string
}

func (r *FWDProxyDescriptor) asConfigurationString() string {
	conf := "\tserver {\n"
	conf += fmt.Sprintf("\t\tlisten %d;\n", r.ListenPort)
	conf += fmt.Sprintf("\t\tproxy_pass host.docker.internal:%d; # %s\n", r.RedirectPort, r.Name)
	conf += "\t}\n"

	return conf
}

func GenerateNginxConf(proxies *[]FWDProxyDescriptor) string{
	conf := "worker_processes 1; \n\n"
	conf += "events {\n worker_connections 1024;\n}\n\n"
	conf += "stream {\n"

	for _, proxy := range *proxies {
		conf += proxy.asConfigurationString()
	}
	conf += "}\n"

	return conf
}

