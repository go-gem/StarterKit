import 'bootstrap/dist/css/bootstrap.min.css';
import './app.css';
import 'bootstrap/dist/js/bootstrap.min.js';

$(function () {
    var btnLogout = $('#btnLogout');
    var logoutForm = $('#logoutForm');
    btnLogout.click(function () {
        var confirm = window.confirm("Are you sure want to sign out?");
        if (!confirm) {
            return;
        }

        logoutForm.submit();
    });
});