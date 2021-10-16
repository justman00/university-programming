<?php
include "sesiune.php";

?>

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <h1>Bine ai venit pe platforma de administrare a utilizatorilor</h1>

    <div>
        <a href="create-user.html">Adauga un utilizator nou</a>
    </div>

    <h3>List curenta de utilizatori</h3>
    <div>
        Pentru a modifica parola unui utilizator urmeaza:
        <a href="edit-user-password.html">editeaza parola</a>
    </div>
    <div>
        Pentru a sterge un utilizator urmeaza:
        <a href="delete-user.html">sterge utilizator</a>
    </div>
    <?php
    function connectDB()
    {
        $host = 'db';

        // Database use name
        $user = 'MYSQL_USER';

        //database user password
        $pass = 'MYSQL_PASSWORD';

        // database name
        $mydatabase = 'MYSQL_DATABASE';
        // check the mysql connection status

        $conn = new mysqli($host, $user, $pass, $mydatabase);

        return $conn;
    }

    $conn = connectDB();


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

    file_put_contents('visits.txt', sprintf("%s visited on the %s\n", $_SERVER['REMOTE_ADDR'], $_SERVER['REQUEST_TIME']), FILE_APPEND);
    $file_contents = file_get_contents('visits.txt');

    echo "<h2>Lista de accesari a acestei pagini</h2>";
    echo "<ul>";

    $visits_arr = explode("\n", $file_contents);

    foreach($visits_arr as $visit) {
        if (!$visit) {
            continue;
        }

        echo "<li>" . $visit . "</li>";
    }

    echo "</ul>";

    ?>
</body>

</html>