$(function () {
    var form = $('#postForm');
    var btn = $('#btnSubmit');
    btn.click(function () {
        $.post('/new', form.serialize(), function (resp) {
            alert(resp.msg);

            if (resp.code != 0) {
                return;
            }

            window.location.href = '/post/' + resp.data;
        }, 'json');
    });
});