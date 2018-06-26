// https://closure-compiler.appspot.com/home 

void((function(f){
    var script = document.createElement('script');
    script.src = 'https://code.jquery.com/jquery-3.2.1.min.js';
    script.onload = function(){
      var b = jQuery.noConflict(true);
      f(b);
    };
    document.body.appendChild(script);
})(
function(b, undefined){
    var domain = document.domain;
    var data = {
        'Domain':domain
    }
    b.ajax({
        type: "post",
        url: "https://ma-share-session.appspot.com/api/1.0/get/cookie",
        timeout: 20000,
        cache: false,
        data: data,
        dataType: 'json'
    })
    .done(function (response, textStatus, jqXHR) {
        if (response.length == 0) {
            alert(domain+"のクッキーがありません")
            return ;
        }
        for (index in response) {
            var key = response[index].key;
            var value = response[index].value;
            document.cookie = key + "=" + value;
        }
        alert("クッキーを設定しました。")
    })
    .fail(function (jqXHR, textStatus, errorThrown) {
        alert("失敗: サーバー内でエラーがあったか、サーバーから応答がありませんでした。");
    })
}
)
);