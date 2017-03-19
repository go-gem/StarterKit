$(function () {
    var form = $('#joinForm');
    var btn = $('#btnSubmit');
    btn.click(function () {
        $.post('/join', form.serialize(), function (resp) {
            alert(resp.msg);

            if (resp.code != 0) {
                return;
            }

            window.location.href = '/login';
        }, 'json');
    });
});