<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<!-- de continua tema din lab.2, unele pagini fiind transformate astfel incit sa fie create dinamic
b) trebuie sa contina fonctionalitati ca conectare la BD, gestionarea informatiei din citeva tebele
(citire, scriere, modificare, stergere)
c) utilizati sesiuni pentru pastrarea datelor de authenificare pentru paginile utilizator cu
functionalitati privilegiate
d) exersati citirea/scrierea in fisiere prin implementarea unor exemple, chiar daca in logica siteului nu e nevoie. -->

<body>
    <?php
    $host = 'db';

    // Database use name
    $user = 'MYSQL_USER';
    
    //database user password
    $pass = 'MYSQL_PASSWORD';
    
    // database name
    $mydatabase = 'MYSQL_DATABASE';
    // check the mysql connection status
    
    $conn = new mysqli($host, $user, $pass, $mydatabase);
    
    // select query
    $sql = 'SELECT * FROM users';
    
    if ($result = $conn->query($sql)) {
        while ($data = $result->fetch_object()) {
            $users[] = $data;
        }
    }
    
    foreach ($users as $user) {
        echo "<br>";
        echo $user->username . " " . $user->password;
        echo "<br>";
    }
    
    ?>
</body>

</html>

</html>