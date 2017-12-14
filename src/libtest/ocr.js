var cursor = db.restaurants.find();
while(cursor.hasNext()){
    var doc = cursor.next();
    var grades = doc.grades
    var id = doc._id
    for (var i = 0; i < grades.length; i++) {
        var grade = grades[i].grade
        var t = {}
        t[`grades.${i}.greade_test`] = grade
        t[`grades.${i}.greade_test1`] = grade
        t[`grades.${i}.greade_test2`] = grade
        t[`grades.${i}.greade_test3`] = grade
        t[`grades.${i}.greade_test4`] = grade
        t[`grades.${i}.greade_test5`] = grade

        db.restaurants.update({_id: id},{$unset: t})
    }
}


身份证号码：idNo --  id_card_number

姓名：trueName -- id_card_name


var cursor = db.restaurants.find({"bizdata.idNo": {$exists: false}});
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = doc.bizdata
    var newpages = pages
    var id = doc._id
    for (var i = 0; i < grades.length; i++) {
        newpages[i].bizdata["greade_test"] = grades[i].bizdata.grade
    }
    db.restaurants.update({"_id": id,"grades.grade_test2": {$exists: false}},{$set: {"grades": up}})
}




var cursor = db.documents.find({"app_id": "5834a331","bizdata.idNo": {$exists: true},"id_card_number": {$exists: false}});
while(cursor.hasNext()){
    var doc = cursor.next();
    var bizdata = doc.bizdata
    var newbizdata = doc.bizdata
    newbizdata["id_card_number"] = bizdata.idNo
    newbizdata["id_card_name"] = bizdata.trueName
    db.restaurants.update({"app_id": "5834a331","bizdata.idNo": {$exists: true},"bizdata.id_card_number": {$exists: false}},{$set: {"bizdata": newbizdata}})
}


mongoexport -d tassadar_test -c documents -q '{"doc_id": {$in :["UDMP-0d4ee5c774c59bb35f49a92e71e33012d27aa85e68-bbd90991-20170227033519-00000002","UDMP-a7c4c886a9c31e498f5b8aeaafa4290bd5662bf8be-bbd90991-20170227033519-00000002"]}}' -o document.json
mongoimport -d tassadar_test -c documents --file ./document.json

db.documents.deleteOne({ "doc_id" : "UDMP-a7c4c886a9c31e498f5b8aeaafa4290bd5662bf8be-bbd90991-20170227033519-00000002" })
db.documents.deleteOne({ "doc_id" : "UDMP-0d4ee5c774c59bb35f49a92e71e33012d27aa85e68-bbd90991-20170227033519-00000002" })

