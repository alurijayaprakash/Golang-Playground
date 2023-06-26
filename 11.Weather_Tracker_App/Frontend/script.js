// fetch('http:localhost:8080/track?q=Mumbai&appid=1f485364c2c92f74eaca9d0bb5768b5f')
//     .then(res => res.json())
//     .then(data => console.log(data))
//     .catch(error => console.error(error))


let MyForm = document.getElementById("form_id");
MyForm.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log(e,"event123");
    // handle submit
    let city_name = document.getElementById("city_name");
    // check input values
    if (city_name.value == "") {
        alert("Ensure you input a value in both fields!");
    } else {
        // perform operation with form input
        // alert("This form has been successfully submitted!");
        console.log(
            `This form has a CityName of ${city_name.value}`
        );
    }
    console.log(city_name.value)
    
    API_URL = "http:localhost:8080/track?q=" + city_name.value
    fetch(API_URL)
    .then(res => res.json())
    // .then(data => console.log(data))
    .then(data => {
        // console.log(data)
        let tableData = ""
        tableData += `<tr>
        <td>${data.name}</td>
        <td>${Math.round(data.main.temp - 273.15)}</td>
        <td>${data.visibility}</td>
        <td>${data.main.pressure}</td>
        <td>${data.main.humidity}</td>
        <td>${data.wind.speed}</td>
        </tr>`;


        document.getElementById("table_body").innerHTML=tableData

    }).catch(err => console.log(err));
});