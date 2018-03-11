# 1. Websocket

Date: 2018-03-09

## Context

Realtime applikasi membutuhkan interaksi dua arah.


## Decision

Untuk kebutuhan chat melalui browser Websocket sudah cukup untuk menyelesaikan masalah pada realtime application. Websocket itu sendiri adalah protokol untuk komunikasi webserver browser secara real time. Protokolnya memungkinkan untuk mengirim dan menerima informasi secara terus menerus diatas koneksi TCP.

Ada beberapa candidat realtime messaging lain seperti XMPP, MQTT yang perlu dicoba.


## Consequences

Beberapa proxy dan browser ada yang belum mendukung websocket
