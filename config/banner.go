package config

import (
	"fmt"
)

func BannerInit() {
	blue := "\033[34m"
	fmt.Println(blue + `
	██████╗ ██████╗  ██████╗ ███╗   ██╗██╗   ██╗ █████╗ ██████╗  ██████╗ ████████╗
	██╔══██╗██╔══██╗██╔═══██╗████╗  ██║╚██╗ ██╔╝██╔══██╗██╔══██╗██╔═══██╗╚══██╔══╝
	██████╔╝██████╔╝██║   ██║██╔██╗ ██║ ╚████╔╝ ███████║██████╔╝██║   ██║   ██║   
	██╔══██╗██╔══██╗██║   ██║██║╚██╗██║  ╚██╔╝  ██╔══██║██╔══██╗██║   ██║   ██║   
	██████╔╝██║  ██║╚██████╔╝██║ ╚████║   ██║   ██║  ██║██████╔╝╚██████╔╝   ██║   
	╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═════╝  ╚═════╝    ╚═╝   
                                                                              
		version 0.0.1-beta2

`)
}
