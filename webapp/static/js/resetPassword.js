/**
 * Created by yashagarwal on 27/04/17.
 */

var elem = document.getElementById("alertify-element");
alertify.parent(elem);

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

$("#submit").click(function () {
    var pass = $("#newPassword").val()
    var confirmPass = $("#confirmNewPassword").val()
    if (pass != confirmPass) {
        console.error("Both passwords dont match")
        alertify.error("Both passwords dont match")
        return
    }
    if (pass.length < 4) {
        console.error("Password should be more than 3 chars")
        alertify.error("Password should be more than 3 chars")
        return
    }
    var data = {
        Password: pass,
        Email: getParameterByName("email"),
        Token: getParameterByName("token")
    }
    $.ajax({
        url: "/api/resetPassword",
        type: 'post',
        dataType: 'json',
        data: JSON.stringify(data),
        success: function (data) {
            console.log("Password successfully changed")
            alertify.success("Password successfully changed");
            setTimeout(function () {
                window.location.href = "/"
            }, 3000)
        },
        error: function (error) {
            console.error(error)
            alertify.error("Error in changing password");
        }
    })
})