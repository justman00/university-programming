<?php
session_start();
if(!isset($_SESSION['SESS_LOG']))
{
header("location: logare.php");
exit();
}
