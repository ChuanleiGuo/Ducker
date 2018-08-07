package main

import (
	"chuanleiguo.com/Ducker/container"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := container.RunContainerInitProcess(tty, command)
	if err := parent.start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
