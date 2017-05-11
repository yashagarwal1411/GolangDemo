/**
 * Created by yashagarwal on 27/04/17.
 */

var elem = document.getElementById("alertify-element");
var autocomplete;

function isValidEmailAddress(emailAddress) {
    var pattern = /^([a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+(\.[a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+)*|"((([ \t]*\r\n)?[ \t]+)?([\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*(([ \t]*\r\n)?[ \t]+)?")@(([a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.)+([a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.?$/i;
    return pattern.test(emailAddress);
};

function initAutocomplete() {
    autocomplete = new google.maps.places.Autocomplete(
        /** @type {!HTMLInputElement} */(document.getElementById('autocomplete')),
        {types: ['geocode']});

}

function geolocate() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(function(position) {
            var geolocation = {
                lat: position.coords.latitude,
                lng: position.coords.longitude
            };
            var circle = new google.maps.Circle({
                center: geolocation,
                radius: position.coords.accuracy
            });
            autocomplete.setBounds(circle.getBounds());
        });
    }
}

$("#saveProfile").click(function () {
    var data = $('form#profile').serializeArray()
    var newData = {}
    data.forEach(function (obj) {
        newData[obj.name] = obj.value
    });
    if (!isValidEmailAddress(newData.Email)) {
        alertify.error("Invalid email provided")
        return
    }
    $.ajax({
        url: "/api/user/update",
        data: JSON.stringify(newData),
        dataType: 'json',
        contentType: "application/json; charset=utf-8",
        type: 'post',
        success: function (data) {
            console.log(data);
            alertify.success("User updated successfully")
        },
        error: function (error) {
            console.error(error);
            alertify.error("Error in saving profile")
        }
    })
});

$("#editProfile").click(function () {
    window.location.href="/profile?edit=1"
});

$('#signout').click(function () {

    $.ajax({
        url: "/api/user/signout",
        type: 'get',
        dataType: 'json',
        success: function (data) {
            var auth2 = gapi.auth2.getAuthInstance();
            auth2.signOut().then(function () {
                console.log('User signed out.');
                window.location.href = "/";
            });
        },
        error: function (error) {
            console.error(error)
            alertify.error("Error in signing out")
        }
    })
})

var init = function () {
    alertify.parent(elem);

    gapi.load('auth2', function() {
        gapi.auth2.init();
    });

    if (window.location.href.indexOf("view=1") > -1) {
        $("#saveProfile").hide();
        $("input[name=FullName]").prop('disabled', true)
        $("input[name=Address]").prop('disabled', true)
        $("input[name=Email]").prop('disabled', true)
        $("input[name=Telephone]").prop('disabled', true)
    } else {
        $('#editProfile').hide();
    }

    $.ajax({
        url: "/api/user/get",
        type: 'get',
        dataType: 'json',
        success: function (data) {
            console.log("LALLA data is: ")
            console.log(data);
            $("input[name=Id]").val(data.Id)
            $("input[name=FullName]").val(data.FullName)
            $("input[name=Address]").val(data.Address)
            $("input[name=Email]").val(data.Email)
            $("input[name=Telephone]").val(data.Telephone)
        },
        error: function (error) {
            console.error(error)
            alertify.error("Error in getting user data")
        }
    });
}

init();

