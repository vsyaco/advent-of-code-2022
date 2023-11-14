<?php

$file = fopen('../data.txt', 'r');
$current = 0;
$max = 0;

function onUserElfChanged()
{
    global $current, $max;

    if ($current >= $max) {
        $max = $current;
    }

    $current = 0;
}

while (($buffer = fgets($file, 4096)) !== false) {
    if ($buffer === PHP_EOL) {
        onUserElfChanged();
    } else {
        $current += intval($buffer);
    }
}
fclose($file);

echo $max . PHP_EOL;
