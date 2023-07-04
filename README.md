# Fetch-BackEnd-Test

#Installation Instructions:
Download the github repository provided
Open terminal and navigate to the folder containing the HTTP-SERVER repository 
Then run the following commands:
docker build -t http-server
docker run -p 8080:8080 -tid http-server

#About
Created by Carson McCombs. A web-service written in Go/Golang with the purpose of processing receipts.

#Usuage
After following the installation instructions, the web-service is able to receive POST requests from "localhost:8080/receipts/process" which will return a JSON containing a corresponding ID. 
You can then take this ID and with a GET request at "localhost:8080/{id}/points" will return the corresponding JSON containing the number of points that the receipt is worth.
You can also delete entries by using a DELETE request at "localhost:8080/receipts/{id}".

Uses port 8080 to execute.
Data is only saved in memory at runtime and is deleted after execution ends.
Tested and compiled for Windows 10/11, but should be able to create a docker images for other operating systems and have it still work.
