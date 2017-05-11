/**
 * Created by yashagarwal on 27/04/17.
 */

var elem = document.getElementById("alertify-element");
alertify.parent(elem);

function isValidEmailAddress(emailAddress) {
    var pattern = /^([a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+(\.[a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+)*|"((([ \t]*\r\n)?[ \t]+)?([\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*(([ \t]*\r\n)?[ \t]+)?")@(([a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.)+([a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.?$/i;
    return pattern.test(emailAddress);
};

$("#signUp").submit(function (e) {
    e.preventDefault()
    console.log("IN SIGN UP SUBMIT")
    console.log($('form#signUp').serializeArray())
    var data = $('form#signUp').serializeArray()
    var newData = {}
    data.forEach(function (obj) {
        newData[obj.name] = obj.value
    })
    if (!isValidEmailAddress(newData.Email)) {
        alertify.error("Invalid email provided")
        return
    }
    if (newData.Password.length < 4) {
        alertify.error("Password needs to be atleast 4 characters")
        return
    }
    console.log(newData)
    $.ajax({
        url: "/api/signup",
        data: JSON.stringify(newData),
        dataType: 'json',
        contentType: "application/json; charset=utf-8",
        type: 'post',
        success: function (data) {
            console.log(data);
            window.location.href = "/profile?edit=1"
        },
        error: function (error) {
            console.error(error);
            alertify.error("Error in signing up")
        }
    })
})

$("#signIn").submit(function (e) {
    console.log("IN SIGN In SUBMIT")
    e.preventDefault()
    var data = $('form#signIn').serializeArray()
    var newData = {}
    data.forEach(function (obj) {
        newData[obj.name] = obj.value
    })
    if (!isValidEmailAddress(newData.Email)) {
        alertify.error("Invalid email provided")
        return
    }
    console.log(newData)
    $.ajax({
        url: "/api/signin",
        data: JSON.stringify(newData),
        dataType: 'json',
        contentType: "application/json; charset=utf-8",
        type: 'post',
        success: function (data) {
            console.log(data);
            window.location.href = "/profile?view=1"
        },
        error: function (error) {
            console.error(error);
            alertify.error("Email/Password invalid")
        }
    })
})

function onSignIn(googleUser) {
    console.log("HERERER GOOGLE SIGN IN SUCCESS YO");
    var profile = googleUser.getBasicProfile();
    var idToken = googleUser.getAuthResponse().id_token;
    console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
    console.log('Name: ' + profile.getName());
    console.log('IdToken: ' + idToken);
    console.log('Image URL: ' + profile.getImageUrl());
    console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.

    $.ajax({
        url: "/api/googlesignin",
        data: JSON.stringify({IdToken: idToken}),
        dataType: 'json',
        contentType: "application/json; charset=utf-8",
        type: 'post',
        success: function (data) {
            console.log(data);
            window.location.href = "/profile?view=1"
        },
        error: function (error) {
            console.error(error);
            alertify.error("Error in Google signin")
        }
    })
}
