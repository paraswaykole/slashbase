package sshtunnel

import (
	"time"
)

type sshTunnelInstance struct {
	sshTun   *SSHTun
	LastUsed time.Time
}

var sshtunnels map[string]sshTunnelInstance = map[string]sshTunnelInstance{}

func GetSSHTunnel(dbConnID string, sshAuthType, sshHost string, remoteHost string, remotePort int, sshUser string, sshPassword, sshKeyFile string) *SSHTun {
	if eSSHtun, exists := sshtunnels[dbConnID]; exists {
		if eSSHtun.sshTun.started {
			sshtunnels[dbConnID] = sshTunnelInstance{
				sshTun:   eSSHtun.sshTun,
				LastUsed: time.Now(),
			}
			return eSSHtun.sshTun
		}
	}
	newPort := 4000 + len(sshtunnels)
	sshtun := New(newPort, sshHost, remoteHost, remotePort)
	sshtun.SetUser(sshUser)
	if sshAuthType == "KEYFILE" {
		sshtun.SetKeyFile(sshKeyFile)
	} else if sshAuthType == "PASSKEYFILE" {
		sshtun.SetEncryptedKeyFile(sshKeyFile, sshPassword)
	} else if sshAuthType == "PASSWORD" {
		sshtun.SetPassword(sshPassword)
	}
	go sshtun.Start()
	sshtunnels[dbConnID] = sshTunnelInstance{
		sshTun:   sshtun,
		LastUsed: time.Now(),
	}
	return sshtun
}

func RemoveUnusedTunnels() {
	for dbConnID, instance := range sshtunnels {
		now := time.Now()
		diff := now.Sub(instance.LastUsed)
		if diff.Minutes() > 20 {
			delete(sshtunnels, dbConnID)
			go instance.sshTun.Stop()
		}
	}
}
