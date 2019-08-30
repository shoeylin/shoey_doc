package main

import (
	"fmt"
	"shoey_doc/golang/ftp/models"
)

func main() {

	// pars := &models.FtpConnPars{
	// 	FtpAddress:     "127.0.0.1",
	// 	FtpPort:        21,
	// 	Username:       "shoey",
	// 	Password:       "12345678",
	// 	Homefolder:     "/Data/",
	// 	DebugFtp:       false,
	// 	RemoteFileName: "fileName.zip",
	// 	SaveFileName:   "./myFileName.zip"}
	// models.NewDownloadFtp(pars)

	// models.DeCompress("myFileName.zip")

	mdAry := models.Opnefile("myFileName.txt")

	fmt.Println("mdAry:", mdAry)
}
