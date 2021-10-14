<?php
session_start();
$login = $_POST['login'];
if ($login == "parola") {
    session_regenerate_id();
    $_SESSION['SESS_LOG'] = "1";
    session_write_close();
    header("location: index.php");
    exit();
} else {
    header("location: logare.php");
    exit();
}
