// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package grpc

import "github.com/grpc/grpc-go"

func Start() {
	server := grpc.NewServer()
	server.Serve()

}
