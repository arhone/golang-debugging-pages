window.addEventListener("DOMContentLoaded", function(event) {
    document.getElementsByTagName('header')[0].setAttribute('style', 'height:' + (window.innerHeight / 3) + 'px;');
    document.getElementById('logo-400').setAttribute('style', 'height:' + (window.innerHeight / 3 - 80) + 'px;');

    if (window.location.pathname === '/') {

        // Получение времени с клиента
        let dataTimeClient = document.querySelector('[data-time="client"] strong');
        if (dataTimeClient) {
            let today = new Date();
            let dd   = ('0' + (today.getDate()).toString()).slice(-2);
            let mm   = ('0' + (today.getMonth()+1).toString()).slice(-2);
            let yyyy =  today.getFullYear();
            let h    =  ('0' + today.getHours().toString()).slice(-2);
            let m    =  ('0' + today.getMinutes().toString()).slice(-2);
            let s    =  ('0' + today.getSeconds().toString()).slice(-2);
            dataTimeClient.innerHTML = dd + '.' + mm + '.' + yyyy + ' ' + h + ':' + m + ':' + s;
        }

        // Получение времени с сервера
        let xhr = new XMLHttpRequest();
        xhr.withCredentials = true;
        let dataTimeServerJson = document.querySelector('[data-time="server-json"] strong');
        window.i = 0;
        dataTimeServerJson.innerHTML = window.i.toString();
        window.timeInterval = setInterval(function () {
            window.i++;
            dataTimeServerJson.innerHTML = window.i.toString();
        }, 1000)

        xhr.addEventListener("readystatechange", function() {
            if (this.readyState === 4) {
                let responseJson = JSON.parse(this.responseText);
                clearInterval(window.timeInterval)
                dataTimeServerJson.innerHTML = responseJson.status.message;
            }
        });

        xhr.open("GET", "/api/0/time/current.json");

        xhr.send();

    }



});
