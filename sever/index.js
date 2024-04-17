// const grpc = require("grpc");
import grpc from "@grpc/grpc-js";

import protoLoader from "@grpc/proto-loader";
import { dirname } from "path";
import { fileURLToPath } from "url";
import EventEmitter from "events";
const eventEmitter = new EventEmitter();

const __dirname = dirname(fileURLToPath(import.meta.url));

var PROTO_PATH = __dirname + "/../mediaDecoding.proto";
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
var object = grpc.loadPackageDefinition(packageDefinition);
var mediaDecoding = object.mediaDecoding;
var server = new grpc.Server();

const StartVideoCutting = (call) => {
  call.on("data", async (message) => {
    console.log("Received message from client:", message);
    // İstemciye bir mesaj gönderme
    videoCutting(call);
    eventEmitter.on("finished", (data) => {
      console.log("data => ", data);
      let id = data.id;
      call.write({ Status: "Finished id : ", id });
    });
    call.write({ Status: "Hello from server" });

    // İstemci ile bağlantı kesildiğinde
    call.on("end", function () {
      // Akışı kapat
      call.end();
    });
  });

  call.on("end", function () {
    console.log("END");
    call.end(); // Akışı sonlandır
  });
};

const videoCutting = (call) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log("return");
      resolve(true);
      let id = 1;
      call.write({ Status: "Finished id : ", id });
      eventEmitter.emit("finished", { id: 1 });
    }, 7000);
  });
};

function main() {
  server.addService(mediaDecoding.mediaDecodingService.service, {
    StartVideoCutting: StartVideoCutting,
  });

  server.bindAsync(
    `0.0.0.0:${8080}`,
    grpc.ServerCredentials.createInsecure(),
    () => {
      console.log("bindAsync 0.0.0.0:" + 8080);
    }
  );
}

main();
