package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	for {
		RenderIntro()
		fmt.Scanln(&input)

		switch input {
		case "1":
			DisplayServers()
		case "2":
			AddingServer()
		case "3":
			RemovingServer()
		case "4":
			fmt.Println("EXITING")
			return
		}
	}
}

func RenderIntro() {
	logo := `
╔═════════════════════════════════════════════════════════════════╗
║     ░██████╗░██████╗░███████╗███╗░░░███╗██╗░░░░░██╗███╗░░██╗    ║
║     ██╔════╝░██╔══██╗██╔════╝████╗░████║██║░░░░░██║████╗░██║    ║
║     ██║░░██╗░██████╔╝█████╗░░██╔████╔██║██║░░░░░██║██╔██╗██║    ║
║     ██║░░░██╗██╔══██╗██╔══╝░░██║╚██╔╝██║██║░░░░░██║██║╚████║    ║
║     ╚██████╔╝██║░░██║███████╗██║░╚═╝░██║███████╗██║██║░╚███║    ║
║     ░╚═════╝░╚═╝░░╚═╝╚══════╝╚═╝░░░░░╚═╝╚══════╝╚═╝╚═╝░░╚══╝    ║
║                                                                 ║
║                  ░██████╗░██████╗██╗░░██╗                       ║
║                  ██╔════╝██╔════╝██║░░██║                       ║
║                  ╚█████╗░╚█████╗░███████║                       ║
║                  ░╚═══██╗░╚═══██╗██╔══██║                       ║
║                  ██████╔╝██████╔╝██║░░██║                       ║
║                  ╚═════╝░╚═════╝░╚═╝░░╚═╝                       ║
║                                                                 ║
║  ███╗░░░███╗░█████╗░███╗░░██╗░█████╗░░██████╗░███████╗██████╗   ║   
║  ████╗░████║██╔══██╗████╗░██║██╔══██╗██╔════╝░██╔════╝██╔══██   ║
║  ██╔████╔██║███████║██╔██╗██║███████║██║░░██╗░█████╗░░██████╔   ║ 
║  ██║╚██╔╝██║██╔══██║██║╚████║██╔══██║██║░░╚██╗██╔══╝░░██╔══██   ║
║  ██║░╚═╝░██║██║░░██║██║░╚███║██║░░██║╚██████╔╝███████╗██║░░██   ║
║  ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═╝░░╚═╝░╚═════╝░╚══════╝╚═╝░░╚═   ║ 
║                                                                 ║
╚═════════════════════════════════════════════════════════════════╝
`
	fmt.Println(logo)
	fmt.Println("Press Following Choice:")
	fmt.Println("1.) List Servers")
	fmt.Println("2.) Add Server")
	fmt.Println("3.) Remove Server")
	fmt.Println("4.) Quit")
	fmt.Print(">")
}

func AddingServer() {
	var server Server
	fmt.Println("----------------------------------")
	fmt.Print("Name:")
	fmt.Scanln(&server.Name)
	fmt.Print("Host:")
	fmt.Scanln(&server.Host)
	fmt.Print("Port:")
	fmt.Scanln(&server.Port)
	fmt.Print("UserName:")
	fmt.Scanln(&server.UserName)
	fmt.Print("AuthType(pwd/key):")
	fmt.Scanln(&server.AuthType)

	err := AddServer(server)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
	}
}

func RemovingServer() {
	var serverName string
	var confirmation string
	fmt.Println("----------------------------------")
	fmt.Print("Enter Server Name for Deletion:")
	fmt.Scanln(&serverName)
	fmt.Printf("Are you sure you want to delete %s:(Y/N)", serverName)
	fmt.Scanln(&confirmation)

	if strings.ToUpper(confirmation) == "Y" {
		err := DeleteServer(serverName)
		if err != nil {
			fmt.Printf("Error Occured while Deleteing Server %s \n", err.Error())
			fmt.Scanln()
			return
		} else {
			fmt.Printf("Successfully Deleted Server %s \n", serverName)
			fmt.Scanln()
		}

	} else {
		fmt.Println("Cancelling Deletion")
	}
}

func DisplayServers() {
	fmt.Println("----------------------------------")
	servers, err := GetServers()
	if err != nil {
		fmt.Println("Error getting servers:", err)
		fmt.Scanln()
		return
	}

	for index, value := range servers.Server {
		choice := index + 1
		fmt.Printf("%v.) %s \n", choice, value.Name)
	}
	fmt.Print(">")
	var input int
	fmt.Scanln(&input)

	selectedServer := servers.Server[input-1]

	var execError = ExeSSHPwd(selectedServer.Host, selectedServer.Port, selectedServer.UserName)
	if execError != nil {
		fmt.Printf("Error Occured while Executing Server %s \n", err.Error())
		fmt.Scanln()
	} else {
		fmt.Printf("Server Exec SSH Done")
		fmt.Scanln()
	}

}
