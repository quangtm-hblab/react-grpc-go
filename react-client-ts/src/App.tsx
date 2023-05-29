import React from 'react';
import logo from './logo.svg';
import './App.css';
import * as grpcWeb from "grpc-web"
import { GreeterClient } from './proto/src/output/HelloServiceClientPb';
import { HelloRequest, HelloReply } from './proto/src/output/hello_pb';

function App() {
  const grpcCall =()=>{
    
    const req = new HelloRequest()
    req.setName("quang dep zai")
    const greeterService = new GreeterClient('http://localhost:8080', null, null)
    const call = greeterService.sayHello(req, null,(err: grpcWeb.RpcError, response: HelloReply)=>{
      if (err) {
        console.log(err)
        return
      }
      console.log(response)
    })

    call.on('status', (status: grpcWeb.Status) => {
      console.log("status:", status);
    });
  }
  return (
    <div className="App">
      <button onClick={grpcCall}>Click</button>
    </div>
  );
}

export default App;
