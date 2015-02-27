// Copyright 2015 toby.zxj
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// simple udp test
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Create a udp socket
	raddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:3333")
	if err != nil {
		fmt.Println("Resolve UDP Address failed:", err.Error())
		os.Exit(0)
	}
	socket, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		fmt.Println("Create udp socket failed:", err.Error())
	}
	defer socket.Close()

	// Set options
	t := time.Now().Add(time.Second * time.Duration(3)) // default is 3s timeout
	socket.SetWriteDeadline(t)
	t = time.Now().Add(time.Second * time.Duration(5))
	socket.SetReadDeadline(t)

	// Send data
	send_data := []byte("716017B69950ABC28178E19156B4EDB7:0D")
	_, err = socket.Write(send_data)
	if err != nil {
		fmt.Println("Send udp data failed:", err.Error())
		os.Exit(0)
	}

	// Receive data
	recv_data := make([]byte, 4096)
	n, raddr, err := socket.ReadFromUDP(recv_data)
	if err != nil {
		fmt.Println("Recv data failed:", err.Error())
		os.Exit(0)
	}
	fmt.Printf("Recv: % X\n", recv_data[:n])
	fmt.Println("[ToStrings] -->", string(recv_data[:n]))
}
