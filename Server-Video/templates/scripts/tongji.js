$(document).ready(function () {
    /**
     * 上报用户信息，以及访问数据到打点服务器
     */
    var lock = true;

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/dig',
        type: 'get',
        data: {
            "time": gettime(),
            "url": geturl(),
            "refer": getrefer(),
            "ua": getuser_agent()
        },
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                callback(null, "Internal Error");
                lock = false; //使用锁防止引用多次
            }
        }
    }).done(function (data, statusText, xhr) {
        callback(data, null);
        lock = false; //使用锁防止引用多次
    });


})


function gettime() {
    var nowDate = new Date();
    return nowDate.toLocaleString();
}

function geturl() {
    return window.location.href;
}

function getrefer() {
    return document.referrer;
}

function getuser_agent() {
    return navigator.userAgent;
}