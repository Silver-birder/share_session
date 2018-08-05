$(document).ready(function(){
    // modal初期設定
    $('.modal').modal({
        onOpenStart:function(modal, target){
            var is_add = $(target).text() == 'add';
            if (is_add) {
                $('#modal #title').text('ADD');
            } else {
                var $tr = $(target).parent().parent();
                var host = $tr.find('td[data="Host"]').text();
                var key = $tr.find('td[data="Key"]').text();
                var value = $tr.find('td[data="Value"]').text();
                var domain = $tr.find('td[data="Domain"]').text();
                var path = $tr.find('td[data="Path"]').text();
                var expire = $tr.find('td[data="Expire"]').text().slice(0,10);
    
                $('#modal [name=Host]').val(host).attr('readonly',true);
                $('#modal [name=Key]').val(key).attr('readonly',true);
                $('#modal [name=Value]').val(value);
                $('#modal [name=Domain]').val(domain).attr('readonly',true);
                $('#modal [name=Path]').val(path).attr('readonly',true);
                $('#modal [name=Expire]').val(expire);
                $('#modal #title').text('EDIT');
                $('#modal #Delete').css('display','');
            }
        },
        onCloseEnd:function(modal, target) {
            location.reload()
        }
    });
    $('#Done').on('click', function(){
        $.ajax({
            url:'/api/1.0/set/cookie',
            type: 'POST',
            data: $('form').serialize(),
            dataType:'json'
        })
    });
    $('#Delete').on('click', function(){
        $.ajax({
            url:'/api/1.0/delete/cookie',
            type: 'POST',
            data: $('form').serialize(),
            dataType:'json'
        })
    });
});