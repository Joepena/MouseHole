require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

$(() => {

});

window.onload = function () {
    var conn;
    var eventsTable = document.getElementById("events-table");

    function appendEventToTable(jsonEvent) {
        var event = JSON.parse(jsonEvent.data);
        var newRow = eventsTable.insertRow();
        newRow.innerHTML = `
                           <td class="centered">
                                ${event.Title}
                             </td>
                             <td>
                                ${event.Content} 
                             </td>
                             <td>
                                ${event.Tags}
                             </td>`
    }

    if (window["WebSocket"]) {
        console.log(document.location.host);
        conn = new WebSocket("ws://" + document.location.host + '/events_socket');
        conn.onclose = function (evt) {
            var item = document.createElement('div');
            item.innerHTML = "<b>Connection closed.</b>";
        };
        conn.onmessage = function (event) {
            appendEventToTable(event)
            }
    };
}

