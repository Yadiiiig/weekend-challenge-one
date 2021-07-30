function copy() {
    var copyText = document.getElementById("endpoint");
    copyText.select();
    document.execCommand("copy");

}

var countDownDate = new Date("08/01/2021").getTime();
var x = setInterval(function() {
    var now = new Date().getTime();
    var distance = countDownDate - now;
    var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
    var seconds = Math.floor((distance % (1000 * 60)) / 1000);

    if (hours < 10) {
        hours = "0" + hours;
    }

    if (minutes < 10) {
        minutes = "0" + minutes;
    }

    if (seconds < 10) {
        seconds = "0" + seconds;
    }
    document.getElementById("timer").innerHTML = hours + ":" + minutes + ":" + seconds;

    if (distance < 0) {
        clearInterval(x);
        document.getElementById("timer").innerHTML = "ERROR - TERMINATING";
    }
}, 1000);