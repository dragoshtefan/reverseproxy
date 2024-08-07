package generator

import "fmt"

func GenerateDockerYmlFile(filePath string, proxies *[]FWDProxyDescriptor) string {
	stringFormat := "services:\n" +
		"  nginx-proxy:\n" +
		"    image: nginx\n" +
		"    ports:\n" +
		"%s" +
		"    volumes:\n" +
		"      - ${NGINX_FILE}:/etc/nginx/nginx.conf"
	var portsMapped string = ""

	for _, proxy := range *proxies {
		portsMapped += fmt.Sprintf("      - %d:%d\n", proxy.ListenPort, proxy.ListenPort)
	}

	return fmt.Sprintf(stringFormat, portsMapped)
}
