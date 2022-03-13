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

    db.albums.insertOne({"name":"tutorials point", "content": "some content"})

    db.albums.findOne({"name":"tutorials point"})
EOF