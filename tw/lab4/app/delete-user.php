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

    $sql_query = "DELETE FROM users
    WHERE username = '%s';";
    $sql = sprintf($sql_query, $_POST['name']);

    if($result = $conn->query($sql)) {
        header('Location: /');
    } 

    echo "Something went wrong";
?>