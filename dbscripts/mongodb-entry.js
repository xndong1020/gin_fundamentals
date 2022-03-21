var db = connect("mongodb://root:password@localhost:27017/admin");

db = db.getSiblingDB('albumsDb');

db.createCollection("albums")

db.albums.insertOne({"name":"[Blue Train] featured by John Coltrane", "content": "content body for Blue Train", "albumId": 1})
db.albums.insertOne({"name":"[Jeru] featured by Gerry Mulligan", "content": "content body for Jeru", "albumId": 2})
db.albums.insertOne({"name":"[Sarah Vaughan and Clifford Brown] featured by Sarah Vaughan", "content": "content body for Sarah Vaughan and Clifford Brown", "albumId": 3})

db.albums.findOne({"albumId":1})