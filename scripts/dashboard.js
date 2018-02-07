var output = document.getElementById('output');
var elementlist = [];
var dashboardRequest = new XMLHttpRequest();
dashboardRequest.open('GET', 'http://localhost:3000/people', true);
dashboardRequest.onload = function () {
    if (dashboardRequest.readyState == 4 && dashboardRequest.status == 200) {
        var data = JSON.parse(dashboardRequest.responseText);
        // console.log(data);

        for (let index = 0; index < data.length; index++) {
            var element = data[index].firstname;
            console.log(element)
            elementlist.push(element)
        }

        output.innerHTML = elementlist;
    }

}

dashboardRequest.send();


