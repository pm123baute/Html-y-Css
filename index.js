function Reset(){
    formulario.Reset()
}

function Enviar(){
    const mysql = require('mysql');
    const connection = mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: '',
        database: 'casa'
      });
      
}