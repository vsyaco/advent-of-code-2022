<?php

$file = fopen('../data.txt', 'r');
$current = 0;
$topMax = [0];

function onUserElfChanged()
{
    global $current, $topMax;

    $minValue = intval(min($topMax));
    $currentInt = intval($current);
    $current = 0;

    if ($currentInt <= $minValue) {
        return;
    }

    if (count($topMax) < 3) {
        $topMax[] = $currentInt;
        return;
    }

    foreach ($topMax as &$value) {
        if ($minValue === $value) {
            $value = $currentInt;
        }
    }
    unset($value);
}

while (($buffer = fgets($file, 4096)) !== false) {
    if ($buffer === PHP_EOL) {
        onUserElfChanged();
    } else {
        $current += intval($buffer);
    }
}
fclose($file);

echo array_sum($topMax) . PHP_EOL;
