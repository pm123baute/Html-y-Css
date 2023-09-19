document.getElementById('enviar').addEventListener('click', function() {
   
  const url = 'http://127.0.0.1:8080/api';
  const data = {
    user: document.getElementById("user").value,
    email: document.getElementById("email").value,
    pass: document.getElementById("pass").value
  };
  console.log(data);
  console.log(JSON.stringify(data));
  fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
  })
  .then(response => response.json())
  .then(result => console.log(result))
});

function Reset(){
  form.Reset();
}