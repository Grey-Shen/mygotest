//备份要修改的数据
mongoexport -d tassadar_release -c documents -q '{"app_id": "97dcfa02"}' -o document.json
// 创建索引
db.documents.createIndex({"app_id": 1});

//确认创建的索引
db.documents.getIndexes();

// 要修改的数量
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true}}).count()

// 执行修改
var cursor = db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}});
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = doc.bizdata
    var newbizdata = doc.bizdata
    newbizdata["id_card_number"] = bizdata.idNo
    newbizdata["id_card_name"] = bizdata.trueName
    db.documents.update({_id: doc._id}, {$set: {"bizdata": newbizdata}})
}

// 确认未修改的数量
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}}).count()

// 确认已修改的数量
db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: true}}).count()




mongoexport -d tassadar_release -c apps -q  -o apps.json


db.documents.createIndex({"operator.created_at": 1});
db.documents.createIndex({"operator.updated_at": 1});
db.createIndex({"app_name": 1});
db.apps.createIndex({"access_key": 1});
tassadar_release.apps.createIndex({"bucket_name": 1});
tassadar_release.apps.createIndex({"bucket_domain": 1});