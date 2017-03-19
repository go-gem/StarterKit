$(function () {
    var form = $('#loginForm');
    var btn = $('#btnSubmit');
    btn.click(function () {
        $.post('/login', form.serialize(), function (resp) {
            alert(resp.msg);

            if (resp.code != 0) {
                return;
            }

            window.location.href = '/';
        }, 'json');
    });
});