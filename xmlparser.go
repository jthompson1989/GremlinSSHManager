package main

import (
	"encoding/xml"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type Server struct {
	Name     string `xml:"name"`
	Host     string `xml:"host"`
	Port     int    `xml:"port"`
	UserName string `xml:"username"`
	AuthType string `xml:"auth"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Server  []Server `xml:"server"`
}

func getXMLPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "servers.xml"), nil
}

func GetServers() (Servers, error) {
	var servers Servers

	xmlPath, err := getXMLPath()
	if err != nil {
		return servers, err
	}

	xmlFile, err := os.Open(xmlPath)

	if err != nil {
		return servers, err
	}

	//Close closes the File, rendering it unusable for I/O.
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)

	if err != nil {
		return servers, err
	}

	xml.Unmarshal(byteValue, &servers)

	return servers, nil
}

func SaveServersToXml(servers Servers) error {
	xmlData, err := xml.MarshalIndent(servers, "", "  ")
	if err != nil {
		return err
	}

	xmlPath, err := getXMLPath()
	if err != nil {
		return err
	}

	// Create/truncate file for writing (0644 permissions)
	err = os.WriteFile(xmlPath, xmlData, 0644)
	return err
}

func GetSavedServerByName(serverName string, servers Servers) Server {
	for _, value := range servers.Server {
		if value.Name == serverName {
			return value
		}
	}
	return Server{}
}

func DeleteServer(serverName string) error {
	var servers, err = GetServers()

	if err != nil {
		return err
	}

	var server = GetSavedServerByName(serverName, servers)
	if server.Name == "" {
		return nil
	}

	var filteredServers []Server
	for _, s := range servers.Server {
		if s.Name != serverName {
			filteredServers = append(filteredServers, s)
		}
	}

	servers.Server = filteredServers
	return SaveServersToXml(servers)
}

func AddServer(server Server) error {
	var servers, err = GetServers()

	if err != nil {
		return err
	}

	existingServer := GetSavedServerByName(server.Name, servers)
	if existingServer.Name != "" {
		return errors.New("Server Already Exists in XML File")
	}

	servers.Server = append(servers.Server, server)

	err = SaveServersToXml(servers)

	if err != nil {
		return err
	}

	return nil
}
