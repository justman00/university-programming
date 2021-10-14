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

    $sql_query = "INSERT INTO users VALUES (NULL, '%s', '%s');";
    $sql = sprintf($sql_query, $_POST['name'], $_POST['password']);

    if($result = $conn->query($sql)) {
        header('Location: /');
    } 

    echo "Something went wrong";
?>