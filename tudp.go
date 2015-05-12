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
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	remote := flag.String("r", "127.0.0.1:3333", "Remote udp address")
	msg := flag.String("m", "Hello GiterLab!", "Udp message.")
	enter := flag.Bool("n", false, "Send message end with \\n.")
	timeout := flag.Int("t", 0, "Setup udp timeout.")
	debug_mode := flag.Bool("d", false, "Open debug mode.")
	flag.Parse()
	debug := *debug_mode

	if debug {
		fmt.Printf("Send message: \"%s\" to %s\n", *msg, *remote)
	}

	if *enter {
		*msg = *msg + "\n"
	}

	// Create a udp socket
	raddr, err := net.ResolveUDPAddr("udp4", *remote)
	if err != nil {
		if debug {
			fmt.Println("Resolve UDP Address failed:", err.Error())
		}
		os.Exit(0)
	}
	socket, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		if debug {
			fmt.Println("Create udp socket failed:", err.Error())
		}
	}
	defer socket.Close()

	// Set options
	if *timeout == 0 {
		t := time.Now().Add(time.Second * time.Duration(1)) // default is 1s timeout: send
		socket.SetWriteDeadline(t)
		t = time.Now().Add(time.Second * time.Duration(2)) // default is 2s timeout: recv
		socket.SetReadDeadline(t)
	} else {
		t := time.Now().Add(time.Second * time.Duration(*timeout)) // default is 1s timeout: send
		socket.SetWriteDeadline(t)
		socket.SetReadDeadline(t)
	}

	// Send data
	// "716017B69950ABC28178E19156B4EDB7:0D"
	send_data := []byte(*msg)
	_, err = socket.Write(send_data)
	if err != nil {
		if debug {
			fmt.Println("Send udp data failed:", err.Error())
		}
		os.Exit(0)
	}

	// Receive data
	recv_data := make([]byte, 4096)
	n, raddr, err := socket.ReadFromUDP(recv_data)
	if err != nil {
		if debug {
			fmt.Println("Recv data failed:", err.Error())
		}
		os.Exit(0)
	}
	if debug {
		fmt.Printf("Recv: % X\n", recv_data[:n])
		fmt.Println("[ToStrings] -->", string(recv_data[:n]))
	}
}
