fetch('http:localhost:8080/track?q=Mumbai&appid=1f485364c2c92f74eaca9d0bb5768b5f')
    .then(res => res.json())
    .then(data => console.log(data))
    .catch(error => console.error(error))
