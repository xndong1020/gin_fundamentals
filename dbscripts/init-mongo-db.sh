#!/bin/bash
set -e

echo "creating mongodb"

mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
    var rootUser = '$MONGO_INITDB_ROOT_USERNAME';
    var rootPassword = '$MONGO_INITDB_ROOT_PASSWORD';
    var admin = db.getSiblingDB('admin');
    admin.auth(rootUser, rootPassword);

    # for creating a non-admin user
    # var user = '$MONGO_INITDB_ROOT_USERNAME';
    # var passwd = '$MONGO_INITDB_ROOT_PASSWORD';
    # db.createUser({user: user, pwd: passwd, roles: ["readWrite"]});

    show dbs

    # create albumsDb database, then switch to albumsDb
    use albumsDb

    # create collection albums
    db.createCollection("albums")

    db.albums.insertOne({"name":"[Blue Train] featured by John Coltrane", "content": "content body for Blue Train", "albumId": 1})
    db.albums.insertOne({"name":"[Jeru] featured by Gerry Mulligan", "content": "content body for Jeru", "albumId": 2})
    db.albums.insertOne({"name":"[Sarah Vaughan and Clifford Brown] featured by Sarah Vaughan", "content": "content body for Sarah Vaughan and Clifford Brown", "albumId": 3})

    db.albums.findOne({"albumId":1})
EOF