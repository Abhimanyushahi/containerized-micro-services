const http = require('http');

const server = http.createServer((request,response)=>{
        response.end("Hello, World! By Abhimanyu Shahi");
});

server.listen(5000,()=>{
        console.log("Server is Running ...");
});
