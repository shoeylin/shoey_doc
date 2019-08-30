package models

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shenshouer/ftp4go"
)

type FtpConnPars struct {
	FtpAddress     string
	FtpPort        int
	Username       string
	Password       string
	Homefolder     string
	DebugFtp       bool
	RemoteFileName string
	SaveFileName   string
}

//ftp
func NewDownloadFtp(pars *FtpConnPars) {
	var logl int
	if pars.DebugFtp {
		logl = 1
	}

	ftpClient := ftp4go.NewFTP(logl)
	defer ftpClient.Quit()

	ftpClient.SetPassive(true)
	// ftpClient.Connect(pars.ftpAddress, pars.ftpPort, "")

	// connect
	_, err := ftpClient.Connect(pars.FtpAddress, pars.FtpPort, "")
	if err != nil {
		fmt.Printf("The FTP connection could not be established, error: %v", err.Error())
	}

	fmt.Printf("Connecting with username: %s and password %s", pars.Username, pars.Password)
	_, err = ftpClient.Login(pars.Username, pars.Password, "")
	if err != nil {
		fmt.Printf("The user could not be logged in, error: %s", err.Error())
	}

	_, err = ftpClient.Cwd(pars.Homefolder)
	CheckError(err)

	err = ftpClient.DownloadFile(pars.RemoteFileName, pars.SaveFileName, false)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		// os.Exit(0)
	}
}

//解壓縮
func DeCompress(filename string) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(filename[:len(filename)-4], 0644)
	fmt.Println("error: ", err)
	err = os.Chdir(filename[:len(filename)-4])
	fmt.Println("error2: ", err)
	for _, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(k.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("unzip: ", k.Name)
		defer r.Close()
		NewFile, err := os.Create(k.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
	err = os.Chdir("../")
	fmt.Println("error3: ", err)
}
