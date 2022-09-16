package main

import (
	"todoapps/application"
	"todoapps/shared/driver"
)

var Version = "0.0.1"

func main() {
	//appMap := map[string]func() driver.RegistryContract{
	//	"apptodo": application.NewApptodo(),
	//}
	//flag.Parse()
	//
	//app, exist := appMap[flag.Arg(0)]
	//if exist {
	//	fmt.Printf("Version %s", Version)
	//	driver.Run(app())
	//} else {
	//	fmt.Println("You may try 'go run main.go <app_name>' :")
	//	for appName := range appMap {
	//		fmt.Printf(" - %s\n", appName)
	//	}
	//}

	driver.Run(application.NewApptodo()())

}

//func openbrowser(url string) {
//	var err error
//
//	switch runtime.GOOS {
//	case "linux":
//		err = exec.Command("xdg-open", url).Start()
//	case "windows":
//		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
//	case "darwin":
//		err = exec.Command("open", url).Start()
//	default:
//		err = fmt.Errorf("unsupported platform")
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
