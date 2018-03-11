# Test for backend engineer Prismapp

## Challenge

Please create a chat backend application that exposes 3 public API:

1. API to send message:
   * Accept one parameter: the string message to sent.
   * This API should be a pull API (e.g. REST/GraphQL, etc).

2. API to retrieve messages:
   * To retrieve all previously-sent messages.
   * This API should be also an pull API.

3. API to retrieve message at realtime:
   * This will create a long lived connection that will retrieve message at realtime.
   * This API should be a push API (Websocket/MQTT, etc).

Some notes:

1. Feel free to use whatever tools/library necessary.
2. No need to think about authentication, etc.
3. The goal here is not only to measure your technical capabilities, but also to understand how you make design decision and its tradeoffs. So, creating a simple design-decision document would be huge plus for us. Short and concise document is preferred.

## Identify Issue

Tugasnya adalah membuat backend untuk chat application. Tugas dibagi menjadi 3 API:

1. API untuk mengirimkan pesan, API ini bertugas melakukan pengiriman pesan dan pesan akan dicatat kedalam storage
2. API untuk mengambil pesan yang telah dikirimkan, disini semua chat akan ditampilkan.
3. API untuk mendapatkan pesan secara realtime, jadi ketika user mengirimkan dari API 1 maka API 3 akan secara langsung mendapat pesannya

## Solution

Kita membutuhkan suatu proses yang memecah antara pesan yang tersimpan distorage dan pesan yang langsung dikirimkan.
~~Ketika API no 1 menerima pesan maka akan ada "worker" yang mem-"publish" ke dua service yang berberda yaitu, penyimpanan dan penerimaan pesan ke client.~~
Ketika API no 1 menerima pesan maka ada 2 step yang dilakukan , pertama kirim pesan ke websocket lalu menginput pesan kedalam database

## Use Case

Dengan ini kita bisa membataskan konteks dalam 3 use case yaitu:

    - Send New Message in a room of chat.
    - Retrieve Previous Message in a room of chat.
    - Retrieve Message in realtime.

## Design API

1. [POST] /chats    - Membuat chat baru - satu arah
2. [GET]  /chats    - Membaca semua chat
3. [GET]  ws:/realtime - Membaca chat terbaru secara realtime / dua arah


## Requirement 

1. Programming Language:
    * Go Langunage 19.3

2. Persistance Storage
    * Mariadb 10.1

3. Realtime
    * Websocket

## How to running

1. Install go language
2. Install mariadb 10.1
3. Create a new database with name : `prismapp-test`
4. Create new table
    ```SQL
        CREATE TABLE `room` (
            `id` varchar(36) NOT NULL,
            `message` longtext NOT NULL,
            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
            PRIMARY KEY (`id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
    ```
5. clone this project from 
6. open terminal and goto project path 
7. cd /cmd/chatd
8. run go main.go or go build -o chatd && ./chatd
9. go to browser open localhost:8080
10. or goto postman

Example in Gif

https://i.imgur.com/OyU763Z.gif
