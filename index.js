import mysql from 'mysql';
var sql = "SELECT `ID`,`NombreCompleto` From registros WHERE User=? AND Contraseña=? AND Email=?"
const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '',
    database: 'casa'
});

export function Enviar() {
    
    connection.connect()

    connection.query(sql, function (error, results, fields) {
        if (error) throw error;        
        console.log(results);
        connection.end();
      });
    
}


function Reset(){
    form.Reset();
}