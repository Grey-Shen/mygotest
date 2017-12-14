db.documents.createIndex({"app_id": 1});
db.documents.getIndexes();
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true}}).count()
var cursor = db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}});
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = doc.bizdata
    var newbizdata = doc.bizdata
    newbizdata["id_card_number"] = bizdata.idNo
    newbizdata["id_card_name"] = bizdata.trueName
    db.documents.update({_id: doc._id}, {$set: {"bizdata": newbizdata}})
}
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}}).count()
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: true}}).count()


var cursor = db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}}).limit(200);
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = doc.bizdata
}

var bizdata = {
        "channel" : "2",
        "idNo" : "12121222112",
        "channelDate" : "1988-11-10",
        "trueName" : "叶良辰",
        "subAccountNo" : "66666",
        "seqNo" : "5555",
        "userId" : "81212",
        "bankAcnt" : "123456"
    }

var cursor = db.documents.find({"app_id": "5834a331"});
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = {
            "channel" : "2",
            "idNo" : "12121222112",
            "channelDate" : "1988-11-10",
            "trueName" : "叶良辰",
            "subAccountNo" : "66666",
            "seqNo" : "5555",
            "userId" : "81212",
            "bankAcnt" : "123456",
            "my_id": doc._id
        }
    db.documents.update({_id: doc._id}, {$set: {"bizdata": bizdata}})
}

var cursor = db.documents.find({"bizdata.my_id": {$exists: true}});
while(cursor.hasNext()){
    var doc = cursor.next();
    db.documents.update({_id: doc._id}, {$unset: {"bizdata.my_id": 1}})
}
db.documents.find({"bizdata.my_id": {$exists: true}}).count()

db.documents.update({"bizdata.my_id": {$exists: true}},{$unset: {"bizdata.my_id": 1}})
db.documents.find({"bizdata.my_id": {$exists: true},"_id" : ObjectId("586e0d657984c422f8da2c74")})
