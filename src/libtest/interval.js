var tassadar_release = db.getSiblingDB("tassadar_release");
var mapFunction = function() {
    var keyArray = new Array()
    for (var idx = 0; idx < this.pages.length; idx++) {
        var bizdata = this.pages[idx].bizdata;
        for (var bizdatakey in bizdata) {
            keyArray.push(bizdatakey)
        }
    }

    for (var bizdatakey in this.bizdata) {
        keyArray.push(bizdatakey)
    }

    if (keyArray.length >0) {
        emit(this.app_id,{bizdatakeys:keyArray});
    }
}

var reduceFunction = function(key,values) {
    result = {bizdatakeys:[]};
    values.forEach(function (v) {
        for (var key of v.bizdatakeys) {
            if (result.bizdatakeys.indexOf(key) == -1){
                result.bizdatakeys.push(key)
            }
        }
    });
    return result;
}


var mapFunction = function() {
    var ahead = this.end_at.valueOf()- new Date("2017-08-15").valueOf()
    var quotient = parseInt(ahead / 1000000)
    if (quotient > 0) {
        emit(quotient,this.count)
    }
}

var reduceFunction = function(key,values) {
    return Array.sum(values);
}

var queryFunction = function() {
    this.end_at > new Date("2017-06-06")
}

db.app_stats.mapReduce(mapFunction,reduceFunction,{ out: { merge: "app_test" }})

db.runCommand(
                {
                    mapReduce: "app_stats",
                    map: mapFunction,
                    reduce: reduceFunction,
                    out: {merge: "app_test"}
                }
            )

db.app_stats.mapReduce(
                            mapFunction,
                            reduceFunction,
                            {
                                out: { merge: "app_test" },
                                query: {"begin_at": {"$gte": new Date("2017-08-15")}}
                            }
                      )